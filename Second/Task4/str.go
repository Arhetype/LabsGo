package Task4

import (
	"fmt"
)

func Str() {
	fmt.Println("Task 4")

	var str string

	fmt.Println("Enter string:")
	fmt.Scan(&str)

	fmt.Println("Length: ", len(str))
	fmt.Println()
}
