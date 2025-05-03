package main

import "fmt"

func main() {
	games := [5]string{"cricket", "football", "hockey", "tennis", "badminton"}
	// Slice is a dynamic array in Go.
	fruits := []string{"apple", "banana", "cherry"}
	number := []int{1, 2, 3, 4, 5}
	// append is used to add elements to a slice.
	// The append function takes a slice and one or more elements to add to the slice.
	number = append(number, 6, 7, 8, 9, 10)
	fruits = append(fruits, "orange", "grape", "kiwi")

	// Creating a slice from an array.
	slice := games[0:3] // 3 is exclusive.
	fmt.Println(fruits)
	fmt.Println(number)
	fmt.Println(slice)
}
