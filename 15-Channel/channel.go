package main

import "fmt"

func printNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send the number to the channel
	}
	close(ch) // Close the channel after sending all numbers
}

func main() {
	/*ch1 := make(chan string) // Create a channel of type string
	ch2 := make(chan int)
	go func() {
		ch1 <- "Hello from Goroutine!\n" // Send a message to the channel
	}()

	go printNumbers(ch2)

	for num := range ch2 {
		fmt.Println("recived: ", num) // Receive numbers from the channel and print them
	}

	s := <-ch1 // Receive the message from the channel
	fmt.Print("Received: ", s)

	fmt.Println("Main is finished.")*/

	// Create a buffered channel of type int with a capacity of 5
	ch3 := make(chan int, 3)

	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	//ch3 <- 4 // This will block the program because the channel is full

	fmt.Println("Recived", <-ch3)
	fmt.Println("Recived", <-ch3)
	fmt.Println("Recived", <-ch3)
	//fmt.Println("Recived", <-ch3) // This will block the program because the channel is empty
}
