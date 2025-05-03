package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func checker(website string, wg *sync.WaitGroup, online_websites *[]string, mu *sync.Mutex) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	status := "Offline"
	fmt.Printf("Checking %s...\n", website)
	if rand.Intn(3) == 1 {
		status = "Online"
		mu.Lock()
		*online_websites = append(*online_websites, website)
		mu.Unlock()
	}
	fmt.Printf("%s is %s\n", website, status)
}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	online_websites := []string{}
	websites := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
	}

	for _, website := range websites {
		wg.Add(1)
		go checker(website, &wg, &online_websites, &mu)
	}
	wg.Wait()
	fmt.Println("All websites checked.")
	fmt.Println("Online websites:\n", online_websites)
}
