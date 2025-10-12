package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

func NewEmployee(id int, firstName, lastName, position string, isActive bool) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		IsActive:  isActive,
		JoinedAt:  time.Now(),
	}
}

func main() {

	jane := Employee{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Position:  "Night",
		Salary:    1000,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}

	fmt.Println(jane.ID)
	fmt.Println(jane.FirstName)
	fmt.Println(jane.LastName)
	fmt.Println(jane.Position)
	fmt.Println(jane.Salary)

	joe := NewEmployee(1, "John", "Doe", "Jane", true)

	fmt.Println(joe.FirstName)
	fmt.Println(joe.LastName)
	fmt.Println(joe.Position)

	joe.Salary = 100000000
	fmt.Println(joe.Salary)

	joePtr := &joe
	joePtr.IsActive = true
	joePtr.LastName = "John Adam"
	fmt.Println(joe)

}
