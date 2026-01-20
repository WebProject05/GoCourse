package main

import "fmt"




func printElements() {
	for i := range 5 {
		fmt.Println(i)
	}
}

func main() {
	printElements()
}