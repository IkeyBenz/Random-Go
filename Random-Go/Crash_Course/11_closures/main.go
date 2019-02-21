package main

import "fmt"

// This is a function that returns a function
// .. that returns an integer
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := adder()
	/*
		sum equals adder, which equals the anonymous
		function inside of adder. Calling 'sum(i)' only
		calls the inside function, so even though local
		'sum' is 0 when the adder function is initialized,
		it never gets reset. So calling sum multiple times
		will result in the sum accumulating.
	*/
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
