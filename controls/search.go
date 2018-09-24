package controls

import (
	"database/sql"
	"fmt"
	"log"
	"phonebook/models"

	ws "github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	upgrader = ws.Upgrader{}
	//FindeNumbers chan
	FindeNumbers chan []int
)

const (
	// заменить данные пользователь:пароль
	connect = "root:pass@/phonebookdb"
	driver  = "mysql"
)

//InitDB подключение к базе данных
func initDB(driver, connect string) *sql.DB {
	db, err := sql.Open(driver, connect)

	if err != nil || db == nil {
		log.Fatalf("Ошибка %v,\n при инициализации базы данных: %v", err, db.Stats())
	} else {
		log.Printf("Подключение к базе данных. '%v'\n", driver)
	}

	return db
}

//Search for search number
func Search(ctx echo.Context) error {
	var (
		numbers []int
	)

	wsc, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}
	defer wsc.Close()

	db := initDB(driver, connect)
	defer db.Close()
	models.SelectAllNumbers(db, FindeNumbers)

	for {
		numbers = <-FindeNumbers
		fmt.Printf("Цикл обработки %s", numbers)

		// Write
		err := wsc.WriteJSON(numbers)

		if err != nil {
			ctx.Logger().Error(err)
		}

		// Read
		_, msgByte, err := wsc.ReadMessage()

		if err != nil {
			ctx.Logger().Error(err)
		}
		fmt.Printf("%s\n", msgByte)
	}
}
