package main

// IMPORTANT CONCEPT
/*
if we have two goroutines with channels inside them, the goroutine which will execute first will send its value to the receiver. So if the other goroutine is taking time to workout it wont get completed cause the other goroutine is must faster
*/

import (
	"fmt"
	"time"
)

// Multiplexing handles multiple concurrent operations in a single Goroutine

func multiSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 1000
	}()

	go func() {
		time.Sleep(500 * time.Millisecond) // This is faster than the first goroutine so ch1 will receive the value first and look instantly for the receiver
		ch1 <- 20202
	}()

	// If there is no sender but have a receiver, to prevent the DEADLOCK case we use select

	time.Sleep(2 * time.Second) // To prevent the default case completed, giving time for the gorouitnes to get finished
	
	// Looping to get the value from both goroutines
	for range 2 {
		select { // Select is like a switch case
		case msg := <-ch1: // This is a condition so the compiler wont wait for it execute.
			fmt.Println("Message: (ch1)", msg)
		case msg := <-ch2:
			fmt.Println("Message: (ch2)", msg)
		default: // This default case will prevent the deadlock error
			fmt.Println("No Channels ready...")
		}
	}
	fmt.Println("Message received!")
}
