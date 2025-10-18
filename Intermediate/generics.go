package main

import "fmt"

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

type Stack[T comparable] struct {
	elements []T	// List(Slice) of elements of type any
}

func (s *Stack[T]) push(element T) {	// We are modifying the elements list by appending so use the pointer
	s.elements = append(s.elements, element)
}


func (s *Stack[T]) delete(element T) (T, bool) {
	for i, val := range(s.elements) {
		if val == element {
			s.elements = append(s.elements[:i], s.elements[i+1:]...)
			return val, true
		}
	}
	var zero T
	return zero, false
}


func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements) - 1]
	s.elements = s.elements[:len(s.elements) - 1]
	return element, true
}


func (s *Stack[T]) isEmpty() bool {
	if len(s.elements) == 0 {
		return true
	}
	return false
}


func (s Stack[T]) printStack() {
	if s.isEmpty() {
		fmt.Println("The stack is empty!")
	}
	for i, val := range s.elements {
		fmt.Println(i, val)
	}
}

func main() {
	// x, y := 1, 2
	// x, y = swap(x, y)
	// fmt.Println(x, y)

	// a, b := "San", "tosh"
	// a, b = swap(a, b)
	// fmt.Println(a, b)
	stack := Stack[int]{}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)
	stack.printStack()
	fmt.Println(stack.pop())
	fmt.Println(stack.delete(2))
	stack.printStack()
}