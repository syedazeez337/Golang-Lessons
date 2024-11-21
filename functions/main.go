package main

import "fmt"

func main() {
	/*
		greet()
		customGreet("Golang")

		num := add(3, 4)
		fmt.Printf("Add 3 and 4 = %d\n", num)

		word := addString(97, "Three")
		fmt.Printf("Adding number and string %s\n", word)

		x, y, z := multipleReturns(3, 4)
		fmt.Printf("Add %d, Sub %d, Mul %d\n", x, y, z)
	*/

	/*
		res := sum(1,2,3,4,5,6,7,8,9,10)
		fmt.Printf("Sum is %d\n", res)


		func(name string) {
			fmt.Println("Hello", name)
		}("World")

		f := func(x int, y int) int {return x + y }(3, 4)
		fmt.Println(f)

		s := func() string {return fmt.Sprintln("Hello world")}()
		fmt.Println(s)



		a := 20
		f := func() {
			fmt.Println(a)
			a := 30
			fmt.Println(a)
		}
		f()
		fmt.Println(a)

		nextInt := seq()

		fmt.Println(nextInt())
		fmt.Println(nextInt())
		fmt.Println(nextInt())

		newInt := seq()
		fmt.Println(newInt())

		i := 4
		fmt.Println(i)
		normal()
		fmt.Println(i)
	*/

	// 5! = 5 x 4 x 3 x 2 x 1

	/*
		res := 1
		for i:=1; i <=5; i++ {
			res *= i
			// fmt.Println(i)
		}
		fmt.Println(res)

		fmt.Println(fact(5), fact(7))

		var fib func(n int) int

		// fib 5 = fib(4) + fib(3)

		fib = func(n int) int {
			if n < 2 {
				return n
			}
			return fib(n-1) + fib(n-2)
		}

		fmt.Println(fib(7), fib(8))

	*/
	s := []int{1, 2, 3, 4, 5}

	result := mapping(addTwo, s)
	fmt.Println(result)

	result = mapping(addThree, s)
	fmt.Println(result)

	defer fmt.Println("Hello")
	fmt.Println("World")
}

func addTwo(n int) int {
	return n + 2
}
func addThree(n int) int {
	return n + 3
}

// []int{1,2,3,4,5}
// add()

func mapping(function func(int) int, slic []int) []int {
	res := make([]int, len(slic))
	for i, v := range slic {
		res[i] = function(v)
	}
	return res
}

// recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func normal() int {
	i := 2
	return i
}

func seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func greet() {
	fmt.Println("Hello world")
}

func customGreet(s string) {
	fmt.Printf("Hello %s\n", s)
}

func add(x, y int) int {
	return x + y
}

func addString(x int, word string) string {
	res := string(x) + word
	return res
}

func multipleReturns(x, y int) (int, int, int) {
	add := x + y
	sub := x - y
	mul := x * y

	return add, sub, mul
}

// variadic functions
func sum(nums ...int) int {
	fmt.Println(nums)

	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}
