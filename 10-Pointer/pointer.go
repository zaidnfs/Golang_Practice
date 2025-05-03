package main

import "fmt"

//  Pointers with functions calling by reference
func updateValue(a *float64) {
	*a = *a / 10
}

func main() {
	a := 110.0
	p := &a                          // var p *float64 = &a
	fmt.Println("Value of a:", a)    // 7
	fmt.Println("Address of a:", &a) // Address of a
	fmt.Println("Value of p:", p)    // Address of a
	fmt.Println("Value of *p:", *p)  // 7
	//*p = 100
	fmt.Println("New value of a:", a) // 100
	fmt.Println("Address of a:", &a)  // Address of a stay the same

	updateValue(&a)
	fmt.Print("Updated value of a:", a)
}
