package main

import (
	"fmt"
	"os"
)

func simpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	defer fmt.Println("Function simpleDefer: deferred")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")

}

func lifoSimpleDefer() {
	fmt.Println("Function lifoSimpleDefer: Start")
	defer fmt.Println("First: deferred")
	defer fmt.Println("Second: deferred")
	fmt.Println("Function lifoSimpleDefer: Middle")
}
func main() {
	file, err := os.Create("./defer.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//simpleDefer()
	lifoSimpleDefer()

	fmt.Println("Last in main()")

}
