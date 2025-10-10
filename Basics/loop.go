package main

import "fmt"

func main() {

	// Iteration over a range
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Iterate over collection
	numbers := []int{1,2,3,4,5,6}
	for index, val := range numbers {
		fmt.Printf("Index: %d, Value %d \n", index, val)
	}

	for i := range 10 {
		fmt.Println(i - 1)
	}
}