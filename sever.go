package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	//Инициализация базы данных
	db := initDB("./db/storage.db")
	Migrate(db)
	defer db.Close()
}

func main() {
	//Создание приложения
	server := echo.New()
	server.HidePort = false

	server.File("/", "public/index.html")
	server.GET("/phonebook", GetContacts(db))
	server.PUT("/phonebook", PutContact(db))
	server.DELETE("/phonebook/:id", DelContact(db))

	//Запуск сервера с логированием
	server.Logger.Fatal(server.Start(":8080"))
	//Остановка сервера
	defer server.Close()
}

//подключение к базе данных
func initDB(DBPath string) *sql.DB {
	db, err := sql.Open("sqlite3", DBPath)

	if err != nil || db == nil {
		log.Fatalf("Ошибка %v,\n при инициализации базы данных, путь: %v", err, DBPath)
	}

	return db
}
