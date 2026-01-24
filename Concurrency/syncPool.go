package main

import (
	"fmt"
	"sync"
)

/*
sync.Pool is a data structure provided by the sync package that facilitates the reuse of objects to help reduce the overhead of memory allocation, particularly in performance-sensitive applications.
*/

// Each pool can only have one Consumer at a time
// In data base connections creating new instances is expensive so we using pools.

type person struct {
	name string
	age int
}



func main() {
	var pool = sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new person")
			return &person{}
		},
	}


	person1 := pool.Get().(*person)
	person1.name = "Santosh"
	person1.age = 20
	fmt.Println("Got Person:", person1)

	fmt.Printf("Person1 - Name: %s, Age: %d \n", person1.name, person1.age)

	pool.Put(person1)
	fmt.Println("Returned Person 1 to the Pool")

	person2 := pool.Get().(*person)
	person2.name = "Rajanna"
	person2.age = 46

	fmt.Printf("Person2 - Name: %s, Age: %d \n", person2.name, person2.age)

}