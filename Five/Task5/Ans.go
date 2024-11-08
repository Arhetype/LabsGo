package Task5

import (
	"Five/Task4"
	"fmt"
)
import "Five/Task3"

func Ans() {
	fmt.Println("Task 5")

	circle := Task3.Circle{Radius: 5.0}
	rectangle := Task4.Rectangle{Width: 4.0, Height: 6.0}

	shapes := []Task4.Shape{circle, rectangle}

	printAreas(shapes)

	fmt.Println()
}
