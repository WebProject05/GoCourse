package main

import (
	"fmt"
)

func NonBlockingChannels() {
	ch := make(chan int)

	// Non blocking receive operation
	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received msg: ", msg)
	// default:
	// 	fmt.Println(" No Message avaliable...")
	// }

	// Non blocking send operation
	select {
	case ch <- 1:
		fmt.Println("Sent Value", <- ch)
	default:
		fmt.Println("Not ready cause there is no receiver...")
	}

	
}
