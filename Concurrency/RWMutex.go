package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Multiple GoRoutines can read data at a time but the write method is locked for one GoRoutine
	Read Lock: RLock()
	Write Lock: Lock()
*/

var (
	rwmu      sync.RWMutex
	rwCounter int
)

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	rwmu.RLock()
	fmt.Println("Read Counter:", rwCounter)
	rwmu.RUnlock()
}

func writeCounter(wg *sync.WaitGroup, val int) {
	defer wg.Done()
	rwmu.Lock()
	rwCounter = val
	fmt.Println("Written value", val)
	rwmu.Unlock()
}

func RWMutex() {
	var wg sync.WaitGroup
	for range 5 {
		wg.Add(1)
		go readCounter(&wg)
	}

	wg.Add(1)
	time.Sleep(1 * time.Millisecond)
	go writeCounter(&wg, 18)

	wg.Wait()
}
