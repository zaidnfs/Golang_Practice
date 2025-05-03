package main

import "fmt"

func main() {
	// In go, the for loop is the only loop construct available.
	j := 0
	// for loop with initialization, condition, and increment
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Print("\n")
	//While loop using for loop
	for j < 10 {
		fmt.Print(j, " ")
		j++
	}
	fmt.Print("\n")
	for k := 0; k < 20; k++ {
		if k%2 == 0 {
			continue // skip even numbers
		} else if k == 15 {
			break // break the loop when k is 15
		} else {
			fmt.Print(k, " ") // print odd numbers
		}
	}
}
