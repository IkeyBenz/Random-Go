package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

// Since both params are ints, you just have
// to say it once:
func getDifference(num1, num2 int) int {
	return num1 - num2
}

func main() {
	fmt.Println(greeting("Ikey"))
	fmt.Println(getSum(3, 4))
	fmt.Println(getDifference(5, 2))
}
