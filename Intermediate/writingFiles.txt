package main

import (
	"fmt"
	"os"
)

// OS package provides all the operations to handle the file operations


func main() {
	file, err := os.Create("output.txt")   // Creating a pointer to the file
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()   // this will be exceuted right at the end of the function

	// Writing the data into the file
	data := []byte("Hello world mate \n")
	_, err1 := file.Write(data)   // this will return the number of lines it has written and error
	if err1 != nil {
		fmt.Println("Error:", err1)
	}
	fmt.Println("Data has been written!")

	file1, err2 := os.Create("writeString.txt")
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	defer file1.Close()

	_, err3 := file1.WriteString("Hello mate! How is it going.? \n")
	if err3 != nil {
		fmt.Println("Error:", err3)
		return
	}
	fmt.Println("Data has been written successfully!")
}