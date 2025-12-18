package main

import (
	"fmt"
	"time"
)

// Ensures that the data is exchanged between the Goroutines properly

// func main() {
// 	done := make(chan struct{})

// 	// Sending struct to the channel via goroutine
// 	go func () {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		done <- struct{}{}
// 	}()

// 	input := <- done
// 	fmt.Println(input)
// 	fmt.Println("Finished...")
// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		fmt.Println("Sending the data...")
// 		ch <- 10   // blocking until a value is sent to the channel
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("Sent value...")	// Wont get executed cause the channel has received the value so it will look for a receiver
// 	}()

// 	value := <- ch   // Blocking until a value is recived from the channel
// 	fmt.Println(value)
// }

// Sync Multiple Goroutines and ensuring that all the goroutines are completed
// func main() {
// 	numGoroutines := 5
// 	ch := make(chan int, 5)

// 	for i := range numGoroutines {
// 		time.Sleep(2 * time.Second)
// 		go func (id int)  {
// 			fmt.Printf("Goroutine %d is working \n", id)
// 			time.Sleep(1 * time.Second)
// 			ch <- id
// 		}(i)
// 	}

// 	// In the for block if the range is upto 4 then the 5th goroutine wont send the data so the value will be nil and the goroutine does not exist

// 	for range numGoroutines - 1 {
// 		value := <- ch  // Wait for each goroutine to finish
// 		fmt.Println(value)
// 		fmt.Println("Received from channel")
// 	}
// }

func main() {
	ch := make(chan string)

	go func() {
		for i := range 5 {
			ch <- "hello " + string('0'+i)
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()

	// The below for loop will result in an error cause there are no channel receivers to create a channel sendder so we use close() method
	for value := range ch {
		fmt.Println("Received the value", value, ":", time.Now())
	}
}
