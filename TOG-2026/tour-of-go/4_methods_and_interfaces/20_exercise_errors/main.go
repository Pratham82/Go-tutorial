package main

import (
	"fmt"
	"math"
)

type ErrorNegativeSqrt float64

func (e ErrorNegativeSqrt) Error() string {

	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))

}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrorNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
