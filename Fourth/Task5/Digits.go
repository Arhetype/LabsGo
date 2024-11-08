package Task5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Digits() {
	fmt.Println("Task 5")

	var (
		sum float64
	)

	fmt.Println("Enter digits separated by space:")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	numbers := strings.Fields(str)

	for _, numStr := range numbers {
		number, err := strconv.ParseFloat(numStr, 64)
		if err == nil {
			sum += number
		} else {
			fmt.Printf("Conversion error '%s' to number. Skip.\n", numStr)
		}
	}
	fmt.Printf("Sum: %.2f\n", sum)
	fmt.Println()
}
