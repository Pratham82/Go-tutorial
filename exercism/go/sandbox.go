package main

import "fmt"

func main() {
	// candidateName := "Prathamesh"
	// votes := 100

	// fmt.Println(fmt.Sprintf("%s (%d)", candidateName, votes))

	candidate := "Mary"
	var results = map[string]int{
		"Mary": 10,
		"John": 51,
	}

	for k, v := range results {
		if k == candidate {
			results[k] = results[k] - 1
		}
		fmt.Println(k, v)
	}
	fmt.Println(results)

}

// Valid dereferencing: c points to b (c = &b), so *c reads/writes b's value.
// c is valid to dereference because it holds a real address, unlike a above.
