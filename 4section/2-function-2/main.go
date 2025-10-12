package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	fmt.Println(factorial(5))

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	logger := func(msg string) {
		fmt.Println(msg)
	}

	logger("Hello World")
	logger("Hello World")

}
