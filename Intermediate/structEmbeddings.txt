package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	employeeDetials person
	employeeID int
	salary     float64
}

func main() {
	emp := employee{
		employeeDetials:     person{name: "Santosh", age: 20},
		employeeID: 20240020279,
		salary:     1200000,
	}

	fmt.Println(emp.employeeDetials.name)
}