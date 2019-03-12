package services

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var (
	Db *sql.DB
)

func InitiateDatabase() {
	dbCon, err := sql.Open("sqlite3", "/Users/calebmoore/GoLang/src/PythonUpdate/475.db")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Connected to database")
	}

	Db = dbCon
}
