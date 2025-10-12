package main

import (
	"embed"
	"fmt"
	"log"
)

// enterprise application in Go
// ----------------------
var name = "Joseph"

//go:embed public
var public embed.FS

func main() {

	data, err := public.ReadFile("public/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	fmt.Println(name)

}
