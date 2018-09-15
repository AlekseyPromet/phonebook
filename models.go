package main

import (
	"database/sql"
	"log"
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

//Migrate миграция базы данных, создание таблиц
func Migrate(db *sql.DB) {
	create := `
		CREATE TABLE IF NOT EXISTS phonebook(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			firstname VARCHAR NOT NULL,
			secondname VARCHAR,
			sinonim VARCHAR,
			prefix VARCHAR,
			phone INTEGER,
			note VARCHAR
		);
	`
	_, err := db.Exec(create)
	if err != nil {
		log.Fatalf("Не удалось создать таблицу phonebook\n %v", err)
	}
}

//ReadNumber contact func
func ReadNumber(db *sql.DB, id int) Contacts {
	var contacts Contacts
	selectContact := `SELET * FROM phonebook WHERE number LIKE ?% `

	sql, err := db.Prepare(selectContact)
	if err != nil {
		log.Fatal(err)
	}
	defer sql.Close()

	contactRows, errRead := sql.Query(id)
	if errRead != nil {
		log.Println(errRead)
	}

	for contactRows.Next() {
		var cont Contact

		errCont := contactRows.Scan(
			&cont.ID,
			&cont.Firstname,
			&cont.Secondname,
			&cont.Sinonim,
			&cont.Prefix,
			&cont.Number,
			&cont.Note,
		)

		if errCont != nil {
			log.Printf("Не удалось прочитать контакт, ошибка %v \n", errCont)
		}

		contacts.Contacts = append(contacts.Contacts, cont)
	}

	return contacts
}

//Create contact func
func Create(db *sql.DB) (int64, error) {
	var contact Contact
	createCont := `INSERT INTO phonebook
								VALUES(?, ?, ?, ?, ?, ?)`

	sql, err := db.Prepare(createCont)
	if err != nil {
		log.Fatal(err)
	}
	defer sql.Close()

	result, errCre := sql.Exec(
		contact.ID,
		contact.Firstname,
		contact.Secondname,
		contact.Sinonim,
		contact.Prefix,
		contact.Number,
		contact.Note,
	)

	if errCre != nil {
		log.Fatal(err)
	}

	return result.LastInsertId()
}

//Delete contact func
func Delete(db *sql.DB, id int) (int64, error) {
	delCont := "DELETE FROM tasks WHERE id = ?"

	// выполним SQL запрос
	sql, err := db.Prepare(delCont)
	// выход при ошибке
	if err != nil {
		log.Fatal(err)
	}

	// заменим символ '?' в запросе на 'id'
	result, errDel := sql.Exec(id)
	// выход при ошибке
	if errDel != nil {
		log.Fatal(errDel)
	}

	return result.RowsAffected()
}

//ReadAll select all rows in phonebook
func ReadAll(db *sql.DB) Contacts {
	var contacts Contacts
	selectAllCont := `SELECT * FROM phonebook`
	contactRows, err := db.Query(selectAllCont)
	if err != nil {
		log.Fatal(err)
	}

	defer contactRows.Close()

	for contactRows.Next() {
		var newContact Contact

		errCont := contactRows.Scan(
			&newContact.ID,
			&newContact.Firstname,
			&newContact.Secondname,
			&newContact.Sinonim,
			&newContact.Prefix,
			&newContact.Number,
			&newContact.Note,
		)

		if errCont != nil {
			log.Printf("Не удалось прочитать контакт, ошибка %v \n", errCont)
		}

		contacts.Contacts = append(contacts.Contacts, newContact)
	}

	return contacts
}
