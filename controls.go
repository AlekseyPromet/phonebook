package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//JBODY любой JSON
type JBODY map[string]interface{}

//GetContact test handler
func GetContact(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, ModelContact(db))
	}
}

//PutContact test handler
func PutContact(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, JBODY{
			"Контакт обновлён №": id,
		})
	}
}

//DelContact test handler
func DelContact(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, JBODY{
			"Контакт удалён №": id,
		})
	}
}
