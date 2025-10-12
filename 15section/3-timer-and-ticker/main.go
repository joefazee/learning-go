package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Hour)
	counter := 0
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			counter++
			fmt.Println("Tick")
			if counter >= 5 {
				fmt.Println("stopped")
				return
			}
		}
	}
}

func timerExample() {

	timer := time.NewTimer(5 * time.Second)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-timer.C
		fmt.Println("After 1 second")
	}()

	fmt.Println("This is happening inside the main goroutine")

	wg.Wait()

	fmt.Println("program ends")
}
