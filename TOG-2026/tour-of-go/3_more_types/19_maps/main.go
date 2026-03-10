package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var fullName map[string]Vertex

type Info struct {
	firstName string
	lastName  string
}

type User struct {
	id        int
	firstName string
	lastName  string
}

type EmployeeInfo struct {
	fName string
	lName string
	phone int
	email string
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	// f := fullName["Prathamesh"]{}

	k := Vertex{
		232323,
		-232232,
	}

	fmt.Println(m["Bell Labs"])
	fmt.Println(k)

	person := Info{
		firstName: "Prathamesh",
		lastName:  "Mali",
	}

	fmt.Println(person)

	// using map
	users := map[string]Info{
		"u1": {firstName: "John", lastName: "Doe"},
	}

	users["u2"] = Info{
		firstName: "John",
		lastName:  "Snow",
	}

	fmt.Println(users)

	// Creating map then adding structs

	employyee := make(map[string]EmployeeInfo)

	employyee["emp1"] = EmployeeInfo{
		fName: "Judy",
		lName: "Smith",
		phone: 98989988,
		email: "judy.smit@email.com",
	}

	fmt.Println(employyee)

	dbUser := map[int]User{
		1: {
			firstName: "User",
			lastName:  "1",
		},
	}

	fmt.Println(dbUser)

}
