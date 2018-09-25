package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"phonebook/controls"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "golang.org/x/net/websocket"
)

//DB database driver
var (
	DB *sql.DB
)

const (
	// заменить данные
	host = "localhost:8080"
	//пользователь:пароль
	connect     = "root:pass@/phonebookdb"
	connect2tcp = "root:pass@tcp(server_ip:port)/phonebookdb"
	driver      = "mysql"
)

//InitDB подключение к базе данных
func initDB(driver, connect string) *sql.DB {
	DB, err := sql.Open(driver, connect)

	if err != nil || DB == nil {
		log.Fatalf("Ошибка %v,\n при инициализации базы данных: %v", err)
	} else {
		log.Printf("Подключение к базе данных. '%v'\n", driver)
	}

	return DB
}

//Migrate миграция базы данных, создание таблиц
func migrate(db *sql.DB) {
	var err error

	usedb := `use phonebookdb;`
	createTablePhonebook := `create table phonebook	(
			id int primary key auto_increment not null,
			firstname varchar(20) not null,
			secondname varchar(40) not null,
			sinonim varchar(20),
			prefix varchar(3),
			number int not null,
			active bool
		) engine=innodb
		auto_icrement=1
		default charset=utf8mb4;
	`
	db.Exec(usedb)
	_, err = db.Exec(createTablePhonebook)

	result, err := testContact(db)

	if err != nil {
		log.Fatalf("Не удалось создать таблицу phonebook\n %v", err)
	} else {
		//Информация о базе данных
		id64, _ := result.RowsAffected()
		fmt.Printf("Тип подключения к БД: %#v\n", db.Stats())
		fmt.Printf("Успешно: таблица создана, количество строк %v \n", id64)
	}
}

func main() {
	//Инициализация базы данных
	db := initDB(driver, connect)
	migrate(db)
	//Создание приложения
	server := echo.New()
	server.Debug = true

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.HeaderUpgrade},
		}))

	//serve rout for phonebook
	server.File("/", "public/index.html")
	server.Static("/public", "public/static")
	server.GET("/contacts", controls.GetContacts(db))
	server.GET("/finde", controls.GetContact(db))
	server.PUT("/newсontact", controls.UpdateContact(db))
	server.POST("/contacts/:id", controls.CreateContact(db))
	server.DELETE("/delсontact/:id", controls.DeleteContact(db))
	//serve rout for websocket search
	server.GET("/search", controls.Search)

	//Запуск сервера с логированием
	fmt.Println("Успешно: север запущен")
	go server.Logger.Fatal(server.Start(host))

	//Остановка сервера
	defer server.Close()
	defer db.Close()
}

//test contact create and write to DB
func testContact(db *sql.DB) (sql.Result, error) {

	numberRnd := rand.Int63n(987654320) + int64(7000000000)

	result, err := db.Exec(`insert into phonebook
												 (firstname, secondname, sinonim, prefix, number, active )
													values ('Тестовый','пользователь', 'Ник', 'mts', ?, true)`, numberRnd)
	if err != nil {
		log.Println("Не удалось созадать запись в бд.")
	}
	return result, err
}
