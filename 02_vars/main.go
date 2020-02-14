package main

// https://golang.org/pkg/fmt/
import "fmt"

// ===== Vars can be defined outside of functions Globally ===
var name string = "Justin"

func main() {
	//Main Types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complext64 complex128

	// Creating Variable using var - can declare types
	// var name string = "Justin"
	var age int = 31
	const isCool bool = true

	// Creating variable shorthand - Only in functions - Cant not declare types
	// double assigning shorthand
	job := "Developer"
	size := 1.4

	lastName, email := "Haney", "haney037332@gmail.com"

	fmt.Println(name, age, isCool, job, lastName, email)

	//type of value with %T
	fmt.Printf("%T\n", size)

}
