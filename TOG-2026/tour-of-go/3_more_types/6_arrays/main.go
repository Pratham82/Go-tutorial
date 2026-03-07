package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	// slice (dynamic array)
	// This creates a slice that is nil.
	// nums = nil

	var nums []int

	for i := range 6 {
		nums = append(nums, i)
	}

	// slices (not null values)
	// This creates an empty slice but NOT nil.
	// nums2 = []
	nums2 := []int{}

	for i := range 6 {
		nums2 = append(nums2, i)
	}

	// array (fixed size)
	var arr [3]int

	for i := range 3 {
		arr[i] = i
	}

	fmt.Println("primes", primes)
	fmt.Println("nums", nums)
	fmt.Println("nums2", nums2)
	fmt.Println("arr", arr)
}
