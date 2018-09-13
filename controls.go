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
		return c.JSON(http.StatusOK, ModelGet(db))
	}
}

//PutContact test handler
func PutContact(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		var contact Contact
		// привязываем пришедший JSON к новому контакту
		c.Bind(&contact)

		id, err := ModelPutContact(db, contact)

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
		id, _ := strconv.Atoi(c.Param("id"))

		_, err := ModelDelContact(db, id)

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
