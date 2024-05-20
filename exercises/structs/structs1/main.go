// structs1
// Make me compile!
//
package main

import "fmt"

type Person struct {
 Name string
	Age int
}

func main() {
	person := Person{Name:"Adam", Age:21}
	fmt.Printf("Person %s and age %d", person.Name, person.Age)
}
