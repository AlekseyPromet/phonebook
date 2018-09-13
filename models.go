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

//AllContacts in phonebook
type AllContacts struct {
	Contacts []Contact `json:"contacts"`
}

//Read contact func
func (c Contact) Read(db *sql.DB) []string {
	selectContact := `SELET * FROM phonebook WHERE id=1;`

	contactRow, err := db.Query(selectContact)
	if err != nil {
		log.Fatal(err)
	}
	defer contactRow.Close()

	contact, err := contactRow.Columns()
	if err != nil {
		log.Fatal(err)
	}
	return contact
}

//Create contact func
func (c Contact) Create(db *sql.DB) (int, error) {
	return 0, nil
}

//Delete contact func
func (c Contact) Delete(db *sql.DB, id int) (int, error) {
	return 0, nil
}

//ReadAll select all rows in phonebook
func (contacts AllContacts) ReadAll(db *sql.DB) AllContacts {
	selectAllCont := `SELECT * FROM phonebook;`
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
