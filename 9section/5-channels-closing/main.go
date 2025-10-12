package main

import (
	"fmt"
	"sync"
)

func main() {

	jobs := make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			r, ok := <-jobs
			if ok {
				fmt.Println("Got this message", r)
			} else {
				fmt.Println("Channel closed")
				return
			}
		}
	}(&wg)

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Sending ", i)
	}

	close(jobs)

	wg.Wait()

}

func doubleChannelConcept() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			r, ok := <-jobs
			if ok {
				fmt.Println("Got this message", r)
			} else {
				fmt.Println("Channel closed")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Sending ", i)
	}

	close(jobs)

	<-done
}
