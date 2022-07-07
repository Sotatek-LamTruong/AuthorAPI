package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {

	var err error

	DB, err = sql.Open("mysql", "root:Lucian2001@@tcp(database.dev:3306)/BookDB")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
