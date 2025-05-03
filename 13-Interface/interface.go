package main

import "fmt"

// Define an interface
type Shape interface {
	Area()
}

// Creating a struct that implements the interface
type Circle struct{}

// Extending the interface to include parameters
func (c Circle) Area(r float64) {
	area := 3.14 * r * r
	fmt.Printf("Area of Circle: %.3f\n", area)
}

type Rectangle struct{}

func (r Rectangle) Area(l, b float64) {
	area := l * b
	fmt.Printf("Area of Rectangle %.3f", area)
}

func main() {
	var c Circle
	var r Rectangle
	c.Area(5.0)
	r.Area(5.0, 10.0)
}
