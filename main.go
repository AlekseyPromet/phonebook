package main

import (
	"database/sql"
	"fmt"
	"log"
	"phonebook/controls"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

//DB database driver
var DB *sql.DB

const (
	connect = "root:pass@/storage"
	driver  = "mysql"
)

//InitDB подключение к базе данных
func initDB(driver, connect string) *sql.DB {
	DB, err := sql.Open(driver, connect)

	if err != nil || DB == nil {
		log.Fatalf("Ошибка %v,\n при инициализации базы данных: %v", err)
	} else {
		fmt.Printf("Подключение к базе данных. '%v'\n", driver)
	}

	return DB
}

//Migrate миграция базы данных, создание таблиц
func migrate(db *sql.DB) {
	createdb := `CREATE DATEBASE IF NOT EXISTS phonebookdb;`
	usedb := `USE phonebookdb;`
	createtabl := ` CREATE TABLE	phonebook	(
			ID INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
			Firstname VARCHAR(20),
			Secondname VARCHAR(40),
			Sinonim VARCHAR(20),
			Prefix VARCHAR(3),
			Number INT,
			Note VARCHAR(120)
		);
	`
	_, err := db.Exec(createdb)
	_, err = db.Exec(usedb)
	_, err = db.Exec(createtabl)
	result, err := db.Exec("INSERT INTO phonebook (Firstname, Secondname, Sinonim, Prefix, Number, Note )VALUES ('Тест','товый', 'пользователь', '+7', '987654321', 'Заметка')")

	if err != nil {
		log.Fatalf("Не удалось создать таблицу phonebook\n %v", err)
	} else {
		//Информация о базе данных
		id64, _ := result.RowsAffected()
		fmt.Printf("%#v\n", db.Stats())
		fmt.Printf("Успешно: таблица создана, количество строк %v \n", id64)
	}
}

func init() {
	//Инициализация базы данных
	DB := initDB(driver, connect)
	migrate(DB)
}

func main() {
	//Создание приложения
	server := echo.New()
	//server.HidePort = false

	server.File("/", "public/index.html")
	server.GET("/contacts", controls.GetContacts(DB))
	server.PUT("/contacts", controls.PutContact(DB))
	server.DELETE("/contacts/:id", controls.DelContact(DB))
	//Запуск сервера с логированием
	fmt.Println("Успешно: север запущен")
	go server.Logger.Fatal(server.Start(":8080"))

	//Остановка сервера
	defer server.Close()
	defer DB.Close()
}
