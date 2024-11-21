package main

import "fmt"

func main() {
	/*
		// slice
		var x []int
		fmt.Println(x, len(x))

		var p = []int{1,2,3,4,5}
		// fmt.Println(p[3])

		p = append(p, 6, 7, 8, 9, 10)
		fmt.Println(p, len(p))

		for i, elem := range p {
			fmt.Printf("i = %d, elem = %d, address = %v\n", i, elem, &elem)
		}


		var x = make([]int, 5)
		fmt.Println(x, len(x), cap(x))

		x = append(x, 1, 2, 3)
		fmt.Println(x, len(x), cap(x))

		y := make([]int, 4, 6)
		fmt.Println(y, len(y), cap(y))

		// make(), len(), cap(), append(), clear()
		var p = []int{1,2,3,4,5}
		fmt.Println(p, len(p), cap(p))
		clear(p)
		fmt.Println(p, len(p), cap(p))


		// slicing [1:3]
		slic := []int {1,2,3,4,5}
		//             0-1-2-3-4
		f := slic[:3]
		fmt.Println(f)
		for _, elem := range f {
			fmt.Printf("%v %v\n", elem, &elem)
		}
		y := slic[2:4]
		fmt.Printf("%v %[1]T\n", y)
		for _, elem := range y {
			fmt.Printf("%v %v\n", elem, &elem)
		}

		new_slic := make([]int, 5)
		num := copy(new_slic, slic)
		fmt.Println(new_slic, num)
	*/

	s := "Hello world"
	for _, n := range s {
		fmt.Println(n)
	}
}
