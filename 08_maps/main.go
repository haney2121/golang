package main

import "fmt"

// Maps are Key Value Pairs

func main() {
	// Define Map
	//make map [string] = key  string = value
	emails := make(map[string]string)

	// Declare map and add key Values
	lastNames := map[string]string{"Justin": "Haney", "Shana": "Haney", "Chase": "Harrison"}

	// Assign keyValues
	emails["Justin"] = "haney037332@gmail.com"
	emails["Shana"] = "shana@gmail.com"
	emails["Caden"] = "caden@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Justin"])

	// Delete from map
	delete(emails, "Caden")
	fmt.Println(emails)
	fmt.Println(lastNames)

}
