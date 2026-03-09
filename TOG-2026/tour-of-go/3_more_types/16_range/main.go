package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for k, j := range nums {
		fmt.Println("j", j, "k:", k)

	}
}
