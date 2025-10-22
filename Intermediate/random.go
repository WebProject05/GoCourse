package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Using the seed
	// val := rand.New(rand.NewSource(404))
	val := rand.New(rand.NewSource(time.Now().Unix())) // New seed every time the program is run

	// fmt.Println(rand.Intn(101))   Here the number generated will be different
	fmt.Println(val.Intn(101)) // same number will be genetate because we have used the seed

}