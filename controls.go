package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//JBODY любой JSON
type JBODY map[string]interface{}

//GetContacts test handler
func GetContacts(db *sql.DB) echo.HandlerFunc {
	var contacts = Contacts{}
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, contacts.ReadAll(db))
	}
}

//GetContact test handler
func GetContact(db *sql.DB) echo.HandlerFunc {
	var contacts = Contacts{}
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, contacts.ReadNumber(db, id))
	}
}

//PutContact test handler
func PutContact(db *sql.DB) echo.HandlerFunc {
	var contact = Contact{}

	return func(ctx echo.Context) error {
		//получаем номер контракта из контекста
		//id, _ := strconv.Atoi(c.Param("id"))
		//contact.ID = id
		// привязываем пришедший JSON к новому контакту
		ctx.Bind(&contact)

		id64, err := contact.Create(db)

		if err == nil {
			return ctx.JSON(http.StatusOK, JBODY{
				"Контакт обновлён №": id64,
			})
		}
		//Обработка ошибок
		log.Fatal(err)

		return err
	}
}

//DelContact test handler
func DelContact(db *sql.DB) echo.HandlerFunc {
	var contact = Contact{}
	return func(ctx echo.Context) error {

		id, _ := strconv.Atoi(ctx.Param("id"))

		_, err := contact.Delete(db, id)

		if err == nil {
			return ctx.JSON(http.StatusOK, JBODY{
				"Контакт удалён №": id,
			})
		}
		//Обработка ошибок
		log.Fatal(err)

		return err
	}
}
