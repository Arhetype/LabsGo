package main

import (
	"fmt"
)

func main() {
	strings := []string{
		"Hellow",
		"How are you?",
		"Go is forever.",
		"I love dance.",
		"Good luck!",
	}

	var longestString string

	for _, str := range strings {
		if len(str) > len(longestString) {
			longestString = str
		}
	}

	fmt.Println("The biggest one:", longestString)
}
