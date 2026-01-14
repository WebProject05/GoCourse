package main

import (
	"fmt"
	"time"
)

type StateFull struct {
	count int
	ch    chan int
}

func (w *StateFull) start() {
	go func() {
		for {
			select {
			case val := <-w.ch:
				w.count += val
				fmt.Println("Current count:", w.count)
			}
		}
	}()
}

func (w *StateFull) send(val int) {
	w.ch <- val
}

func main() {
	stWorker := &StateFull{
		ch:    make(chan int),
		count: 0,
	}

	stWorker.start()

	for i := range 5 {
		stWorker.send(i)
		time.Sleep(500 * time.Millisecond)
	}
}
