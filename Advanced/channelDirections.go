package main

import "fmt"

// These are indentend to use in functions and goroutines

func channelDirections() {
	ch := make(chan int)
	producer(ch)

	consumer(ch)
}

// Producer only channel (send only channel)
func producer(ch chan<- int) {
	go func() { // The function will allow only the send only to the channel
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// Receive only channel
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received: ", value)
	}
}
