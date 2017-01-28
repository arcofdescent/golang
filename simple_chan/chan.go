package main

import "fmt"
import "time"

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 2) // sleep for 2 seconds
		ch <- "Hello, chan"
	}()

	msg := <-ch
	fmt.Printf("Received from channel: %s\n", msg)
}
