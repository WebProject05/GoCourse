package main

// There are basically key value pairs
// Generally used in API development to timeout nil returns and other error requests etc...
// One example can be, fetching a data from API exceeds a time limit we can use this context to store the timeouts and deadlines


/*
We have "withTimout where we pass a timeout deadline,
But in withCancel we dont have to pass a deadline period but just the context (ctx),
the syntax will look like
	ctx, cancel := context.withCancel(ctx)
	go func(){
		time.Seep(4 * time.Second)   Tis maybe some heavyduty task
		cancel()   After completion of the task cancel the context
	}()"
*/


import (
	"context"
	"fmt"
	"time"
)

func checkEvenOdd(ctx context.Context, num int) string {
	select {
	case <-ctx.Done():	// After 1 second a cancellation signal will be send to this and return whatever is inside of the case
		return "Operation Canceled..."
	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is Even", num)
		} else {
			return fmt.Sprintf("%d is Odd", num)
		}
	}
}

func ContextFunc() {   //main function actually
	ctx := context.TODO()

	result := checkEvenOdd(ctx, 5)
	fmt.Println("Result using context.TODO(): ", result)

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)   // After 1 second the context ctx will get cancelled, meaning the conext will send a deadline signal
	defer cancel()

	result = checkEvenOdd(ctx, 5)	// Extraction of a value from the context!
	fmt.Println("Result using the timeout context (background): ", result)

	time.Sleep(2 * time.Second)
	result = checkEvenOdd(ctx, 19)
	fmt.Println("Result after the timeout:", result)
}


// func main() {
// 	todoContext := context.TODO()
// 	contextBkg := context.Background()

// 	ctx := context.WithValue(todoContext, "name", "hihihi")
// 	fmt.Println(ctx.Value("name"))

// 	ctx1 := context.WithValue(contextBkg, "city", "Toronto")
// 	fmt.Println(ctx1.Value("city"))
// }
