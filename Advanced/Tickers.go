package main

import (
	"fmt"
	"time"
)

/*
	These tickers are used in making systems consistent
	and simple. Espically these are used in rate limiting
	for an API.
*/

// =======BASIC TICKERS======//
// func main() {
// 	/*
// 		The below line is for the time duration the
// 		ticker should tick for. Here in this case it is one second.
// 	*/
// 	ticker := time.NewTicker(2 * time.Second)
// 	defer ticker.Stop()

// 	i := 1
// 	for range 5 {
// 		fmt.Println(i)
// 		i += 1
// 	}

// 	// for tick := range ticker.C {
// 	// 	fmt.Println(tick)
// 	// }
// }

// ====SCHEDULING PERIODIC TASKS=====//
// func periodicTask() {
// 	fmt.Println("Performing periodic task:", time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(1 * time.Second)

// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

func main() {
	ticker := time.NewTicker(time.Second)
	stop := time.After(5 * time.Second)

	defer ticker.Stop()
	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick at:", tick)
		case <-stop:
			fmt.Println("Stopping ticker...")
			return
		}
	}
}
