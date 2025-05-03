package main

import "fmt"

type car struct {
	brand string
	speed int
}

// A Function that works with structs is called a method
// A method is a function with a special receiver argument

func (c car) drive() {
	fmt.Printf("The %s is driving at %d km/h \n", c.brand, c.speed)
}
func (c *car) accalarate() {
	c.speed += 10
}

func main() {
	mycar := car{brand: "porsche", speed: 100}
	mycar.drive()
	mycar.accalarate()
	mycar.drive()
}
