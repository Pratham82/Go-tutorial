// Type conversions
// The expression T(v) converts the value "v" to the Type "T";w
/*
var i = 43
var f float64 =  float64(i)
var u uint = uint(f)

Or using short declarations
i := 42
f := float64(i)
u := uint(f)
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 23, 45
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Printf("x:  %v y: %v f: %v z: %v\n", x, y, f, z)
}
