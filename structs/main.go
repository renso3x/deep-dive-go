package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo

	// shorthand alternative
	contactInfo
}

func main() {
	// declare a struct
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	// alternative
	var bob person
	bob.firstName = "Bob"
	bob.lastName = "Foo"
	fmt.Println(bob)

	jim := person{
		firstName: "Jimy",
		lastName:  "Butter",
		// alternative
		// contact: contactInfo{
		// 	email:   "jimmy@buckets.com",
		// 	zipCode: 1231,
		// },
		// alternative
		contactInfo: contactInfo{
			email:   "jimmy@buckets.com",
			zipCode: 1231,
		},
	}

	// &variable_name - give me the memory variable pointing at
	// jumPointer := &jim
	// jumPointer.updateName("Jimmy")

	// shortcut for a pointer
	jim.updateName("Jimmy")
	jim.print()
}

// *person - give me the value of this memory address
func (pointerToPerson *person) updateName(fname string) {
	// pointer of this person then change the property
	(*pointerToPerson).firstName = fname
}

// structs can have a receiver
func (p person) print() {
	// %+v prints out all the property name of struct
	fmt.Printf("%+v", p)
}
