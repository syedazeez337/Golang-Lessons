package main

import "fmt"

func main() {
	// classic for loop
	/*
		for i:=0; i < 10; i++ {
			fmt.Printf("Number %d\n", i)
		}*/

	// while loop
	/*
		i := 10
		for i < 20 {
			fmt.Printf("Number %d\n", i)
			i++
		}*/

	for {
		fmt.Println("infinite loop")
	}

	/*
		for i := range 10 { // n - 1 <
			fmt.Printf("Number %d\n", i)
		}*/

	/*
		for i := range 10 {
			if i == 3 {
				break
			}
			fmt.Println(i)
		}*/

	/*
		for i := range 10 {
			if i == 3 {
				continue
			}
			fmt.Println(i)
		}*/
}
