package main

import "fmt"

func main() {
	x := 5
	y := 10

	/* IF STATEMENTS */

	// Common practice to leave out parenthesis
	if x <= y {
		// %d prints out the value of any decimal type (numbers)
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	/* IF, ELSE IF, ELSE */
	color := "blue"
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not blue or red")
	}

	/* SWITCH */
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not blue or red")
	}

}
