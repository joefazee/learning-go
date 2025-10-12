package main

import (
	"database/sql"
	"learning-go/12section/6-repository/repository"

	//"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {

	dbName := "users_database.db"
	db, err := connectToDatabase(dbName)
	checkErr(err)

	fmt.Println("database connection established")

	repo := repository.NewSQLUserRepository(db)

	printUsers(repo)

}

func printUsers(repo repository.UserRepository) {
	users, err := repo.GetUsers()
	checkErr(err)
	for _, user := range users {
		fmt.Println(user.ID, user.Email)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func connectToDatabase(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
