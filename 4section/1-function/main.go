package main

import (
	"fmt"
)

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func calculateArea(width float64, height float64) float64 {
	if width < 0 || height < 0 {
		fmt.Println("Error: width and height must be positive")
		return 0
	}
	return width * height
}
func main() {
	greet("Bob Wonderland")
	add(1, 2)

	area := calculateArea(4, 4)
	fmt.Println(area)
}
