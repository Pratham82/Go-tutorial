package main

import "fmt"

func main() {
	var sum = 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	for i := range 5 {
		fmt.Println(i)
	}

	fmt.Println(sum)
}
