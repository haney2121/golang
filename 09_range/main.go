package main

import "fmt"

// Range is used to loop through arrays, maps, and slices

func main() {

	ids := []int{22, 34, 57, 98, 2, 3}

	//js for of loop _ means not used
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	var total int = 0

	for _, id := range ids {
		total += id
	}

	fmt.Println(total)

	// Range with map
	emails := map[string]string{"Justin": "justin@gmail.com", "Shana": "Shana@gmail", "Chase": "chase@gmail"}

	for key, value := range emails {
		fmt.Printf("%s: %s\n", key, value)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}

}
