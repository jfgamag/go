package main

import "fmt"

type car struct {
	brand  string
	model  string
	motor  string
	weight string
	year   int
}

func main() {
	myCar := car{brand: "Ford", model: "F-150", motor: "4.5 L", weight: "5 ton", year: 2020}
	fmt.Println(myCar)

	//Alternative

	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
