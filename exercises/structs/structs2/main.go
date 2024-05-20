// structs2
// Make me compile!
//
package main

import "fmt"

type ContactDetails struct {
	Phone int
}

type Person struct {
	// don't just create the phone field here. embed a new struct
	name string
	age int
	contact ContactDetails
}

func main() {
	contactDetails := ContactDetails{Phone:1234}
	person := Person{name: "John", age: 32, contact:contactDetails}
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.contact.Phone)
}
