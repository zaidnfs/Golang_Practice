package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Select statement is used to wait on multiple channel operations.
// Basically help us accept the first channel that is ready to send or receive data.
// It is similar to switch statement but for channels.
// It is used to handle multiple channels in a concurrent program.

func main() {
	/*ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second) // Simulate some work
		ch1 <- "Hello from channel 1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Hello from channel 2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg) // This will be printed first because ch2 is ready first
	case msg := <-ch2:
		fmt.Println(msg) // This will be printed second because ch1 is ready second
	}*/

	/*ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second) // Simulate delay
		ch <- "Data recived"
	}()
	select {
	case msg := <-ch:
		fmt.Println("Success:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: No data received within 2 seconds")
	}*/

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		delay := (time.Duration(rand.Intn(3000)) * time.Millisecond)
		time.Sleep(delay) // Simulate some work
		fmt.Println("Sending data from channel 1 after", delay)
		ch1 <- "Data from channel 1"
	}()
	go func() {
		delay := (time.Duration(rand.Intn(3000)) * time.Millisecond)
		time.Sleep(delay) // Simulate some work
		fmt.Println("Sending data from channel 2 after", delay)
		ch2 <- "Data from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}
}
