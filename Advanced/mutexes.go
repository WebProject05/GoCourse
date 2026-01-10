package main

import (
	"fmt"
	"sync"
)

/*
	It is a way of stopping multiple concurrent processes
	accessing shared resources simultaneously.
	These also help in previnting race conditions

	Lock() -> locks the data from other goroutines
	Unlock() -> Unlocks the data and allows other goroutines from accessing the data
	TryLock() -> It returns a boolean value and tries to access the mutex
*/

type counter struct {
	mu    sync.Mutex
	count int
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func mutexes() {
	var wg sync.WaitGroup
	counter := &counter{} // creating an instance of the struct and haveing the pointer pointing to it

	numGoroutines := 10
	wg.Add(numGoroutines)

	for range numGoroutines {
		// wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				// counter.count++   -> This statment will not perform correctly as multiple gorouines will have access to the shared value count
			}
		}()
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter.getValue())
}
