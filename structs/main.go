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
	// Declare struct
	ola := person{"Ola", "Norman", contactInfo{"ola@test.no", 12345}}
	fmt.Println(ola)
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
	var harald person
	fmt.Println(harald)
	// %+v - print out filed names and values
	fmt.Printf("%+v", harald)
	// Update struct
	harald.firstName = "Harald"
	harald.lastName = "Hårfagre"
	harald.contact = contactInfo{"fugl@mail.no", 1}
	harald.print()
	// This will not update the original object, only its copy
	harald.updateFirstName("HARALD")
	harald.print()
	// &variable name - give access to memory address of the variable (reference to memory address)
	haraldPointer := &harald
	haraldPointer.updateFirstN("HARALD")
	harald.print()
	// we can use the function with pointer with types to shortcut reference
	harald.updateFirstN("HaRaLd")
	harald.print()
	bil := "Bil"
	fmt.Println(*&bil)
}

func (p person) updateFirstName(updatedName string) {
	p.firstName = updatedName
}

func (pointerToPerson *person) updateFirstN(updateName string) {
	// *pointer - uses memory address to get the value from this address (direct access to value)
	(*pointerToPerson).firstName = updateName
}

func (p person) print() {
	fmt.Println()
	fmt.Println(p.firstName)
	fmt.Println(p.lastName)
	fmt.Println(p.contact.email)
	fmt.Println(p.contact.zipCode)
}
