package main

import "fmt"

func main() {
	value1 := 2
	value2 := 5
	if value1 == 1 {
		fmt.Println("It's one")
	} else {
		fmt.Println("It's not one")
	}
	//with and
	if value1 == 2 && value2 == 5 {
		fmt.Println("It's true")
	} else {
		fmt.Println("It's False")
	}
	if value1 == 3 || value2 == 5 {
		fmt.Println("That is correct")
	} else {
		fmt.Println("That is wrong")
	}

	// Even numbers
	if value1%2 == 0 {
		fmt.Printf("The value %d is an odd number", value1)
	} else {
		fmt.Printf("The value %d is even", value1)
	}
}
