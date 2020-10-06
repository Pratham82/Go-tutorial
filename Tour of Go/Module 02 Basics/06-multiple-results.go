package main

import (
	"fmt"
)

// A function can return any number of arguments in go, we just have to
func swap_values(s1, s2 string) (string, string) {
	return s2, s1
}

func main() {
	// This is short declaration in this we dont't have to explicitly provide the data types or the var
	a, b := swap_values("hello", "world")
	fmt.Println(swap_values("go", "corona"))
	fmt.Println(a, b)
}
