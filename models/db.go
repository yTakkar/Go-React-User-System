package models

import (
	"database/sql"

	// for mysql
	_ "github.com/go-sql-driver/mysql"
)

// DB function
func DB() *sql.DB {
	db, _ := sql.Open("mysql", "root:iamaprogrammer@tcp(127.0.0.1:3306)/go-notes-app")
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
