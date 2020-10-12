//A var declaration can include initializers, one per variable.
//If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

package main

import "fmt"

var n1, n2 int = 10, 20

func main() {
	var c, python, javascript = false, true, "Hell yes"
	fmt.Println("Numbers", n1, n2)
	fmt.Println("Languages Known", "c: ", c, "python: ", python, "javascript", javascript)
}
