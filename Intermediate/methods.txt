package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

// Method with value receiver

func (r *Rectangle) getArea() float64 {
	return r.length * r.width
}


// Method with the pointer receiver

func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func main() {
	rectange := Rectangle{
		length: 10,
		width:  10,
	}

	rectange.Scale(5)

	area := rectange.getArea()
	fmt.Println(area)

	num := MyInt(5)
	fmt.Println(num.isPositive())
	fmt.Println(num.welcomeMessage())
}


type MyInt int

func (m MyInt) isPositive() bool {
	return m > 0
}


func (MyInt) welcomeMessage() string {
	return "Hello myInt Structure"
}