package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
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

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}

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

	users, err := GetUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

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

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	stmt := `SELECT id, name, email,  hashed_password, created_at FROM users WHERE email = ?`

	row := db.QueryRow(stmt, email)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func GetUsers(db *sql.DB) ([]User, error) {
	stmt := `SELECT id, name, email,  hashed_password, created_at FROM users`
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID,
			&user.Name,
			&user.Email,
			&user.HashedPassword,
			&user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
