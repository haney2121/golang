package main

import "fmt"

// %d is a number placeholder followed by the incoming args for the values

func main() {
	x := 10
	y := 10

	// If Else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// Else If
	color := "Red"

	if color == "Blue" {
		fmt.Println("Color is blue")
	} else if color == "Red" {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Color is not blue or red")
	}

	// Switch

	switch color {
	case "Red":
		fmt.Println("color is red")
	case "Blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is a mystery")
	}

}
