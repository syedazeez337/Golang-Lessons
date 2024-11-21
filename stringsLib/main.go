package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello Go language"
	// b := "ld"

	substring := a[4:8]
	fmt.Println(substring)

	fmt.Println(strings.ToUpper(a))
	fmt.Println(strings.ToLower(a))

	// fmt.Println(strings.Split(a, ","))

	// fmt.Println(strings.Replace(a, "world", "gopher", 1))

	/*
		stringSlice := []string{
			"go",
			"is",
			"awesome",
		}

		fmt.Println(strings.Join(stringSlice, "-"))

		// file name main_test.go

		fmt.Println(strings.HasPrefix(a, "hello"))
		fmt.Println(strings.HasSuffix(a, "old"))

		fmt.Println(strings.Contains(a, b))
		fmt.Println(strings.Compare(a, b))
	*/

}
