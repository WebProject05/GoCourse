package main

import "fmt"

func main() {
	fmt.Println(factorial(10))
	fmt.Println(sumOfDigits(398095087))
}

func factorial(n int) int {
	// base case should be the main prioriyt (for my kn)
	if n == 0 {
		return 1
	}

	// recursive case
	return n * factorial(n-1)
}



func sumOfDigits(num int) int {
	if num < 10 {
		return num
	}

	return num % 10 + sumOfDigits(num / 10)
}