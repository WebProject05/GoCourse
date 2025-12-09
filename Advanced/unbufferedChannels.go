package main

import (
	"fmt"
	"time"
)

// allows the channel to hold some values in the buffer before logging into the receiver
// helps in the rate of data transfer between the consumer and the maker

// Receiver will wait for all the goroutines to complete and then end


func unbuffered() {
	ch := make(chan int)
	go func() {   // This goroutine is already ready to receive the value from and the channel
		receiver := <-ch  // This will block the code until it recives the value or the main thread ends
		fmt.Println(receiver)
		time.Sleep(1 * time.Second)  // Waiting for the go Runtime to start the goroutine
	}()
	ch <- 11  // As soon as this channel receives the data, it looks for the reciver which is already existent in the other thread
	// time.Sleep(2 * time.Second)
}
