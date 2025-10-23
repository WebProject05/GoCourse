package main

import (
	"fmt"
	"strconv"
)

func main() {
	numString := "12345"
	number, err := strconv.Atoi(numString)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("The converted number:", number)

	// Parsing the number
	paredInt, err1 := strconv.ParseInt(numString, 10, 32)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(paredInt)
	fmt.Printf("%T \n", paredInt)


	flaotStr := "3.142325"
	floatNum, err2 := strconv.ParseFloat(flaotStr, 32)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(floatNum)
	fmt.Printf("%T \n", floatNum)

	binaryStr := "1010"
	decimal, err3 := strconv.ParseInt(binaryStr, 2, 64)
	if err3 != nil {
		fmt.Println("Error:", err3)
	}
	fmt.Println(decimal)

	hexStr := "1010"
	hexaDec, err4 := strconv.ParseInt(hexStr, 16, 64)
	if err3 != nil {
		fmt.Println("Error:", err4)
	}
	fmt.Println(hexaDec)
}