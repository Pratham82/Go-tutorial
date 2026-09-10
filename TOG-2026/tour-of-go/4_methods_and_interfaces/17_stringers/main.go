package main

import "fmt"

type Person struct {
	id   int
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{1, "Arthur Dent", 42}
	z := Person{2, "Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
