package main

import "fmt"

func main() {

	/* RANGE WITH SLICES */

	ids := []int{33, 76, 54, 23, 11, 2}

	// Loop through ids (with indices):
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Loop through ids (without indices):
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	/* RANGE WITH MAPS */
	emails := map[string]string{
		"Bob":   "bob@gmail.com",
		"Peter": "peter@gmail.com",
	}
	// Loop through key, value of map
	for name, email := range emails {
		fmt.Printf("%s: %s\n", name, email)
	}

	// If you omit the value, it will only give you keys
	for name := range emails {
		fmt.Println("Name: " + name)
	}

	// If you only want the values:
	for _, email := range emails {
		fmt.Println("Email: " + email)
	}
}
