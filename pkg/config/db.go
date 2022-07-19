package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {

	var err error

	DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/author_book_db")
	if err != nil {
		panic(err)
	}
	fmt.Println("connect success")
}
