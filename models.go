package main

import (
	"database/sql"
	"log"
)

//Contact in phonebook
type Contact struct {
	ID         int    `json:"id"`
	Firstname  string `json:"fistname"`
	Secondname string `json:"secondname"`
	Sinonim    string `json:"sininime"`
	Prefix     string `json:"prefix"`
	Number     int    `json:"number"`
}

//Contacts in phonebook
type Contacts struct {
	Contacts []Contact `json:"contacts"`
}

//ReadNumber contact func
func (contacts Contacts) ReadNumber(db *sql.DB, id int) Contacts {
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
		)

		if errCont != nil {
			log.Printf("Не удалось прочитать контакт, ошибка %v \n", errCont)
		}

		contacts.Contacts = append(contacts.Contacts, cont)
	}

	return contacts
}

//Create contact func
func (cont Contact) Create(db *sql.DB) (int64, error) {
	createCont := `INSERT INTO phonebook
								VALUES(?, ?, ?, ?, ?, ?)`

	sql, err := db.Prepare(createCont)
	if err != nil {
		log.Fatal(err)
	}
	defer sql.Close()

	result, errCre := sql.Exec(
		cont.ID,
		cont.Firstname,
		cont.Secondname,
		cont.Sinonim,
		cont.Prefix,
		cont.Number,
	)

	if errCre != nil {
		log.Fatal(err)
	}

	return result.LastInsertId()
}

//Delete contact func
func (cont Contact) Delete(db *sql.DB, id int) (int64, error) {
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
func (contacts Contacts) ReadAll(db *sql.DB) Contacts {
	selectAllCont := `SELECT * FROM phonebook`
	contactRows, err := db.Query(selectAllCont)
	if err != nil {
		log.Fatal(err)
	}

	defer contactRows.Close()

	for contactRows.Next() {
		var cont Contact

		errCont := contactRows.Scan(
			&cont.ID,
			&cont.Firstname,
			&cont.Secondname,
			&cont.Sinonim,
			&cont.Prefix,
			&cont.Number,
		)

		if errCont != nil {
			log.Printf("Не удалось прочитать контакт, ошибка %v \n", errCont)
		}

		contacts.Contacts = append(contacts.Contacts, cont)
	}

	return contacts
}
