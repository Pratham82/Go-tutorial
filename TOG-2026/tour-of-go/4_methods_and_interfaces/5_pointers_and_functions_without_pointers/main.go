/*
* Without pointers
*
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

// Scale takes v as a plain Vertex (no pointer).
// Every call to Scale COPIES the whole struct into a brand-new,
// independent box that only this function can see.
func Scale(v Vertex, f float64) {
	fmt.Println(v) // prints the VALUE, e.g. {3 4} - a separate copy, not main's box
	// These writes only ever touch Scale's own private copy of v.
	v.X = v.X * f
	v.Y = v.Y * f
	// When Scale returns, this copy (now {30, 40}) is discarded entirely -
	// nothing outside Scale ever sees it.
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v) // {3 4}

	// v is passed BY VALUE here - Scale gets its own copy of the struct.
	// main's v and Scale's v are two different boxes in memory from
	// this point on.
	Scale(v, 10)

	// main's v was never touched by Scale - it's still {3, 4}.
	fmt.Println(Abs(v)) // sqrt(3*3 + 4*4) = sqrt(25) = 5 (scaling was lost)
}
