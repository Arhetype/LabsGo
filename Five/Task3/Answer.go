package Task3

import "fmt"

func Answer() {
	fmt.Println("Task 3")

	circle := Circle{Radius: 5.0}

	area := circle.Area()
	fmt.Printf("Area of the circle with Radius %.2f is %.2f\n", circle.Radius, area)

	fmt.Println()
}
