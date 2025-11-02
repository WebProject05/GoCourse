package main

import (
	"fmt"
	"time"
)

// These are simple Threads which handle multiple tasks / concurrent programming. It handles tasks concurrently without manually managing threads.
// We have to use go keyword to start a new goroutine.
// Goroutines are functions that leave the main thread and run in the background and come back to join the main thread once the functions are finished / ready to return any value
// Goroutines are non stopping functions and they dont block any work in the main thread.

func sayHello() {
	fmt.Println("Hello")
}

func printNumbers() {
	for i := 0; i < 11; i++ {
		// fmt.Println("Number:", i)
		fmt.Println(time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "hello" {
		fmt.Println("Char:", string(letter), " time:", time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)

	return fmt.Errorf("Error...")
}

func main() {
	var err error

	go func() {
		err = doWork()
	}()

	go printNumbers()
	go printLetters()

	time.Sleep(5 * time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work Completed!")
	}

}
