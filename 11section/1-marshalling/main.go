package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phone_number"`
	IsActive bool   `json:"active" xml:"active"`
	Role     string `json:"role" xml:"role"`
}

func main() {

	u := user{
		Name:     "John Doe",
		Age:      42,
		Phone:    "555-555-5555",
		IsActive: true,
	}
	bs, err := json.Marshal(u)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}
