package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"age"`
	// Age       int    `json:"-"`   This means that the age key will always be omited
}

func main() {
	person := Person{
		FirstName: "Chopperla",
		LastName:  "Santosh",
		Age:       20,
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("The json data is:", string(jsonData))
}
