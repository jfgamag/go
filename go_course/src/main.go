package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 float64 = 3.1415
	fmt.Println("pi:", pi)

	// declaracion de variables enteras

	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
	// Area cuadrado

	const base_cuadrado = 10
	area_cuadrado := base_cuadrado * base_cuadrado
	fmt.Println("El area del cuadrado es: ", area_cuadrado)

	//Suma
	x := 10
	y := 4

	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = x - y
	fmt.Println("Resta: ", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplacion: ", result)

	//Division
	result = x / y
	fmt.Println("Division: ", result)

	//Modulo

	result = x % y
	fmt.Println("Modulo: ", result)

	//Incremental
	x++

	//Decremental
	x--
	// Calcular area rectangulo, circulo y trapecio

	base_r := 20
	altura_r := 5
	result = base_r * altura_r
	fmt.Println("El area del rectangulo es: ", result)

	var r float64 = 5

	area_cir := math.Pi * math.Pow(r, 2)
	fmt.Println("El area del circulo es", area_cir)

}
