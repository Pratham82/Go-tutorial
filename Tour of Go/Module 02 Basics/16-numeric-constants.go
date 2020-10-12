//Numeric constants
/*
Numeric constants are high-precision values.
An untyped constant takes the type needed by its context.
*/

package main

import "fmt"

const (
	// Creating a huge number by a 1 bit left 100 places
	Big = 1 << 100

	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println(needInt(Small))
	// fmt.Println(needInt(Big)) =>  constatnt 1267650600228229401496703205376 overflows int
	// Reason:  An int can store at maximum a 64-bit integer, and sometimes less.)
	fmt.Println(needFloat(Big))
}
