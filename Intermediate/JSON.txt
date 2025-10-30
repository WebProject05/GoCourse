package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string  `json:"firstName"`
	Age       int     `json:"age,omitempty"`
	Email     string  `json:"email,omitempty"`
	Address   Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	person := Person{
		FirstName: "Santosh",
		Email:     "San@gmail.com",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(jsonData))

	person1 := Person{
		FirstName: "Santosh",
		Age:       20,
		Email:     "Sant@mail.com",
		Address: Address{
			City:  "Hyd",
			State: "Telangana",
		},
	}

	jsonData1, err1 := json.Marshal(person1)
	if err1 != nil {
		fmt.Println("Error:", err1)
	}
	fmt.Println(string(jsonData1))

	jsonObj := `{
		"firstName": "rapa",
		"age": 20,
		"email": "rapa@gmail.com",
		"address": {
			"city": "Hyderabad",
			"state": "Telangana"
		}
	}`

	var personMate Person

	errObj := json.Unmarshal([]byte(jsonObj), &personMate)
	if errObj != nil {
		fmt.Println("Error:", errObj)
	}
	fmt.Println(personMate)
	fmt.Println(personMate.FirstName, "is from", personMate.Address.City)


	listOfCityState := []Address{
		{City: "Hyd", State: "Tg"},
		{City: "Vizag", State: "Ap"},
		{City: "Chennai", State: "Tn"},
		{City: "New Delhi", State: "Dl"},
	}

	// var cityList Address

	jsonCity, errCity := json.Marshal(listOfCityState)
	if errCity != nil {
		fmt.Println("Error:", errCity)
	}
	fmt.Println("The json List is:", string(jsonCity))

	// handling unknow JSON structures
	jsonData2 := `{"name": "john", "age": 20, "address": {"city": "Mim", "state": "fl"}}`

	var data map[string]interface{}

	err4 := json.Unmarshal([]byte(jsonData2), &data)
	if err4 != nil {
		log.Fatalln("Error:", err4)
	}

	fmt.Println(data)
	fmt.Println("Name of the person:",data["name"])



}
