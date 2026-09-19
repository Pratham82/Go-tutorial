package main

import "fmt"

func main() {

	ch := make(chan int, 3)
	ch <- 1
	ch <- 11
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
