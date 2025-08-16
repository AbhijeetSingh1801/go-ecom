package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
    "log"
)

var DB *sql.DB


func Connect() {
	var err error
	DB, err = sql.Open("sqlite3", "./ecom.db")
	if err != nil {
		log.Fatal("Failed to connect DB: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB unreachable: ", err)
	}
	log.Println("DB connected successfully")
}