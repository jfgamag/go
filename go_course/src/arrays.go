package main

import "fmt"

func main() {
	//Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))
	//slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//metodos en el slice
	fmt.Print(slice[0])
	fmt.Println(slice[0:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//append
	slice = append(slice, 7)
	fmt.Println(slice)
	//append new list
	new_slice := []int{8, 9, 10}
	slice = append(slice, new_slice...)
	fmt.Println(slice)

	// Recorriendo slices

	slice2 := []string{"hola", "que", "hace"}
	for i, value := range slice2 {
		fmt.Println(i, value)
	}
}
