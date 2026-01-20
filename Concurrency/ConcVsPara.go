package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// func printElements() {
// 	for i := range 5 {
// 		fmt.Println(i)
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

// func printLetters() {
// 	for _, val := range "ABCED" {
// 		fmt.Println(string(val))
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

func heavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d is starting \n", id)
	for range 100_000_000 {

	}
	fmt.Println((time.Now()))
	fmt.Printf("Task %d is completed \n", id)
}

func cvp() {
	// go printElements()
	// go printLetters()

	// time.Sleep(3 * time.Second)
	numThreads := 4

	runtime.GOMAXPROCS(numThreads)
	var wg sync.WaitGroup
	for i := range numThreads {
		wg.Add(1)
		go heavyTask(i, &wg)
	}

	wg.Wait()
}
