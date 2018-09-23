package main

import (
	"database/sql"
	"fmt"
	"log"
	"phonebook/controls"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "golang.org/x/net/websocket"
)

//DB database driver
var (
	db *sql.DB
)

const (
	// заменить данные пользователь:пароль
	connect     = "root:pass@/phonebookdb"
	connect2tcp = "root:pass@tcp(server_ip:port)/phonebookdb"
	driver      = "mysql"
)

//InitDB подключение к базе данных
func initDB(driver, connect string) *sql.DB {
	db, err := sql.Open(driver, connect)

	if err != nil || db == nil {
		log.Fatalf("Ошибка %v,\n при инициализации базы данных: %v", err)
	} else {
		log.Printf("Подключение к базе данных. '%v'\n", driver)
	}

	return db
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
		) engine=innodb auto_icrement=1 default charset=utf8mb4;
	`
	createTableChat := `create table chat(
			id int primary key auto_incrimant not nul,
			email varchar(40) not null,
			message varchar(128)
			) engine=innodb auto_icrement=1 default charset=utf8mb4;
		`

	db.Exec(usedb)
	_, err = db.Exec(createTablePhonebook)
	_, err = db.Exec(createTableChat)
	result, err := db.Exec("insert into phonebook (firstname, secondname, sinonim, prefix, number, active ) values ('Тестовый','пользо', 'ватель', '+7', '987654321', 'true')")

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

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		}))

	//serve rout for phonebook
	server.File("/", "public/index.html")
	server.GET("/contacts", controls.GetContacts(db))
	server.GET("/finde", controls.GetContact(db))
	server.PUT("/newсontact", controls.PutContact(db))
	server.POST("/contacts/:id", controls.PostContact(db))
	server.DELETE("/delсontact/:id", controls.DelContact(db))
	//serve rout for chat
	server.GET("/chatws", controls.ChatHandler())

	//Запуск сервера с логированием
	fmt.Println("Успешно: север запущен")
	go server.Logger.Fatal(server.Start(":8080"))

	//Остановка сервера
	defer server.Close()
	defer db.Close()
}
