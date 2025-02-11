package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `CREATE TABLE staking (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		wallet_address TEXT UNIQUE,
		amount REAL
	);`

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}
