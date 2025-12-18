package main

import "fmt"

// Channels are used for communication between concurrent Goroutines
// Helps synchronize and manage the flow of data in concurrent programs

func channels() {
	//variable := make(chan type), '<-' recive operator
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString //Sending the value of greetString to greeting channel
		// fmt.Println(greeting)  This is a channel address
		greeting <- "Hello" // To read this value we have to read again
	}()

	receiver := <-greeting // receiver is receving value from the greeting channel
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)
	fmt.Println("End of end")
}