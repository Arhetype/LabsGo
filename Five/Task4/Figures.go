package Task4

import "fmt"
import "Five/Task3"

func Figures() {
	fmt.Println("Task 4")

	circle := Task3.Circle{Radius: 5.0}
	rectangle := Rectangle{Width: 4.0, Height: 6.0}

	shapes := []Shape{circle, rectangle}

	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
	fmt.Println()
}
