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

//JBODY любой JSON
type JBODY map[string]interface{}

//GetContacts handler
func GetContacts(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Получаем JSON и отдаём клиенту")
		return c.JSON(http.StatusOK, models.GetAllContacts(db))
	}
}

//GetContactByID handler
func GetContactByID(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// извлекаем id контакта из параметров переданных клиентом
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, models.GetContactByID(db, id))
	}
}

//PutContact handler
func PutContact(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// создаём пустой контакт
		var contact = models.Contact{}
		// привязываем пришедший JSON к новому контакту
		ctx.Bind(&contact)

		//передаём в модель указатель на контакт
		id64, err := models.PutContact(db, &contact)
		if err == nil {
			//при успехе вернём ответ статут OK, JSON
			return ctx.JSON(http.StatusOK, JBODY{
				"Контакт обновлён №": id64,
			})
		}
		//Обработка ошибок
		log.Println(err)
		return ctx.JSON(http.StatusNotAcceptable, JBODY{
			"Контакт не обновлён №": id64,
		})
	}
}

//DelContact handler
func DelContact(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, _ := strconv.Atoi(ctx.Param("id"))
		_, err := models.Delete(db, id)

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
