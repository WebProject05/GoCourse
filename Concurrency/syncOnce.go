package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("This will not be repeated no matter how many times we run this function")
}


func syncOnce() {
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Goroutine Number:", i)
			once.Do(initialize)   // This will call everytime the loop runs byt once.do will only execute once
		}()
	}

	wg.Wait()
}
