package main

import "fmt"

/*
* Implement a fibonacci function that returns a function
* (a closure) that returns successive fibonacci numbers
* (0, 1, 1, 2, 3, 5, ...).
 */

// Fibonacci function
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

