package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
	// for forever
	counter_fv := 0
	for {
		fmt.Println(counter_fv)
		counter_fv++
	}
}
