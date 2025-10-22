package main

import (
	"fmt"
	"time"
)

func main() {
	// current time
	currTime := time.Now()
	fmt.Println(currTime)

	// Specfic time
	specficTime := time.Date(2024, time.July, 29, 12, 0, 0, 0, time.UTC)
	fmt.Println("Specfic Time:", specficTime)

	// Parse 
	// Golang uses a reference time to parse time to a variable
	// The General form of the time is: Mon Jan 2 15:04:05 MST 2006
	parsedTime, _ := time.Parse("2006-01-02", "2025-10-22")	// The layout of the year - month - days should be same for the two even if the order is different
	parsedTime1, _ := time.Parse("01-2006-02", "09-2025-12")
	fmt.Println("Parsed times")
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)

	fmt.Println(parsedTime.Year())
	fmt.Println(parsedTime.Month())
	fmt.Println(parsedTime.Day())


	// Formating time
	t := time.Now()
	fmt.Println(t.Format("Monday 06-01-02 04-15"))
}