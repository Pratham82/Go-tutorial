package main

import "fmt"

func main() {
	a := 4
	squaredVal := SquareVal(a)

	fmt.Println("Squared value of a before mutating:", squaredVal)
	fmt.Println("Value of a before mutating:", a)
	fmt.Println("Address of a:", &a)

	SquareAdd(&a)

	fmt.Println("Value of a after mutating:", a)
	fmt.Println("Address of a:", &a)
}

func SquareVal(n int) int {
	return n * n
}

func SquareAdd(n *int) int {
	*n = *n * *n

	updatedVal := *n

	return updatedVal
}
