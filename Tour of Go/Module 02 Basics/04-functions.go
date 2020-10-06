package main

import "fmt"

// Simple function in go
// The data type comes after the variable name
// We have to explicitly add the return type for the function as well

func myfunc(n1 int, n2 int) int {
	return n1 + n2
}

// Driver function
func main() {
	fmt.Println(myfunc(323, 435))
}
