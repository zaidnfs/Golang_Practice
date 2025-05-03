package main

import (
	"fmt"
	"os"
)

func main() {

	// Writing to a file in Go
	/*content := "Hello!!! Writing to a file in Go.\n"

	err := os.WriteFile("writefromGo.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File written successfully.")*/

	// Reading from a file in Go
	content, err := os.ReadFile("writefromGo.txt")
	if err != nil {
		fmt.Print("Error reading file:", err)
		return
	}
	fmt.Printf("File content: %s\n", content)
}
