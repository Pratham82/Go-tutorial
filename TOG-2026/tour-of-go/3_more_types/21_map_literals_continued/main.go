package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

type Board struct {
	x, y int
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
	m := map[string]Board{
		"a": {
			x: 223232, y: 2323,
		},
		"b": {
			x: 23, y: 891,
		},
	}

	fmt.Println(m)
}
