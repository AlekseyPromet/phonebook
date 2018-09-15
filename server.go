package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	//Создание приложения
	server := echo.New()
	server.HidePort = false

	//Инициализация базы данных
	db := initDB("./db/storagserver.db")
	migrate(db)
	defer db.Close()

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

//миграция базы данных, создание таблиц
func migrate(db *sql.DB) {
	create := `
		CREATE TABLE IF NOT EXISTS phonebook(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			firstname VARCHAR NOT NULL,
			secondname VARCHAR,
			sinonim VARCHAR,
			prefix VARCHAR,
			phone INTEGER
		);
	`
	_, err := db.Exec(create)
	if err != nil {
		log.Fatalf("Не удалось создать таблицу phonebook\n %v", err)
	}
}
