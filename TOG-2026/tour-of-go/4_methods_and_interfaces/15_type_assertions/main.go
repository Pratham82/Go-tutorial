package main

import "fmt"

type Person struct {
	id      int
	name    string
	surname string
}

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)

	// new testing code
	var i2 interface{} = Person{
		id:      1,
		name:    "Prathamesh",
		surname: "mali",
	}

	info := i2.(Person)
	fmt.Println(info)
}
