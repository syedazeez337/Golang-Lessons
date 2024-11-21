package app

import "fmt"

// does not get exported
func add(x, y int) int {
	fmt.Println("Add function")
	return x + y
}

// exported function
func Mul(x, y int) int {
	fmt.Println("Mul function")
	return x * y
}
