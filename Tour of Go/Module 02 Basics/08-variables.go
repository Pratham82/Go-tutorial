package main

import "fmt"

// The var statement declares a list of variables, for eg. function argument lists the type is the last
// var statement can be at package level or function level

var c, python, java, javascript bool

func main() {
	var i int // initial value of int will be 0

	// initial value of all the boolean
	fmt.Println(i, c, python, java, javascript)
}
