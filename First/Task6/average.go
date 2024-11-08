package Task6

import (
	"fmt"
)

func AverageValue() {
	fmt.Println("Task 6")
	var (
		digit1, digit2, digit3 int
	)
	fmt.Println("Enter 3 numbers:")
	fmt.Scanf("%d %d %d", &digit1, &digit2, &digit3)

	operation := (digit1 + digit2 + digit3) / 3
	fmt.Println("Answer: ", operation)
}
