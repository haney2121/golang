package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getAge(birth int, currentYear int) int {
	return currentYear - birth
}

func main() {
	fmt.Println(greeting("Justin"))
	fmt.Println(getAge(1988, 2020))
}
