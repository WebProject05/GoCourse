package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Definition: A WaitGroup allows you to wait for a
	collection of goroutines to complete their tasks.
	It maintains a counter that tracks the
	number of active goroutines.

	Usage: You use the Add(int) method to increment
	the counter when starting a new goroutine, the
	Done() method to decrement it when a goroutine
	finishes, and the Wait() method to block until
	the counter is zero, indicating all goroutines
	have completed their work.

	Common Scenario: It is particularly useful in scenarios
	where you need to ensure that all parallel tasks are
	completed before proceeding with subsequent operations,
	such as aggregating results or cleaning up resources.
*/
/* BASIC EXAMPLE WITHOUT CHANNELS
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting... \n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d finished... \n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorker := 3

	wg.Add(numWorker)  // Adding the counter to the wait group

	for i := range numWorker {
		go worker(i, &wg)
	}
	wg.Wait()  // Blocks from going to next processes

	fmt.Println("All the processes are completed")
}
*/

// With CHANNELS

func worker(id int, tasks <- chan int, results chan <- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WaitGroup %d is starting \n", id)
	time.Sleep(time.Second)
	for task := range tasks {
		results <- task
	}
	fmt.Printf("WaitGroup %d is completed \n", id)
}


func main() {
	var wg sync.WaitGroup
	numWorkers := 3
	numJobs := 10
	results := make(chan int, numJobs)
	tasks := make(chan int, numJobs)

	wg.Add(numWorkers)

	for i := range numWorkers {
		go worker(i + 1, tasks, results, &wg)
	}

	for i := range numJobs {
		tasks <- i
	}
	close(tasks)

	go func()  {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("result:", result)
	}
}