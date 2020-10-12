/*
Zero values

Variables declared without any explicit values are given thier zero values
0 : for numeric values
false: for boolean type
"" : for strings
*/

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
