package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a buffered reader from a string source
	// reader := bufio.NewReader(strings.NewReader("Hello hello \n"))

	// // Step 1: Read a few bytes (not the whole data)
	// data := make([]byte, 5) // smaller slice so some data remains in buffer
	// n, err := reader.Read(data)
	// if err != nil {
	//   fmt.Println("Error reading bytes:", err)
	//   return
	// }
	// fmt.Printf("Read %d Bytes, data: %q\n", n, data[:n])

	// // Step 2: Read the remaining data until newline '\n'
	// str, err := reader.ReadString('\n')
	// if err != nil {
	//   fmt.Println("Error reading string:", err)
	//   return
	// }
	// fmt.Printf("Read string till newline: %q\n", str)

	writer := bufio.NewWriter(os.Stdout)
	// Writing a byte slice
	data := []byte("Hello, this is via bufio new writer \n")
	numBytes, err := writer.Write(data) // Write a new data into the buffer
	if err != nil {
		fmt.Println("Error Writing:", err)
	}
	fmt.Println("Number of bytes:", numBytes)
  flushErr := writer.Flush()   // Prints the data, Flush the buffer to ensure all the data is written to os.Stdout
  if flushErr != nil {
    fmt.Println(flushErr)
    return
  }

  // Writing string
  str := "Helloooooooo"
  n, err := writer.WriteString(str)
  if err != nil {
    fmt.Println("Error:", err)
  }
  fmt.Println("Number of bytes:", n)
  err1 := writer.Flush()  // Writing the data which is inside the buffer
  if err1 != nil {
    fmt.Println(err1)
  }
}
