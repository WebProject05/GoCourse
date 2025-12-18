package main

import (
	"fmt"
	"time"
)

// These are async, they store data in them for a while
// SEND waits when full.
// RECEIVE waits when empty.

// func main() {
// 	// ===== Blocking on receive only if the buffer is empty =====
// 	ch := make(chan int, 2)
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		ch <- 498572
// 	}()
// 	fmt.Println("Value 1: ", <-ch)
// 	fmt.Println("Value 2: ", <-ch)
// }

func buffered() {
	// variable := make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 1000
	go func() {
		receiver := <-ch
		fmt.Println(receiver)
	}()
	go func() {
		receiver := <-ch
		fmt.Println(receiver)
	}()
	fmt.Println("Sent and received values")
	time.Sleep(1 * time.Second)
}
