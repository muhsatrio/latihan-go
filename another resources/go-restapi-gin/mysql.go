package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connect() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
