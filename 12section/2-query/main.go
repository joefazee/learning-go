package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL, 
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

func main() {

	dbName := "users_database.db"

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("database connection established")

}

func createTable(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(db *sql.DB, name, email, hashedPassword string) (int64, error) {
	stmt := `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`

	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := db.Exec(stmt, name, email, string(hp))
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
