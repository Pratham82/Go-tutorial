package main

/*Every Go program is made up of packages.
Programs start running in package mainA
This program is using the packages with import paths "fmt" and "math/rand".
*/

import (
	"fmt"
	"math/rand"
	"time"
)

// You’ll always see the same sequence every time you run the program.
//This is because by default the seed is always the same, the number 1.
//To actually get a random number, you need to provide a unique seed for your program.
//Use rand.Seed() before calling any math/rand method, passing an int64 value.

func main() {
	//You just need to seed once in your program, not every time you need a random number.
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favourite number is", rand.Intn(10))
}
