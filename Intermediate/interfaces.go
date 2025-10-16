package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	height, width float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("The area of the shape is: ",g.area())
	fmt.Println("The perimeter of the shape is: ",g.perimeter())
}

func main() {
	r := rect{width: 5, height: 10}
	c := circle{radius: 10}

	measure(r)
	fmt.Println("----")
	measure(c)
}
