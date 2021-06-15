package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Juan"] = 28
	m["Vale"] = 21

	fmt.Println(m)
	// Recorrer map

	for i, v := range m {
		fmt.Printf(i, v)
	}
	//Encontrar valores

	value := m["Jose"]
	fmt.Println(value)

	// Verificar que la llave este o no en el diccionario

	value1, ok := m["Jose"]
	fmt.Println(value1, ok)
}
