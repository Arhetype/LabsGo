package main

import (
	"Third/stringutils"
	"fmt"
)

func main() {
	original := "Hello, World!"
	reversed := stringutils.Reverse(original)
	fmt.Println("Original:", original)
	fmt.Println("Reverse:", reversed)
}
