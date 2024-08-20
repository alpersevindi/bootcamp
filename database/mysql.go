package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMySQL() (*sql.DB, error) {
	var err error

	dsn := "root:@tcp(127.0.0.1:3306)/bootcamp"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, nil
}
