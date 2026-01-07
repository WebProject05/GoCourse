package main

import (
	"fmt"
	"time"
)

/*
	A worker pool is a design pattern commonly
	used in Go to manage a group of goroutines
	that process a queue of tasks.

	These help in processing task which are in queue
	in channels.
*/

//Worker function
/*
func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d \n", id, task)
		time.Sleep(time.Second) // Simulate work
		results <- task * 2
	}
}

func main() {
	numWorkers := 3
	numJobs := 10
	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := range numWorkers {
		go worker(i, tasks, results)
	}

	for i := range numJobs {
		tasks <- i
	}

	close(tasks)


	for range numJobs {
		result := <- results
		fmt.Println("Result:", result)
	}
}
*/

type ticketRequest struct {
	personId   int
	numTickets int
	cost       int
}


func ticketProcessing(requests <- chan ticketRequest, results chan <- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of personID %d with total cost %d \n", req.numTickets, req.personId, req.cost)
		time.Sleep(time.Second)
		results <- req.personId
	}
}


func WorkerPools() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	for range 3 {
		go ticketProcessing(ticketRequests, ticketResults) // Like giving access to the channels
	}

	for i := range numRequests {
		ticketRequests <- ticketRequest{personId: (i * 5) + 69, numTickets: (i * 2) ,cost: (i * price)}
	}
	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for personID %d processed successfully \n", <- ticketResults)
	}
}