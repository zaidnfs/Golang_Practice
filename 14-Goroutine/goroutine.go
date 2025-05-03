package main

import (
	"fmt"
	"time"
)

// Goroutines are lightweight threads managed by the Go runtime.
// They are used to perform concurrent tasks in Go programs.(parallel execution)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second) // Sleep for 1 second
	}
}

func printAlphabet() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second) // Sleep for 1 second
	}
}

func main() {
	go printNumbers() // Start the goroutines
	go printAlphabet()
	fmt.Println("Main is running...")
	time.Sleep(5 * time.Second) // Sleep for 3 seconds to allow goroutine to finish
}
