package Task6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Numbers() {
	fmt.Println("Task 6")

	fmt.Println("Enter digits separated by space:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	strNumbers := strings.Fields(input)

	var numbers []int

	for _, strNum := range strNumbers {
		num, err := strconv.Atoi(strNum)
		if err == nil {
			numbers = append(numbers, num)
		} else {
			fmt.Printf("Conversion error '%s' to number. Skip.\n", strNum)
		}
	}

	fmt.Print("Answer: ")
	for i := len(numbers) - 1; i >= 0; i-- {
		fmt.Print(numbers[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
