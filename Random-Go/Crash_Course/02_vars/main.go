package main

import "fmt"

/*
 MAIN TYPES:
	string
	bool
	int
	int int8 int16 int32 int64 (larger postfix, larger the int)
	uint uint8 uint16 uint32 uint64 uintptr
	byte - alias for uint8
	rune - alias for uint32
	float32 float64 - decimal point numbers
	complex64 complex128 - really large numbers
*/

func main() {
	// Using the var keyword:
	/* Don't include a type when you initialize with a value
	   the type is inferred.
	   var name string = "Ikey" */
	var name = "Ikey"
	var age = 20
	var isCool = true

	fmt.Println(name, age, isCool)

	/*
		What if we wanted to print the type of a var?
		This link contains all the symbols you can use
		in your format string:
		https://golang.org/pkg/fmt/
	*/

	fmt.Printf("%T\n", name) // Get Type of name (string)

	// Reseting variable values:
	age = 21
	fmt.Println(age)

	// To declare constants:
	const isAGoodCoder = true
	/*
		See, even Go knows I'm a good coder...
		Can't reset this value!
		isAGoodCoder = false // Yields error
	*/

	// Shorthand variable creation:
	newName := "Ikey"
	newAge := 20

	// Even Shorter:
	newestName, newestAge := "Ikey", 20
	fmt.Println(newName, newAge, "\n", newestName, newestAge)
}
