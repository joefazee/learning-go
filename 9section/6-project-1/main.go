package main

import (
	"context"
	"fmt"
	"time"
)

func ping(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- fmt.Sprintf("ping: %v", time.Now()):
			time.Sleep(1 * time.Second)
		}
	}
}

func pong(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- fmt.Sprintf("pong: %v", time.Now()):
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pingerCh := make(chan string)
	done := make(chan struct{})

	go ping(ctx, pingerCh)
	go pong(ctx, pingerCh)

	go func() {
		timeout := time.After(5 * time.Second)
		for {
			select {
			case <-timeout:
				fmt.Println("operation completed")
				close(pingerCh)
				done <- struct{}{}
				return
			case msg := <-pingerCh:
				fmt.Println(msg)
			}
		}
	}()

	<-done
	fmt.Println("done")

}
