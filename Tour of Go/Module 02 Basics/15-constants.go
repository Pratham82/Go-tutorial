/*
Constants are declared like variables, but with the const keyword.
Constants can be character, string, boolean, or numeric values.
Constants cannot be declared using the := syntax.
*/

package main

import "fmt"

const Pi = 3.14
const cube_sides = 6

func main() {
	const World = "Alfred"
	fmt.Printf("Hello %v\n", World)

	const Truth = true
	fmt.Println("Go rules? ", Truth)
}
