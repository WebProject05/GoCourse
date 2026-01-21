package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	A deadlock is a situation in concurrent computing where two or more processes or
	goroutines are unable to make progress because each
	is waiting for the other to release resources.
	In simpler terms, it's like a standstill where
	processes cannot proceed due to a resource
	blocking issue.
s
	To further elaborate, the four necessary conditions
	for a deadlock to occur are:

	Mutual Exclusion: At least one resource must be held in a non-shareable mode.
	Hold and Wait: A process holding at least one resource is waiting to acquire additional resources.
	No Preemption: Resources cannot be forcibly taken from a process; they must be voluntarily released.
	Circular Wait: There must be a set of processes in a circular chain where each process is waiting for
	a resource held by the next process in the chain.
*/

func deadLock() {
	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		fmt.Println("GoRoutine 1 locked mu1")
		time.Sleep(time.Second)
		mu2.Lock()  // Blocking statments
		fmt.Println("GoRoutine 1 locked mu2")
		mu1.Unlock()
		mu2.Unlock()
	}()

	go func() {
		mu2.Lock()
		fmt.Println("GoRoutine 2 locked mu2")
		time.Sleep(time.Second)
		mu1.Lock()
		fmt.Println("GoRoutine 2 locked mu1")
		mu1.Unlock()
		mu2.Unlock()
	}()

	time.Sleep(3 * time.Second)
}
