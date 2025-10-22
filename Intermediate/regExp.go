package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("He said, \"I am great\"")

	// Now regular expressions
	// compiling a regular expression to validate an email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z{2,}]`)   // all the lower and upper case alphabets in the mail including numbers
	// [a-zA-Z0-9._+%-]+ , in this block of regeexp, the + outside of the bracket is to inform the compiler that the characters inside of the block may occur multiple times
	// after the + we have the @gmail.com part of the regeexp

	email1 := "user@gmail.com"
	email2 := "prab@iiits.in"
	 
	fmt.Println("Email1: ", re.MatchString(email1))
	fmt.Println("Email2: ", re.MatchString(email2))

	// Capturing Groups
	// compile a regex pattern to capture date components
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	date := "2025-22-10"
	subMatches := re.FindStringSubmatch(date)
	fmt.Println(subMatches)
	fmt.Println(subMatches[0])
	fmt.Println(subMatches[1])
	fmt.Println(subMatches[2])
	fmt.Println(subMatches[3])

	// i - case insensitive
	// m - multi line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`)

	text := "Golang is going great"
	fmt.Println(re.MatchString(text))
}