package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "hello \nmate"
	message2 := "hello \tmate"

	fmt.Println(message)
	fmt.Println(message2)


	fmt.Println("The lenght of the message is", len(message))
	fmt.Println("The lenght of the message is", len(message2))

	fmt.Println("FIrst character in the message ",message[0])   // ASCII

	first := "San"
	second := "tosh"

	full := first + second
	fmt.Println(full)

	for i, char := range message {
		fmt.Printf("%d %c\n", i, char)
	}

	fmt.Println("Rune count: ", utf8.RuneCountInString(full))

	var cha rune = '₹'
	fmt.Println(cha)
	fmt.Printf("\n %c", cha)

	convertedStr := string(cha)
	fmt.Printf("Type of converted: %T \n", convertedStr)
	fmt.Printf("Type of converted: %T \n", cha)

	helloT := "హలో"
	fmt.Println(helloT)
	for _, runeValue := range helloT {
		fmt.Printf("%c \n", runeValue)
	}
} 