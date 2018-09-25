package models

import (
	"database/sql"
	"fmt"
	"log"
)

type (
	//Contact in phonebook
	Contact struct {
		ID         int    `json:"id"`
		Firstname  string `json:"firstname"`
		Secondname string `json:"secondname"`
		Sinonim    string `json:"sinonim"`
		Prefix     string `json:"prefix"`
		Number     int    `json:"number"`
		Active     bool   `json:"active"`
	}

	//Contacts in phonebook
	Contacts struct {
		Contacts []Contact `json:"contacts"`
	}
)

//SelectAll select all rows in phonebook
func SelectAll(db *sql.DB) Contacts {
	err := db.Ping()
	if err != nil {
		log.Fatal("\n Пинг БД \n", err)
	}

	selectAllCont := `select * from phonebook`
	rows, err := db.Query(selectAllCont)

	if err == nil {
		fmt.Printf("Запрос выполнен: %s \n", selectAllCont)
	}

	result := Contacts{}

	//цикл по полученным строкам
	for rows.Next() {
		newContact := Contact{}
		//записываем результат в модель
		err := rows.Scan(
			&newContact.ID,
			&newContact.Firstname,
			&newContact.Secondname,
			&newContact.Sinonim,
			&newContact.Prefix,
			&newContact.Number,
			&newContact.Active,
		)
		//обработка ошибок
		if err != nil {
			fmt.Printf("Не удалось прочитать контакты, ошибка %v \n", err)
		} else {
			fmt.Printf("Прочитали контакты %v \n", newContact)
		}
		//добавляем полученный контакт в коллекцию
		result.Contacts = append(result.Contacts, newContact)

	}
	defer rows.Close()
	//Возвращаем коллекцию контактов
	log.Println(result)
	return result
}

//SelectAllNumbers contact func
func SelectAllNumbers(db *sql.DB) ([]int, error) {
	var (
		num     int
		numbers []int
	)
	selectNumbers := `select number from phonebook`
	stmt, err := db.Prepare(selectNumbers)

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		err = rows.Scan(num)
		numbers = append(numbers, num)
	}

	fmt.Println(numbers)
	defer stmt.Close()
	//возвращаем err
	return numbers, err
}

//SelectContact for search one contact
func SelectContact(db *sql.DB, number int) (Contact, error) {
	var selContact Contact

	rowContact := db.QueryRow(`
	 select * from phonebook
	 where Number like = ?
	`, number)

	err := rowContact.Scan(
		&selContact.ID,
		&selContact.Firstname,
		&selContact.Secondname,
		&selContact.Sinonim,
		&selContact.Prefix,
		&selContact.Number,
		&selContact.Active,
	)

	return selContact, err
}

//InsertContact contact func
func InsertContact(db *sql.DB, cont *Contact) (int64, error) {
	insertCont := `
	insert into phonebook (Firstname, Secondname, Sinonim, Prefix, Number, Active)
	 values( ?, ?, ?, ?, ?, true);
	 `
	//записываем результат в модель
	result, errExec := db.Exec(
		insertCont,
		cont.Firstname,
		cont.Secondname,
		cont.Sinonim,
		cont.Prefix,
		cont.Number)

	//обработка ошибок
	if errExec != nil {
		fmt.Printf("Не удалось добавить контак %v. Ошибки: %v \n", cont.ID, errExec)
	}

	return result.LastInsertId()
}

//DeleteByID contact func
func DeleteByID(db *sql.DB, id int) (int64, error) {
	delCont := "delete from phonebookdb.phonebook where id = ?"

	// выполним SQL запрос
	sql, err := db.Prepare(delCont)
	// выход при ошибке
	if err != nil {
		fmt.Println(err)
	}

	// заменим символ '?' в запросе на 'id'
	result, errDel := sql.Exec(id)
	// выход при ошибке
	if errDel != nil {
		fmt.Println(errDel)
	}
	//возврящаем id записи
	return result.RowsAffected()
}

//UpdateContact func
func UpdateContact(db *sql.DB, cont *Contact) (int64, error) {
	updateCont := `update phonebook set firstname=?, secondname=?, sinonim=?, prefix=?, number=?, active=?`

	stmt, err := db.Prepare(updateCont)
	result, err := stmt.Exec(
		cont.Firstname,
		cont.Secondname,
		cont.Sinonim,
		cont.Prefix,
		cont.Number,
		cont.Active,
	)

	if err != nil {
		fmt.Println(err)
	}

	return result.LastInsertId()
}
