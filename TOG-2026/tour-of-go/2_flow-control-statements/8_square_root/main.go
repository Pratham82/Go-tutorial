package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	// for i := range 10 {
	// 	z -= (z*z - x) / (2 * z)
	// 	fmt.Println(z)
	// }
	// return x
	for {
		prev := z
		z -= (z*z - x) / (2 * z)

		if math.Abs(z-prev) < 0.000001 {
			break
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(27))
}
