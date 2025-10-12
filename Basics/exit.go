package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Deffered")
	fmt.Println("Stating the main function")
	os.Exit(1)

	// Below the exit no code will be exceuted
	fmt.Println("After the exit")
	// Even defer function does not execute
}