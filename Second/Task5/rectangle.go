package Task5

import (
	"fmt"
)

type Rectangle struct {
	a int
	b int
}

func (r Rectangle) Area() {
	fmt.Println("Task 5")

	r.a = 4
	r.b = 5

	fmt.Println("Area: ", r.a*r.b)
	fmt.Println()
}
