//go:build ignore

package main

import "fmt"

func main() {

	ch := make(chan int) // unbuffered: capacity 0

	go func() {
		ch <- 42 // blocks until main receives
	}()

	fmt.Println(<-ch) // blocks until goroutine sends
}
