package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math: Square root of a negative number not defined in real plane!")
	}
	return 1, nil
}


func checkError(data []byte) error {
	if len(data) == 0 {
		return errors.New("The data does not exist.")
	}
	return nil
}

// Custom Errors
type myError struct {
	message string
}

func (er *myError) Error() string {
	return fmt.Sprintf("%s", er.message)
}

func errorCheck() error {
	return &myError{"Custom message error"}
}

func main() {
	result, err := sqrt(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)


	// data := []byte{}
	// err = checkError(data)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	error := errorCheck()
	if error != nil {
		fmt.Println("Custom Error: ", error)
		return
	}
}



