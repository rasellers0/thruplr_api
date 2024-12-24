package util

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func dbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:NovemberKazoo01!@tcp(127.0.0.1:3306)/thruplr")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	return db
}
