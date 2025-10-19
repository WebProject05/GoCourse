package main

import (
	"errors"
	"fmt"
)

type customError struct {
	code int
	msg  string
	er   error
}

// Error returns the error message
func (error *customError) Error() string { // The "Error() string" is builtin in golang so we dont need to call the function
	return fmt.Sprintf("Error %d: %s, %v", error.code, error.msg, error.er)
}

// A function to return an error
// func createError() error {
// 	return &customError{
// 		code: 404,
// 		msg:  "Not found! Try again later",
// 	}
// }

func createError() error {
	err := newError()
	if err != nil {
		return &customError{
			code: 404,
			msg:  "Not Found",
			er:   err,
		}
	}
	return nil
}

func newError() error {
	return errors.New("Internal Error. Comeback later")
}


func main() {
	err := createError()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Beyond the error checking block.")
}
