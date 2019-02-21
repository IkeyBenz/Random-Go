package main

import "fmt"

// Person defines the properties that all person
// objects should have
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func main() {
	// Initialize person using struct
	person1 := Person{
		firstName: "Ikey",
		lastName:  "Benzaken",
		city:      "NY",
		gender:    "M",
		age:       20,
	}

	// Alternative:
	person2 := Person{"Ikey", "Benzaken", "NY", "M", 20}
	fmt.Println(person1, person2)
}
