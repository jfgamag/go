package main

import (
	"fmt"

	mypackage "./mypackage"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2021
	fmt.Println(myCar)

	mypackage.Printmessage("Hello everyone")

}
