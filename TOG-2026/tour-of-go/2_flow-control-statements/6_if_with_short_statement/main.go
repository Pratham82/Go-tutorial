package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func getIfShortChar(str string) bool {
	if c := len(str); c < 10 {
		return true
	}
	return false
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(getIfShortChar("Prathamesh"))
	fmt.Println(getIfShortChar("Pratham"))
}
