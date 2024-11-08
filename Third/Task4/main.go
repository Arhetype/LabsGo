package main

import (
	"fmt"
)

func main() {
	var numbers [5]int

	for i := 0; i < 5; i++ {
		fmt.Printf("Enter numbers %d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	fmt.Println("Enterd numbers:")
	for _, number := range numbers {
		fmt.Println(number)
	}
}
