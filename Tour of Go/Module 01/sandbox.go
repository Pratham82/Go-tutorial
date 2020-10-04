package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Go sandbox")
	fmt.Println("Current time is:", time.Now().Format(time.Stamp))
}
