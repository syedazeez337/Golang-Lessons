package main

import "fmt"

func main() {
	fmt.Println("hello world")

	x, y, z := multiple(4, 6)
	fmt.Println(x, y, z)
}

func multiple(x, y int) (int, int, int) {
	add := x + y
	sub := x - y
	mul := x * y

	return add, sub, mul
}
