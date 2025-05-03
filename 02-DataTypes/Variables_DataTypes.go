package main

import "fmt"

func main() {
	// Using var keyword to declare variables
	var name string = "Zaid"
	var age int = 22 //int type (int8, int16, int32, int64 are also available)

	//Using := operator to declare variables
	location := "Lucknow"
	DOB := "2003-08-07" // string type
	c := 1.5555         // float64 type (float32 is also available)
	letter := 'a'       // rune type (char is known as rune in Go)
	isStudent := true   // boolean type

	//printing variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Location:", location)
	fmt.Println("DOB:", DOB)
	fmt.Println("c:", c)
	fmt.Println("isStudent:", isStudent)
	fmt.Println("Letter:", letter)         // Will print the letter as a number (97)
	fmt.Println("Letter:", string(letter)) // Will print the letter as a character (a)
}
