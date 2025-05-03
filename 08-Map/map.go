package main

import "fmt"

func main() {
	// Map is a collection of key-value pairs.
	// The key is used to access the value.
	// The key must be unique.
	// The value can be of any type.
	mymap := map[string]int{
		"laptop":  1000,
		"phone":   500,
		"tablet":  300,
		"earbuds": 100,
	}
	// Adding a new key-value pair to the map.
	mymap["Charger"] = 50

	//Deleting a value from map
	delete(mymap, "laptop")

	//Updating a value in map
	mymap["earbuds"] = 20

	fmt.Println(mymap)

	// Accessing the value using the key.
	fmt.Println(mymap["laptop"])

	// To check if a key exists in the map.
	value, exist := mymap["phone"]
	fmt.Println(value, exist)

}
