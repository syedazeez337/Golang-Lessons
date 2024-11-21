package main

import "fmt"

func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func main() {

	x := 10

	failedUpdate(&x)
	fmt.Println(x)

	update(&x)
	fmt.Println(x)
}

/*

	var x int = 45
	var b bool = true

	addX := &x
	addB := &b

	fmt.Printf("Address %v  and value %v of x\n", addX, *addX)
	fmt.Printf("Address %v and value %v of b\n", addB, *addB)

	var y *int
	y = &x
	fmt.Println(*y)

	var z = new(int)
	fmt.Println(z)
	fmt.Println(*z)
*/
