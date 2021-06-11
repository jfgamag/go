package main

import "fmt"

func main() {
	//Declaracion de variables
	hellomessage := "Hello"
	worldmessage := "world"
	// Println
	fmt.Println(hellomessage, worldmessage)
	fmt.Println(hellomessage, worldmessage)
	//printf
	nombre := "Juan"
	apellido := "Gama"
	edad := 28
	fmt.Printf("%s %s es el hijo mayor, tiene %d anios\n", nombre, apellido, edad)

	//Sprintf
	message := fmt.Sprintf("%s %s tiene %d anios", nombre, apellido, edad)
	fmt.Println(message)

	//Tipo de datos
	fmt.Printf("hellomessage: %T\n", hellomessage)
}
