package main

import "fmt"

func main() {
	/* ARRAYS  (set number of values) */
	var fruitsArr [2]string

	// Assign values
	fruitsArr[0] = "Apple"
	fruitsArr[1] = "Orange"

	// To declare and assign at same time:
	moreFruits := [2]string{"Apple", "Orange"}

	fmt.Println(fruitsArr, moreFruits)
	fmt.Println(fruitsArr[1], moreFruits[1])
	fmt.Printf("%T", fruitsArr)

	/* SLICES (unknown amount of values) */
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)
	fmt.Printf("%T", fruitSlice)

	// To get length of slice or array use len()
}
