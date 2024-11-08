package Task6

import (
	"fmt"
)

func Digits() {
	fmt.Println("Task 6")

	var (
		a, b int
	)

	fmt.Println("Enter a: ")
	fmt.Scan(&a)
	fmt.Println("Enter b: ")
	fmt.Scan(&b)

	num := (a + b) / 2

	fmt.Println("Answer: ", num)
}
