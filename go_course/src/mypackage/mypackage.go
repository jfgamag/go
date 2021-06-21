package mypackage

import "fmt"

//CarPublic Car with public access
type CarPublic struct {
	Brand  string
	Model  string
	Motor  string
	Weight string
	Year   int
}

type carPrivate struct {
	brand string
	year  int
}

func Printmessage(text string) {
	fmt.Println(text)
}
