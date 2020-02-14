package main

import "fmt"

// len is js array.length, you wrap your array with it
// range is done via [1:3]

func main() {
	// Arrays have to be fixed length and types
	var fruitArr [2]string
	// Assign values - Array
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"
	// Array Declare and Assign ShortHand
	skillz := [2]string{"React", "GoLang"}

	// Slice is Array without a fixed length and type
	// Slice Declare and Assign ShortHand
	colors := []string{"Red", "Blue", "Yellow", "Pink"}

	fmt.Println(fruitArr)
	fmt.Println(skillz)
	fmt.Println(colors)

	fmt.Println(len(colors))
	// range selection first item to the second
	fmt.Println(colors[1:2])

}
