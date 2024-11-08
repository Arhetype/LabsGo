package Task2

import (
	"fmt"
)

func List() {
	fmt.Println("Task 2")

	people := make(map[string]int)

	people["Alice"] = 30
	people["Bob"] = 25
	people["Charlie"] = 35
	people["David"] = 28

	fmt.Printf("Average age: %.2f\n", AverageAge(people))

	fmt.Println()
}
