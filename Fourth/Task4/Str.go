package Task4

import (
	"fmt"
	"strings"
)

func Str() {
	fmt.Println("Task 4")

	var input string

	fmt.Print("Enter string: ")
	fmt.Scanln(&input)

	upperInput := strings.ToUpper(input)

	fmt.Println("Up register:", upperInput)
	fmt.Println()
}
