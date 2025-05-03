package main

import "fmt"

func main() {
	// Array declaration and initialization

	var languages [5]string = [5]string{"Go", "Python", "Java", "C++", "JavaScript"}

	// Array declaration and initialization (shorter syntax)
	country := [5]string{"Japan", "Turkey", "Swizerland", "Australia", "Philippines"}

	for i := 0; i < 5; i++ {
		fmt.Println(languages[i], "\t\t", country[i])
	}
	fmt.Print(languages) // Can also print array directly
}
