package main

import "fmt"

var word string

var (
	word1 = "Hello"
	word2 = "world"
)

func main() {
	var num int
	num = 5

	var num2 = 6

	num3 := 7

	num3 = 8

	var x, y = 2, 3

	fmt.Println(x, y)
	fmt.Print(num)
	fmt.Println(num2)
	fmt.Printf("Number %v\n", num3)
	fmt.Printf("words %s, %v\n", word1, word2)
}
