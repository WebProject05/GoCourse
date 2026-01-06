package main

import (
	"fmt"
	"time"
)

/*
- Timeouts
- Delays
- Peroidic Tasks
*/

/*
timer.Reset(some time) will reset the actual timer
*/

// func main() {
// 	timer := time.NewTimer(2 * time.Second) // This does not block
// 	stopped := timer.Stop()
// 	if stopped {
// 		fmt.Println("Timer Stopped...")
// 	} else {
// 		<-timer.C
// 		fmt.Println("Timer Expired...")
// 	}

// }


// func longOperation() {
// 	for i := range 5 {
// 		fmt.Println(i)
// 		time.Sleep(1 * time.Second)
// 	}
// }


// func main() {
// 	timeout := time.After(10 * time.Second)  // This is non Blocking
// 	done := make(chan bool)

// 	go func() {
// 		longOperation()
// 		done <- true	
// 	}()

// 	select {
// 	case <- timeout:
// 		fmt.Println("Operation Timeouted...")
// 	case <- done:
// 		fmt.Println("Operation Completed...")
// 	}
// }


// Scheduling delayed Operations
func Timers() {
	timer := time.NewTimer(2 * time.Second)
	go func() {
		<- timer.C
		fmt.Println("Delayed Operation Exceuted...")	
	}()
	fmt.Println("Waiting...")
	time.Sleep(3 * time.Second)
	fmt.Println("End of Operations...")
}