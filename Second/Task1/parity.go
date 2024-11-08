package Task1

import (
	"fmt"
)

func Parity() {

	fmt.Println("Task 1")
	var (
		digit int
	)

	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &digit)

	if digit%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is not even")
	}
	fmt.Println()
}
