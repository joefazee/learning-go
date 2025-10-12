package main

import (
	"fmt"
)

func main() {

	messages := make(chan string, 3)

	fmt.Println("Sending messages to buffered channel")
	messages <- "first message"
	messages <- "send message"
	messages <- "third message"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
