package main

import "fmt"

func main() {
	var sum = 0

	for i := range 10 {
		sum += i
	}

	for i := range 5 {
		fmt.Println(i)
	}

	fmt.Println(sum)
}
