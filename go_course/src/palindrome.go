package main

import (
	"fmt"
	"strings"
)

func ispalindrome(text string) {
	var textreverse string

	for i := len(text) - 1; i >= 0; i-- {
		textreverse += string(text[i])
	}
	if text == textreverse {
		fmt.Printf(text + " " + "is palindrome")
	} else {
		fmt.Printf(text + "is not palindrome")
	}

}

func main() {
	ispalindrome(strings.ToLower("Ana"))
}
