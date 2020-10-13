// for is go's while

package main

import "fmt"

func main() {
	n := 0
	for n > 10 {
		n += n
	}
	fmt.Println(n)
}
