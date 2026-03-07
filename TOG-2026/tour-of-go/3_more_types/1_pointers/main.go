package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i               // point to i
	fmt.Println("p:", *p) // read i through the pointer
	*p = 21               // set i through the pointer
	fmt.Println("i:", i)  // see the new value of i

	p = &j               // point to j
	*p = *p / 37         // divide j through the pointer
	fmt.Println("j:", j) // see the new value of j

	a := 100
	b := &a

	fmt.Println(a)  // 100
	fmt.Println(*b) // 100

	*b = 200

	fmt.Println(a)  // 200
	fmt.Println(*b) //200

}
