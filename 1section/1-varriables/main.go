package main

import "fmt"

func main() {

	var greeting string // zero-value is an empty string ""
	greeting = "Hello, world!"

	fmt.Println(greeting)

	var count int
	count = 10
	fmt.Println(count)

	var isRunning bool
	isRunning = true
	fmt.Println(isRunning)

	var firstName, lastName string
	firstName = "John"
	lastName = "Doe"
	fmt.Println(firstName, lastName)

	email := "test@test.com"
	fmt.Println(email)

	age := 24
	fmt.Println(age)

	var year = 2025
	fmt.Println(year)

}
