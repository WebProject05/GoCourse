package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("writeString.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	// Keyword to filter lines
	keyword := "golang"

	// Read and filter lines
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			updatedLine := strings.ReplaceAll(line, keyword, "Go")
			fmt.Println("Updated Filtered line:", updatedLine)
		}
	}

	err1 := scanner.Err()
	if err1 != nil {
		fmt.Println("Error:", err1)
		return
	}
}