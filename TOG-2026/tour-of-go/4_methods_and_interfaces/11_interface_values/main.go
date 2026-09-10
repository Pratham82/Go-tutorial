package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

type S string

func (s S) M() {
	fmt.Println(s)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	// Using new string variable
	// var n S
	// n = S("Prathamesh")

	n := S("Prathamesh")
	describe(n)
	n.M()
}

func describe(i I) {

	fmt.Printf("(%v, %T)\n", i, i)
}
