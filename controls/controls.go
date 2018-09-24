package controls

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"phonebook/models"
	"strconv"

	"github.com/labstack/echo"
)

var (
	ctx echo.Context
)

//JBODY любой JSON
type JBODY map[string]interface{}

//GetContacts handler
func GetContacts(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println(db.Ping())
		results := models.SelectAll(db)
		return ctx.JSON(http.StatusOK, results)
	}
}

//GetContact handler
func GetContact(db *sql.DB) echo.HandlerFunc {

	return func(ctx echo.Context) error {
		var (
			contact models.Contact
			err     error
		)

		fmt.Println(db.Ping())
		prefix := ctx.Param("prefix")
		number, err := strconv.Atoi(ctx.Param("number"))
		contact, err = models.SelectContact(db, number, prefix)

		if err != nil {
			ctx.Logger().Error(err)
		}

		return ctx.JSON(http.StatusOK, contact)
	}
}

//UpdateContact handler
func UpdateContact(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// создаём пустой контакт
		var (
			contact = models.Contact{}
			err     error
		)

		// привязываем пришедший JSON к новому контакту
		err = ctx.Bind(&contact)

		if err != nil {
			//Обработка ошибок
			ctx.Logger().Error(err)
		}

		id64, errDB := models.UpdateContact(db, &contact)
		if errDB != nil {
			return ctx.JSON(http.StatusOK, JBODY{
				"Контакт не обновлён №": id64,
			})
		}

		return ctx.JSON(http.StatusOK, JBODY{
			"Контакт обновлён №": id64,
		})
	}
}

//CreateContact handler
func CreateContact(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// создаём пустой контакт
		var contact = models.Contact{}
		// привязываем пришедший JSON к новому контакту
		ctx.Bind(&contact)

		//передаём в модель указатель на контакт
		id64, err := models.InsertContact(db, &contact)
		if err != nil {
			//Обработка ошибок
			log.Println(err)
			return ctx.JSON(http.StatusNotAcceptable, JBODY{
				"Контакт не создан №": id64,
			})
		}

		//при успехе вернём ответ статут OK, JSON
		return ctx.JSON(http.StatusOK, JBODY{
			"Контакт создан №": id64,
		})
	}
}

//DeleteContact handler
func DeleteContact(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, _ := strconv.Atoi(ctx.Param("id"))
		_, err := models.DeleteByID(db, id)

		if err == nil {
			return ctx.JSON(http.StatusOK, JBODY{
				"Контакт удалён №": id,
			})
		}
		//Обработка ошибок
		log.Println(err)

		return err
	}
}
