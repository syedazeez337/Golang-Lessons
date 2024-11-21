package main

import "fmt"

type person struct {
	name string // ""
	age  int    // 0
}

func (p *person) greet() string {
	if p == nil {
		return fmt.Sprintln("Nil instance")
	}
	return fmt.Sprintf("Name %v, age %v\n", p.name, p.age)
}

func main() {
	var p *person
	fmt.Println(p)
	fmt.Println(p.greet())

	p = &person{
		name: "Good",
		age:  34,
	}

	fmt.Println(p.greet())
}

/*
type adder struct {
	start int
}

func (a adder) addTo(val int) int {
	return a.start + val
}

func funcadd(x adder, val int) int {
	return x.start + val
}

func main() {
	x := adder{start: 10}
	fmt.Println(x.addTo(5))
	y := funcadd(x, 5)
	fmt.Println(y)
}
*/
