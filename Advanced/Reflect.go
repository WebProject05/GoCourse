package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string
	Age int
}


func main() {
	p1 := person{Name: "Jonas", Age: 19}

	v := reflect.ValueOf(p1)

	for i := range v.NumField() {
		fmt.Printf("Field %d: %v \n", i, v.Field(i))
	}

	v1 := reflect.ValueOf(&p1).Elem()

	nameField := v1.FieldByName("Name")
	if nameField.CanSet() {
		nameField.SetString("Martha")
	} else {
		fmt.Println("Cannot set name or age")
	}

	fmt.Println("Modified person", p1)
}