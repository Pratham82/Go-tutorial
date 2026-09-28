package main

import (
	"fmt"
	"strconv"
)

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.

func DescribeNumberBox(nb NumberBox) string {
	unwrapped := nb.Number()
	return fmt.Sprintf("This is a box containing the number %.1f", float64(unwrapped))
}
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	unwrapped := fnb.Value()
	converted, err := strconv.Atoi(unwrapped)

	if err != nil {
		return 0
	}

	return converted
}

func main() {
	fmt.Println(ExtractFancyNumber(FancyNumber{"10"}))

}

// Valid dereferencing: c points to b (c = &b), so *c reads/writes b's value.
// c is valid to dereference because it holds a real address, unlike a above.
