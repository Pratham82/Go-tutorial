package main

import (
	"fmt"
	"strings"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func getName(fullName string) (first, last string) {
	parts := strings.Split(fullName, " ")
	first = parts[0]
	last = parts[1]
	fmt.Println(first, "\n", last)
	return
}

func main() {
	fmt.Println(split(17))
	fmt.Println(getName("Prathamesh Mali"))
}
