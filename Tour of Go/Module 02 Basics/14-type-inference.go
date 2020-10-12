//Type inference
/*
When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

var i int
j := i // j is an int
But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant

*/
package main

import "fmt"

func main() {
	val := "qword"
	fmt.Printf("Type of val: %T\n", val)
}
