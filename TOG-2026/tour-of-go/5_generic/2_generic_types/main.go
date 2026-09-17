package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type

type List[T any] struct {
	next *List[T]
	val  T
}

// Push prepends a new node with val to the front of the list, returning the new head.
func (l *List[T]) Push(val T) *List[T] {
	return &List[T]{next: l, val: val}
}

// Len returns the number of nodes in the list.
func (l *List[T]) Len() int {
	count := 0
	for n := l; n != nil; n = n.next {
		count++
	}
	return count
}

// ToSlice collects all values into a slice, in list order (head to tail).
func (l *List[T]) ToSlice() []T {
	var out []T
	for n := l; n != nil; n = n.next {
		out = append(out, n.val)
	}
	return out
}

// Find returns the first node whose val equals target, or nil if not found
func (l *List[T]) Find(target T, eq func(a, b T) bool) *List[T] {
	for n := l; n != nil; n = n.next {
		if eq(n.val, target) {
			return n
		}
	}
	return nil
}

func main() {
	var list *List[int]

	list = list.Push(3)
	list = list.Push(3)
	list = list.Push(1)

	fmt.Println(list.ToSlice())
	fmt.Println(list.Len())

	found := list.Find(2, func(a, b int) bool { return a == b })

	if found != nil {
		fmt.Println("found:", found.val)
	}

}
