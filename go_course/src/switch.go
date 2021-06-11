package main

import "fmt"

func main() {
	module := 4 % 2
	switch module {
	case 0:
		fmt.Println("It's even")
	default:
		fmt.Println("It's odd")
	}

	// Sin condicion

	value := 200
	switch {
	case value > 100:
		fmt.Println("It's bigger than 100")
	case value < 0:
		fmt.Println("It's less than 0")
	default:
		fmt.Println("No condition")
	}
}
