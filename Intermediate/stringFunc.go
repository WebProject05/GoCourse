package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// String Conversion
	num := 12
	str := strconv.Itoa(num)
	fmt.Printf("%T \n", str)

	// String splitting
	fruits := "Hello, My, I, Lie"
	parts := strings.Split(fruits, ", ")
	fmt.Println(parts)

	// Array / Slice of strings to a string
	countries := []string{"India", "aidnI"}
	joined := strings.Join(countries, " - reversed - ")
	fmt.Println(joined)

	// Containes or not
	fmt.Println(strings.Contains(str, "1"))

	// Replaceing
	replacd := strings.Replace(str, "1", "2", 1)
	fmt.Println(replacd)

	// Trimming spaces off the strings
	original := " Hello !    "
	fmt.Println("Original with no trimming", original)
	fmt.Println("Trimmed string with trimmed spaces at the left and the right corner",strings.TrimSpace(original))
}