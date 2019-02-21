package main

import "fmt"

func main() {
	/* FOR LOOP */

	// Long method:
	// Here, 'for' acts like 'while'
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Short (standard) method:
	for i := 1; i < 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		// If i is divisible by 3 and 5:
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
			// If i is only divisible by 3:
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
