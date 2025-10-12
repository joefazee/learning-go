package main

import (
	"fmt"
)

type Person interface {
	GetName() string
}

type Employee struct {
	ID   int
	Name string
}

type BusinessPerson struct {
	ID   int
	Name string
}

func (e BusinessPerson) GetName() string {
	return e.Name
}

func (e Employee) GetName() string {
	return e.Name
}

func displayPerson(p Person) {
	fmt.Println(p.GetName())
}

func main() {

	//joe := Employee{
	//	ID:   1,
	//	Name: "Joe",
	//}

	jane := BusinessPerson{
		ID:   1,
		Name: "Jane",
	}

	//displayPerson(joe)
	//displayPerson(jane)

	fmt.Println(jane)

}
