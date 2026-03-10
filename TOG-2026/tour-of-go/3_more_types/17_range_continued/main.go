package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	for index, value := range pow {
		fmt.Println("index:", index, "value:", value)
	}

	fmt.Println("pow", pow)
}
