package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	// Example usage:
	q, r := divide(10, 3)
	println("Quotient:", q, "Remainder:", r)

	str, err := compare(2,2)
	fmt.Printf("Value: %v \n", str)
	fmt.Printf("Value: %v \n", err)
}

// In the return decleration part of the function we can also do like (value string, err error)
func compare(a, b int) (string, error) {
	if a > b {	//Golang does not have ternary operators
		return "a > b", nil
	} else if a < b {
		return "a < b", nil
	} else {
		return "", errors.New("Error in comparing")
	}
}