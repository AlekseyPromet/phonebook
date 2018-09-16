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
	// заменить данные пользователь:пароль
	connect = "root:pass@/phonebookdb"
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
	usedb := `use phonebookdb;`
	createtabl := ` create table phonebookdb.phonebook	(
			id int primary key auto_increment not null,
			firstname varchar(20),
			secondname varchar(40),
			sinonim varchar(20),
			prefix varchar(3),
			number int,
			note varchar(120)
		);
	`
	_, err := db.Exec(usedb)
	_, err = db.Exec(createtabl)
	result, err := db.Exec("insert into phonebookdb.phonebook (Firstname, Secondname, Sinonim, Prefix, Number, Note ) values ('Тест','товый', 'пользователь', '+7', '987654321', 'Заметка')")

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
	db := initDB(driver, connect)
	migrate(db)
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
