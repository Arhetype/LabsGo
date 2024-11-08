package Task2

import (
	"fmt"
)

func CheckNumber() {

	fmt.Println("Task 2")

	var num int

	fmt.Print("Enter number: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
	fmt.Println()
}
