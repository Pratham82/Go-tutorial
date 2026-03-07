package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Name struct {
	firstName string
	lastName  string
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println(v)

	fullName := Name{"Prathamesh", "Malie"}

	fmt.Println(fullName)
	fmt.Println("FirstName", fullName.firstName)
	fmt.Println("LastName", fullName.lastName)

}
