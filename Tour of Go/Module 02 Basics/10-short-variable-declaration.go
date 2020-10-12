//Inside a function, the := short assignment statement can be used in place of a var declaration  with implicit type.

//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

package main

import "fmt"

// Short vairable declaration
// With this we don't have to add types explicitly
// But we can only use inside the function
// Outside the function we have to use the keywords

func main() {
	username := "Joey"
	john, tyrion, cerci := "Stark", "Lannister", "Lannister"
	length, breadth := 32, 45
	fmt.Println(username)
	fmt.Println("john: ", john, "tyrion: ", tyrion, "cerci: ", cerci)
	fmt.Println("Area: ", length*breadth)
}
