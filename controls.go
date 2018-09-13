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

//GetContact test handler
func GetContact(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var contact Contact
		return c.JSON(http.StatusOK, contact.Read(db))
	}
}

//PutContact test handler
func PutContact(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var contact Contact
		//получаем номер контракта из контекста
		id, _ := strconv.Atoi(c.Param("id"))

		// привязываем пришедший JSON к новому контакту
		c.Bind(&contact)

		id, err := contact.Create(db)

		if err == nil {
			return c.JSON(http.StatusOK, JBODY{
				"Контакт обновлён №": id,
			})
		}
		//Обработка ошибок
		log.Fatal(err)

		return err
	}
}

//DelContact test handler
func DelContact(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var contact Contact

		id, _ := strconv.Atoi(c.Param("id"))

		_, err := contact.Delete(db, id)

		if err == nil {
			return c.JSON(http.StatusOK, JBODY{
				"Контакт удалён №": id,
			})
		}
		//Обработка ошибок
		log.Fatal(err)

		return err
	}
}
