package main

import (
	"fmt"
	"time"
)

type counter struct {
	total       int // zero
	lastUpdated time.Time
}

// pointer receiver
func (c *counter) increment() {
	c.total++ // incrementing it to 1
	c.lastUpdated = time.Now()
}

func (c counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v\n", c.total, c.lastUpdated)
}

func valueUpdate(c counter) {
	c.increment()
	fmt.Println("value updated ", c.String())
}

func pointerUpdate(c *counter) {
	c.increment()
	fmt.Println("pointer updated ", c.String())
}

func main() {

	var c counter
	fmt.Println(c.String())
	c.increment() // (&c).increment()
	fmt.Println(c.String())

	var p counter
	valueUpdate(p)
	fmt.Println("Main function", p.String())
	pointerUpdate(&p)
	fmt.Println("Main function", p.String())

}
