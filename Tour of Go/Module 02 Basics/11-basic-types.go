/*
Go basic typs

bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point
float32 float64
complex64 complex128

Variable declarations may be "factored" into blocks, as with import statements.

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
*/

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	Maxint uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	name   string     = "Duke"
)

func main() {
	fmt.Printf("Type %T value: %v\n", ToBe, ToBe)
	fmt.Printf("Type %T value: %v\n", Maxint, Maxint)
	fmt.Printf("Type %T value  %v\n", z, z)
	fmt.Printf("Type %T value %v\n", name, name)
}
