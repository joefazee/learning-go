package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string  `json:"name" xml:"name"`
	Age      int     `json:"age" xml:"age"`
	Phone    string  `json:"phone" xml:"phone_number"`
	Password string  `json:"-" xml:"-"`
	IsActive bool    `json:"active" xml:"active"`
	Role     string  `json:"-" xml:"role"`
	Profile  profile `json:"profile" xml:"profile"`
}

type profile struct {
	URL string `json:"url"`
}

var payload = `{
 "name": "Jane",
 "age": 20,
 "phone": "123-456-789",
 "active": true,
 "profile": {
  "url": "https://www.jane.co.id"
 }
}
`

func main() {

	var u user
	err := json.Unmarshal([]byte(payload), &u)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", u)

	bs, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}
