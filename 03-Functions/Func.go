package main

import "fmt"

func greet() {
	fmt.Println("Hello, World!")
}

func greetUser(name string) { // Here we dont need to use var keyword to declare variables
	fmt.Println("Hello, ", name, "!")
}

func add(a int, b int) int { // func with return type
	return a + b
}
func mul_div(a int, b int) (int, float64) { // func with multiple return types
	product := a * b
	division := float64(a) / float64(b)
	return product, division
}
func main() {
	greet()
	greetUser("Zaid Alam")
	sum := add(10, 5)
	product, division := mul_div(5, 6)
	fmt.Println("sum:", sum)
	fmt.Println("product:", product)
	fmt.Println("division:", division)
}
