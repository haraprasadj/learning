// Go supports methods defined on struct types
package main

import "fmt"

type rect struct {
	width, height int
}

// Methods can be defined for pointer or value receiver types

// This method has a receiver type of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// This method has a receiver type of rect
func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values and pointers for method calls
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
