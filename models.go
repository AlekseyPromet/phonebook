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

//ModelGet func
func ModelGet(db *sql.DB) []string {
	selectGetContact := `SELET * FROM phonebook WHERE id=1;`

	contactRow, err := db.Query(selectGetContact)
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

//ModelPutContact func
func ModelPutContact(db *sql.DB, contact Contact) (int, error) {
	return 0, nil
}

//ModelDelContact func
func ModelDelContact(db *sql.DB, id int) (int, error) {
	return 0, nil
}
