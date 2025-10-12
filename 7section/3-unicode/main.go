package main

import (
	"fmt"
	"unicode"
)

func main() {

	data := []rune{'絵', '文', '字'}

	for _, v := range data {
		fmt.Println(string(v), unicode.IsLetter(v))
	}

}
