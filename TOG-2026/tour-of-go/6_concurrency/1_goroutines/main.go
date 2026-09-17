package main

import (
	"fmt"
	"time"
)

func say(s string) {
	// for i := 0; i < 5; i++ {
	for range 5 {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func main() {
	go say("hello r")
	say("hello")
}
