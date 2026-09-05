/*
Here we see the Abs and Scale methods rewritten as functions.

Again, try removing the * from line 16. Can you see why the behavior changes? What else did you need to change for the example to compile?

(If you're not sure, continue to the next page.)
* */

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Abs takes v BY VALUE - it receives a copy of the struct.
// Fine here since Abs only reads X and Y, never writes them.
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale takes v as a POINTER (*Vertex), not a value.
// v does not hold a Vertex here - it holds the ADDRESS of one.
// No copy of the struct is made; there is only ONE Vertex in memory,
// and Scale is holding directions to it.
func Scale(v *Vertex, f float64) {
	fmt.Println(v) // prints an address, e.g. &{3 4} - same address as main's &v
	// v.X is shorthand for (*v).X: Go auto-dereferences the pointer,
	// follows it to the original struct, and writes the new value THERE.
	v.X = v.X * f
	v.Y = v.Y * f
	// After these two lines, main's v has actually been mutated - there
	// was never a second copy to lose the change to.
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(&v) // &v = address of this box, e.g. 0xc0000140a0

	// &v passes that address into Scale - not a copy of the struct.
	// Scale and main are now looking at the exact same memory.
	Scale(&v, 10)

	// v has been mutated in place through the pointer: now {30, 40}.
	fmt.Println(Abs(v)) // sqrt(30*30 + 40*40) = sqrt(2500) = 50
}
