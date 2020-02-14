package main

import (
	"fmt"
	"strconv"
)

// Person defined Struct (Class like)
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	//alt
	firstName, lastName, city, gender string
	age                               int
}

// Greeting Method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// Change Age (pointer reciever) -- no return value
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "Female" {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init Person using struct
	person1 := Person{firstName: "Justin", lastName: "Haney", city: "Daphne", gender: "Male", age: 31}

	// Alternative
	person2 := Person{"Shana", "Smith", "Daphne", "Female", 29}

	//get single field of the struct
	person1.hasBirthday()
	person1.getMarried("Jones")
	fmt.Println(person1.greet())

	fmt.Println(person1, person2)
}
