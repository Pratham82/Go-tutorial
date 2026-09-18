// Package main demonstrates the window/offset-limit pattern:
// mapping a 1-indexed unit number (e.g. "week N") to a sub-slice.
// See README.md for the full explanation and diagram.
package main

import "fmt"

// weekSlice returns the size-item window belonging to the given 1-indexed
// week number, clamped to the bounds of xs.
func weekSlice(xs []int, week, size int) []int {
	start := (week - 1) * size
	if start > len(xs) {
		start = len(xs)
	}
	end := min(start+size, len(xs))
	return xs[start:end]
}

// chunks splits xs into consecutive windows of the given size.
func chunks(xs []int, size int) [][]int {
	var result [][]int
	for start := 0; start < len(xs); start += size {
		end := min(start+size, len(xs))
		result = append(result, xs[start:end])
	}
	return result
}

func main() {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for week := 1; week <= 3; week++ {
		fmt.Printf("week %d -> %v\n", week, weekSlice(xs, week, 7))
	}

	fmt.Println("all chunks:", chunks(xs, 7))
}
