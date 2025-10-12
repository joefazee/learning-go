package main

import (
	"fmt"
)

type Person interface {
	GetName() string
}

type BusinessPerson struct {
	ID   int
	Name string
}

func (e BusinessPerson) GetName() string {
	return e.Name
}

func (e BusinessPerson) String() string {
	return fmt.Sprintf("Person[ID:%d, Name:%s]", e.ID, e.Name)
}

type ID int

func (idx ID) String() string {
	return fmt.Sprintf("COMING FROM HERE ID: %d", idx)
}
func main() {

	jane := BusinessPerson{
		ID:   1,
		Name: "Jane",
	}

	fmt.Println(jane)

	var myId ID
	myId = 30
	fmt.Println(myId)

}
