package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (r rectangle) perim() float64 {
	return 2 * (r.length + r.width)
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perim() float64 {
	return 4 * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	/*
		r := rectangle{
			length: 4,
			width:  5,
		}

		fmt.Println(r)
		fmt.Println("Area", r.area())
	*/
	r := rectangle{
		length: 4,
		width:  5,
	}

	c := circle{
		radius: 5,
	}

	s := square{
		side: 6,
	}

	measure(r)
	measure(c)
	measure(s)

}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
