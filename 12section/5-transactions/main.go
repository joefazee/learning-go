package main

import (
	"context"
	"database/sql"
	//"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// Database transaction
//-----------------------

// 1. User creates account
// 2. Create a wallet for the user
// 3. Want to top up the wallet for the user
// 4. You want to write a transaction log

var schema = `
CREATE TABLE IF NOT EXISTS profile (
    user_id INTEGER PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
    avatar TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	Profile        Profile   `json:"profile"`
}

type Profile struct {
	UserID  int       `json:"user_id"`
	Avatar  string    `json:"avatar"`
	Created time.Time `json:"created"`
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

	_, err = createUser(db, "test with defer", "test@tets.s", "qkjsq", "kjqwkq")
	if err != nil {
		log.Fatal(err)
	}

}

func createTable(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

// Begin, Rollback or Commit
func createUser(db *sql.DB, name, email, hashedPassword, avatar string) (int64, error) {
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(name, email, string(hp))
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	profileStm, err := tx.PrepareContext(ctx, `INSERT INTO profiles (user_id, avatar) VALUES( ?, ?)`)
	if err != nil {
		return 0, err
	}

	defer profileStm.Close()
	_, err = profileStm.Exec(userID, avatar)
	if err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	stmt := `SELECT u.id, u.name, u.email,  u.hashed_password, u.created_at, p.avatar FROM users u 
	INNER JOIN profile p ON u.id = p.user_id WHERE u.email = ?`

	row := db.QueryRow(stmt, email)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt, &user.Profile.Avatar)
	if err != nil {
		return nil, err
	}
	user.Profile.UserID = user.ID
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
