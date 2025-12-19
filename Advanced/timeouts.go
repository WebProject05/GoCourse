package main

import (
	"fmt"
	"time"
)

// Which ever happens first wins, it maybe goroutine or timeout (time.After())
// This stops out main code to wait for the gorouitine

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 69
		close(ch)
	}()


	// Can also use this, value, ok := <- ch:  and then if !ok do something
	select {
	case value:= <- ch:
		fmt.Println("Received:", value)
	case <- time.After(500 * time.Millisecond):
		fmt.Println("TimeOut!!!!")
		return   // killing the channel
	}
	time.Sleep(1 * time.Second)
	value2 := <- ch
	fmt.Println(value2)  // Cant access cause the goroutine is dead
}
