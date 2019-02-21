package main

import "fmt"

func main() {
	/* Declare map: map[key]value */
	emails := make(map[string]string)

	/* Assign key values: */
	emails["Bob"] = "bob@gmail.com"
	emails["Peter"] = "peter@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails, len(emails))
	fmt.Println(emails["Bob"])

	/* Delete from map */
	// Delete takes collection and then index/key as arguments:
	delete(emails, "Bob")
	fmt.Println(emails) // No more bob :(

	/* Declaring and setting values at same time */
	newEmails := map[string]string{
		"Bob":   "bob@gmail.com",
		"Peter": "peter@gmail.com",
	}
	fmt.Println(newEmails)
}
