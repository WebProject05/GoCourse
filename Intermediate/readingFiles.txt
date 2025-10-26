package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("writeString.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer func() {
		fmt.Println("Closing the currently opened file.")
		file.Close()
	}()

	fmt.Println("File is opened Now.")

	// Reading the content of the file
	// data := make([]byte, 1024)
	// _, err1 := file.Read(data)
	// if err1 != nil {
	// 	fmt.Println("Error:", err1)
	// 	return
	// }

	// fmt.Println("File Content:", string(data))


	// Scanning the file line by line using the bufio package
	scanner := bufio.NewScanner(file)

	// scanner.Scan() will return false when it encounters EOF in the file. Then the for loop terminates (actually while it is)
	for scanner.Scan() {	// This will loop through the lines of the file
		line := scanner.Text()
		fmt.Println("Line:", line)
	}
	
	err1 := scanner.Err()
	if err1 != nil {
		fmt.Println("Error")
		return
	}
}