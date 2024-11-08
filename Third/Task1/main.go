package main

import (
	"Third/mathutils"
	"fmt"
)

func main() {
	number := 5
	result := mathutils.Factorial(number)
	fmt.Printf("Factorial %d is %d\n", number, result)
}
