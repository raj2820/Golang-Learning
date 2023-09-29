package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// raj := person{"Raj", "Shinde"}

	// otherway

	// raj := person{firstName: "Raj", lastName: "Shinde"}

	// otherway

	// var raj person

	// raj.firstName = "Raj"
	// raj.lastName = "Shinde"

	// fmt.Printf("%+v", raj)

	// fmt.Println(raj)

	// with embedding

	raj := person{
		firstName: "Raj",
		lastName:  "Shinde",
		contact: contactInfo{
			email:   "raj@gmail.com",
			zipCode: 421301,
		},
	}

	rajPointer := &raj
	// fmt.Printf("%+v", raj)

	// fmt.Println(raj)
	rajPointer.updateName("raju")
	raj.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
