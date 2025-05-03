package main

import "fmt"

// Can only Send
func sender(ch chan<- string) {
	ch <- "Hello from sender!"
}

// Can only Receive
func reciver(ch <-chan string) {
	fmt.Print("Recived: ", <-ch)
}

func main() {
	ch1 := make(chan string)

	go sender(ch1)
	reciver(ch1)
}
