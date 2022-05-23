package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err.Error())
	}

	return db
}
