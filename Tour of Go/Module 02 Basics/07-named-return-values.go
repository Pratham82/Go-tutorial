package main

import "fmt"

// Go's return values can be named, and they can be treated as variable defined at the top of the functions
// A return statement without arguments returns the named return value(known as naked return)
// Naked returns should only be used with short functions
func split(sum int) (x, y int) {
	x = sum * 10 / 2
	y = x + 100
	return
}

func main() {
	fmt.Println(split(14))
}
