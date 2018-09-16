package models

import (
	"database/sql"
	"fmt"
)

//Contact in phonebook
type Contact struct {
	ID         int    `json:"id"`
	Firstname  string `json:"firstname"`
	Secondname string `json:"secondname"`
	Sinonim    string `json:"sinonim"`
	Prefix     string `json:"prefix"`
	Number     int    `json:"number"`
	Note       string `json:"note"`
}

//Contacts in phonebook
type Contacts struct {
	Contacts []Contact `json:"contacts"`
}

//GetAllContacts select all rows in phonebook
func GetAllContacts(db *sql.DB) Contacts {
	selectAllCont := `select * from phonebookdb.phonebook;`
	fmt.Println("\n Выполняем запрос к БД")
	contactRows, err := db.Query(selectAllCont)

	if err == nil {
		fmt.Printf("Запрос выполнен: %v \n", selectAllCont)
	}

	var result = Contacts{}
	//цикл по полученным строкам
	for contactRows.Next() {
		var newContact = Contact{}
		//записываем результат в модель
		err := contactRows.Scan(
			&newContact.ID,
			&newContact.Firstname,
			&newContact.Secondname,
			&newContact.Sinonim,
			&newContact.Prefix,
			&newContact.Number,
			&newContact.Note,
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
	defer contactRows.Close()
	//Возвращаем коллекцию контактов
	return result
}

//GetContactByID contact func
func GetContactByID(db *sql.DB, id int) Contact {
	var cont = Contact{}
	selectContact := `select * from phonebookdb.phonebook where id = ?`
	//вычисляем выражение sql
	stmt, err := db.Prepare(selectContact)
	if err != nil {
		fmt.Printf("%#v \n", stmt)
	}
	defer stmt.Close()
	//выполняем запрос к БД
	row := stmt.QueryRow(id)
	//записываем результат в модель
	errScan := row.Scan(
		&cont.ID,
		&cont.Firstname,
		&cont.Secondname,
		&cont.Sinonim,
		&cont.Prefix,
		&cont.Number,
		&cont.Note,
	)
	//обработка ошибок
	if errScan != nil {
		fmt.Printf("Не удалось получить контакт, ошибка %v \n", errScan)
	}
	//возвращаем модель
	return cont
}

//PutContact contact func
func PutContact(db *sql.DB, cont *Contact) (int64, error) {
	insertCont := `
	insert into phonebookdb.phonebook (Firstname, Secondname, Sinonim, Prefix, Number, Note)
	 values( $1, $2, $3, $4, $5, $6);
	 `
	//записываем результат в модель
	row, errExec := db.Exec(
		insertCont,
		cont.ID,
		cont.Firstname,
		cont.Secondname,
		cont.Sinonim,
		cont.Prefix,
		cont.Number,
		cont.Note,
	)

	//обработка ошибок
	if errExec != nil {
		fmt.Printf("Не удалось обновить контак %v. Ошибки: %v \n", cont.ID, errExec)
	}
	defer db.Close()
	//возврящаем id записи
	return row.LastInsertId()
}

//Delete contact func
func Delete(db *sql.DB, id int) (int64, error) {
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
