# phonebook

1. Установить Go версию не ниже 1.10
 внешние зависимости: 	_ "github.com/go-sql-driver/mysql"
	                      "github.com/labstack/echo"
 установить: $ go get github.com/go-sql-driver/mysql && go get github.com/labstack/echo
                     
2. Установить MySQL или MariaDB
3. Создать БД
            $ mysql -u root -p
            MariaDB [(none)]> create database phonebookdb;
4. В main.go заменить данные пользователь:пароль
	                          connect = "root:pass@/phonebookdb" 
 5. Запустить $ go run main.go
