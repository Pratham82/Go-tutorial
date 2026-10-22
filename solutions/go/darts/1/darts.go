package darts

import (
	"math"
)

func Score(x, y float64) int {
	res := math.Sqrt(x*x + y*y)

	switch {
	case res <= 1:
		return 10
	case res <= 5:
		return 5
	case res <= 10:
		return 1
	default:
		return 0
	}
}
