package main

import (
	"Third/mathutils"
	"fmt"
)

func main() {

	var number int

	fmt.Println("Enter number: ")
	fmt.Scan(&number)

	result := mathutils.Factorial(number)

	fmt.Println("Factorial: ", result)
}
