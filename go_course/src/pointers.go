package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (compt *pc) duplicateRAM() {
	compt.ram = compt.ram * 2
}

func (compt pc) ping() {
	fmt.Println(compt.brand, "Pong")
}

func main() {
	a := 50
	b := &a

	fmt.Println(a, b)
	fmt.Println(a, *b)

	myPC := pc{ram: 16, disk: 1024, brand: "msi"}
	fmt.Println(myPC)
	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()
	fmt.Println(myPC)
}
