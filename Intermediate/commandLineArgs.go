package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Command:", os.Args[0])	// This will print out the file path
	fmt.Println("Command:", os.Args[1])

	for i, arg := range os.Args {
		fmt.Printf("Argument %d: %s \n", i, arg)
	}

	// Defining the flags
	// for example the flap in "grep --help" is "--help". It will give all the required information about tools

	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "Santosh", "Name of the Coder")
	flag.IntVar(&age, "age", 20, "Age of the Coder")
	flag.BoolVar(&male, "male", true, "Gender of the Coder")

	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Male:", male)
}