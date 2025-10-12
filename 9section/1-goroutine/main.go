package main

import (
	"fmt"
	"time"
)

func sayHello(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("sayHello", message)
}

func main() {

	fmt.Println("Hello from Main() Goroutine")

	go sayHello("Hello World 1", time.Second)
	go sayHello("Hello World 2", time.Second)
	go sayHello("Hello World from 2 seconds", 2*time.Second)
	go sayHello("Hello World from 3 seconds", 3*time.Second)

	fmt.Println("Last message from Main() Goroutine")

	time.Sleep(5 * time.Second)

}
