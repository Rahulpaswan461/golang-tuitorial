package main

import "fmt"

func addTwoNumber(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Defer keyword")
	fmt.Println("The first statemenet")
	defer fmt.Println("The second one ")
	number := addTwoNumber(2, 3)
	fmt.Println("the return value is", number)
	// fmt.Println("sum is", sum)
	fmt.Println("The third one")
}
