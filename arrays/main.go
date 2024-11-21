package main

import "fmt"

func main() {
	var arr [5]int

	// 0,1,2,3,4
	arr[0] = 23
	arr[2] = 43
	arr[3] = 54
	arr[4] = 88

	fmt.Println(arr)

	var s = [3]string{"a", "b", "c"}
	fmt.Println(s)

	var p = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 3, 4, 5}
	fmt.Println(p, len(p))

	for _, add := range p {
		fmt.Printf("elem = %d, address = %v\n", add, &add)
	}

}
