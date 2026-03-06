package main

import (
	"fmt"
	"math"
)

func isAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(isAdult(20))
}
