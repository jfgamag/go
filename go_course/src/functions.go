package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)

}

func tripleArg(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnvalue(a int) int {
	return a * 2
}

func doublereturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hello world")
	tripleArg(1, 2, "3")
	value := returnvalue(2)
	fmt.Println("Value: ", value)

	value1, _ := doublereturn(2)
	fmt.Println("Value1:", value1)
}
