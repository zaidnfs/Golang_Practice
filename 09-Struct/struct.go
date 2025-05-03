package main

import "fmt"

//We use the 'type' & 'struct' keyworks to declare a struct
type Book struct {
	title  string
	auther string
	price  float64
}

func main() {
	// Struct is a collection of fields.
	// A struct is a composite data type that groups together variables (fields) under one name.
	// The fields can be of different types.
	// A struct is used to create a complex data type that groups together related data.

	// Struct initialization 1
	var b1 Book
	b1.auther = "J.K. Rowling"
	b1.title = "Harry Potter"
	b1.price = 500
	fmt.Println(b1)

	//Struct initialization 2
	b2 := Book{auther: "xyz", title: "abc", price: 2000}
	fmt.Println(b2)

	//Slice of struct initialization
	books := []Book{
		{title: "Harry Potter", auther: "J.K. Rowling", price: 500},
		{title: "The Hobbit", auther: "J.R.R. Tolkien", price: 300},
		{title: "The Catcher in the Rye", auther: "J.D. Salinger", price: 400},
		{title: "To Kill a Mockingbird", auther: "Harper Lee", price: 600},
		{title: "1984", auther: "George Orwell", price: 700},
		{title: "The Great Gatsby", auther: "F. Scott Fitzgerald", price: 798.51},
	}

	//Iterating over the slice of struct
	for _, book := range books { // _ is used to ignore the index of the slice
		fmt.Printf("%s by %s: $%2.f\n", book.title, book.auther, float64(book.price))
	}
}
