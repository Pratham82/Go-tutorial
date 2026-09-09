package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so

func (t T) M() {
	fmt.Println(t.S)
}

// new func
type I2 interface {
	greet(name string) string // just a contract — no code
}

type Foo struct{} // a concrete type — the "who" that will implement the contract

func (f Foo) greet(name string) string { // the actual implementation, attached to Foo
	fmt.Printf("Hello %s !!", name)
	return name
}
func main() {
	var i I = T{"hello"}
	var myName = "Prathamesh"

	var foo I2 = Foo{}

	i.M()

	// foo.greet(myName) calls the greet method (defined above on Foo) through
	// the I2 interface. It has to be called on foo (a value of a type that
	// satisfies I2) because greet is a method, not a standalone function —
	// there's no bare greet(myName) to call; the receiver is required.
	foo.greet(myName)
}
