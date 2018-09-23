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
		return ctx.JSON(http.StatusOK, models.SelectAll(db))
	}
}

//GetContact handler
func GetContact(db *sql.DB) echo.HandlerFunc {
	var findeNumber = make(chan []int, 10)
	var numbers []int
	return func(ctx echo.Context) error {
		fmt.Println(db.Ping())
		number, err := strconv.Atoi(ctx.Param("number"))
		if err != nil {
			ctx.Logger().Error(err)
		}
		go models.SelectAllNumbers(db, findeNumber, number)
		numbers = <-findeNumber
		for n := range numbers {
			fmt.Print(n)
		}
		fmt.Println()
		return ctx.JSON(http.StatusOK, numbers)
	}
}

func GetNumbers() {

}

//PutContact handler
func PutContact(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// создаём пустой контакт
		var contact = models.Contact{}
		// привязываем пришедший JSON к новому контакту
		ctx.Bind(&contact)

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			//Обработка ошибок
			ctx.Logger().Error(err)
		}

		id64, errDB := models.Update(db, &contact, id)
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

//PostContact handler
func PostContact(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// создаём пустой контакт
		var contact = models.Contact{}
		// привязываем пришедший JSON к новому контакту
		ctx.Bind(&contact)

		//передаём в модель указатель на контакт
		id64, err := models.Insert(db, &contact)
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

//DelContact handler
func DelContact(db *sql.DB) echo.HandlerFunc {
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
