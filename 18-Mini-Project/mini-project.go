package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

func logToFile(mu *sync.Mutex, logFile *os.File, logEntry string) {
	mu.Lock()
	defer mu.Unlock()
	logFile.WriteString(logEntry)
	logFile.Sync() // Ensure the log entry is written to the file immediately
}

func checker(id int, jobs <-chan string, wg *sync.WaitGroup, online_websites *[]string, mu *sync.Mutex, logFile *os.File) {
	defer wg.Done()
	logEntry := ""                                   // Initialize logEntry to an empty string
	client := http.Client{Timeout: 10 * time.Second} // Set a timeout for the HTTP request
	for website := range jobs {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		response, err := client.Get(website)
		fmt.Printf("[%s]Worker %d Checking %s...\n", currentTime, id, website)
		if err != nil {
			logEntry = fmt.Sprintf("[%s] %s is Offline (error: %v)\n", currentTime, website, err)
			//fmt.Printf("[%s] %s is Offline\n", currentTime, website)
		} else {
			defer response.Body.Close() // Close the response body to prevent resource leaks

			if response.StatusCode == http.StatusOK {
				mu.Lock()
				*online_websites = append(*online_websites, website)
				mu.Unlock()
				logEntry = fmt.Sprintf("[%s] %s is Online, Response code: %d\n", currentTime, website, response.StatusCode)
				//fmt.Printf("[%s] %s is Online, Response code: %d\n", currentTime, website, response.StatusCode)
			} else {
				logEntry = fmt.Sprintf("[%s] %s is Offline (error: %v)\n", currentTime, website, err)
				//fmt.Printf("[%s] %s is Offline, Response code: %d\n", currentTime, website, response.StatusCode)
			}
		}
		logToFile(mu, logFile, logEntry) // Call the logToFile function to write the log entry to the file
	}

}

func checkingWebsite(logFile *os.File) {
	const maxWorkers = 5
	// Code for reading the file and storing the data in a slice.
	online_websites := []string{}
	websites := []string{}
	jobs := make(chan string, len(websites))
	file, err := os.Open("websites.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Reading file...")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		websites = append(websites, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	// Code for checking if the website is online or offline.
	var wg sync.WaitGroup
	var mu sync.Mutex
	fmt.Println("Checking websites...")
	for i := 1; i <= maxWorkers; i++ {
		wg.Add(1)
		go checker(i, jobs, &wg, &online_websites, &mu, logFile)
	}
	for _, website := range websites {
		jobs <- website
	}
	close(jobs)
	wg.Wait()
	sort.Strings(online_websites)
	fmt.Println("Online websites:\n", online_websites)
	fmt.Println("All websites checked.")
}

func main() {
	// Create a logger to log the output to a file
	logFile, err := os.OpenFile("website_checker.Log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	ticker := time.NewTicker(20 * time.Second) // Change to 1 minute for actual use
	defer ticker.Stop()

	checkingWebsite(logFile) // Initial check
	fmt.Println("Starting website checker in 1 minute...")

	for { // Infinite loop to check every minute
		// Wait for the ticker to tick
		select {
		case <-ticker.C:
			checkingWebsite(logFile)
			fmt.Println("Starting website checker again in 1 minute...")
			// Create a ticker that ticks every minute
		}
	}
}
