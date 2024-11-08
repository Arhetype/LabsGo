package Task5

import "fmt"
import "Five/Task4"

func printAreas(shapes []Task4.Shape) {
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}
