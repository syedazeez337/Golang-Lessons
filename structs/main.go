package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{
		name: name,
		age:  age,
	}
	return &p

}

func main() {
	/*
		var s person

		s.name = "Good"
		s.age = 25
		fmt.Printf("%v %[1]T\n", s.name)
		fmt.Printf("%v %[1]T\n", s.age)
	*/
	m := newPerson("Hello", 45)
	fmt.Println(m.age)
	fmt.Println(m.name)
}
