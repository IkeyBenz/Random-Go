package main

import "fmt"

func main() {
	// a equals 5
	a := 5
	// b equals the memory address of the a variable
	b := &a

	fmt.Println(a, b)
	fmt.Printf("a is a %T\nb is a %T\n", a, b)
	// a is a int
	// b is a *int, the star represents a pointer

	// To read value from memory address, use the star
	fmt.Println(*b) // same as *&a

	// To change value using pointer, also use the star
	*b = 10
	fmt.Println(a)

	// You'd use a pointer to pass a shitload of data to different
	// files. Rather than sending ALL the data... you just send the
	// address of the data.
}
