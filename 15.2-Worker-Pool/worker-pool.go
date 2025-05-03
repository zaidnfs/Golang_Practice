package main

import (
	"fmt"
	"sync" // For WaitGroup
	"time" // For Delay
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		results <- job * 2      // Send the result back to the results channel
	}
}

func square(number int, wg *sync.WaitGroup) {
	defer wg.Done()              // Decrement the counter when the goroutine completes
	time.Sleep(time.Second)      // Simulate work
	fmt.Println(number * number) // Print the square of the number)

}

func main() {
	/*
		// Create a buffered channel of type int with a capacity of 5
		jobs := make(chan int, 5)
		results := make(chan int, 5)

		for i := 1; i <= 3; i++ {
			go worker(i, jobs, results) // Start 3 worker goroutines
		}

		for j := 1; j <= 5; j++ {
			jobs <- j // Send the job to the channel
		}
		close(jobs) // Close the jobs channel to signal that no more jobs will be sent

		for r := 1; r <= 5; r++ { // Receive results from the results channel
			fmt.Println("Recived: ", <-results)
		}*/

	//---------------------------------------------------------.
	//WaitGroup

	var wg sync.WaitGroup // Create a WaitGroup to wait for all goroutines to finish

	numbers := []int{1, 2, 3, 4, 5}

	for _, number := range numbers {
		wg.Add(1)              // Add a count to the WaitGroup for each goroutine
		go square(number, &wg) // Start a goroutine to square the number)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All done...")
}
