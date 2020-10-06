package main

import "fmt"

// If two of more function parameters shares the type, the first type decrlaration can be skipped but the last type declaration needs to be there

func attach_strings(fname, lname string) string {
	return "Hi My name is " + fname + " " + lname
}

func main() {
	fmt.Println(attach_strings("Prathamehs", "Mali"))
}
