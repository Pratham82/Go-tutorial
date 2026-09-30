package main

import "fmt"

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}
func (r *Resident) HasRequiredInfo() bool {
	return r.Name != "" && r.Address["street"] != ""
}

func main() {
	name1 := "Matthew Sanabria"
	age1 := 29
	address1 := map[string]string{"street": "Main St."}

	resident1 := NewResident(name1, age1, address1)

	name2 := "Rob Pike"
	age2 := 0
	address2 := make(map[string]string)

	resident2 := NewResident(name2, age2, address2)

	residents := []*Resident{resident1, resident2}

	for _, v := range residents {
		fmt.Println(v)
		fmt.Println(v.HasRequiredInfo())
	}
}

// Valid dereferencing: c points to b (c = &b), so *c reads/writes b's value.
// c is valid to dereference because it holds a real address, unlike a above.
