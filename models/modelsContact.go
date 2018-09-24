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
			&newContact.Active)
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
func SelectAllNumbers(db *sql.DB, findeNumber chan []int, number int) ([]int, bool) {
	var (
		num     int
		numbers []int
	)
	selectNumbers := `select number from phonebook`
	stmt, err := db.Prepare(selectNumbers)

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
		return nil, false
	}

	for rows.Next() {
		err = rows.Scan(num)
		numbers = append(numbers, num)
		findeNumber <- numbers
	}

	defer stmt.Close()
	//возвращаем true
	return numbers, true
}

//CreateNewContact contact func
func CreateNewContact(db *sql.DB, cont *Contact) (int64, error) {
	updateCont := `insert phonebook
	set  firstname=?, seconsname=?, sinonim=?, prefix=?, number=?, active=?`

	stmt, err := db.Prepare(updateCont)

	row, err := stmt.Exec(
		cont.Firstname,
		cont.Secondname,
		cont.Sinonim,
		cont.Prefix,
		cont.Number,
		cont.Active)
	//обработка ошибок
	id64, err := row.LastInsertId()
	if err != nil {
		fmt.Printf("Не удалось обновить контак id=%v. Ошибки: %v \n", id64, err)
	}
	defer stmt.Close()
	//возврящаем id записи
	return id64, err
}

//Insert contact func
func Insert(db *sql.DB, cont *Contact) (int64, error) {
	insertCont := `
	insert into phonebookdb.phonebook (Firstname, Secondname, Sinonim, Prefix, Number, Active)
	 values( ?, ?, ?, ?, ?, true);
	 `
	//записываем результат в модель
	row, errExec := db.Exec(
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
	defer db.Close()
	//возврящаем id записи
	return row.LastInsertId()
}

//DeleteByID contact func
func DeleteByID(db *sql.DB, id int) (int64, error) {
	delCont := "delet from phonebookdb.phonebook where id = ?"

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
