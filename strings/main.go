package main

import (
	"fmt"
)

func main() {
	/*
		word := "Hello world"
		//       01234567890
		for _, i := range word {
			fmt.Println(string(i))
		}
		fmt.Println("length", len(word))

		s := word[3:7]
		fmt.Println(s)
	*/

	// runes
	// char type in c
	var a rune = 'x'
	fmt.Println(string(a))
	b := 'Y'
	fmt.Printf("%v %[1]T\n", b)

	word := "Hello world"
	// strings are immutable
	runes := []rune(word)
	bytes := []byte(word)
	fmt.Println(runes)
	fmt.Println(bytes)

	fmt.Println(string(word[3]))

}
