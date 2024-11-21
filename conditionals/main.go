package main

import "fmt"

func main() {
	/*
		num := 34

		if num >= 20 {
			fmt.Printf("Number %d is greater than 20\n", num)
		} else if num >= 30 {
			fmt.Printf("Number %d is greater than 30\n", num)
		} else {
			fmt.Println("Number is less than 20")
		}
	*/

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("It is a Monday")
		fallthrough
	case "Tuesday":
		fmt.Println("It is a Tuesday")
	default:
		fmt.Println("It can be any other day")
	}

	/*
		number := 10

		switch {
		case number > 5:
			fmt.Println("Greater than 5")
		case number < 5:
			fmt.Println("Less than 5")
		case number == 5:
			fmt.Println("Equal to 5")
		default:
			fmt.Println("It can be anything")
		}
	*/
}
