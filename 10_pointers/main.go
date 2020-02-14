package main

import "fmt"

// Points to a reference in Memory on the system
// * means pointer

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read val from address otherwise its the memory address
	fmt.Println(*b)

	//Change value with pointer
	*b = 10
	//a was changed when the b reference was changed
	fmt.Println(a)

}
