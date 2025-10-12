package main

import "fmt"

func main() {
	// ... Ellipsis, it will accept variadic parameters
	// func funcName(para 1, para 2, para 3 ...)	 returnType {}

	summ := []int{1,2,3,4,5,6}

	sum := sum(summ...) 
	fmt.Println(sum)

}
// In this case we can also pass static parameters with the variadic parameter but this parameter shoud come at last
func sum(nums ...int) int { // nums will be a list
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}