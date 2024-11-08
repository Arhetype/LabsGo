package Task3

import "math"

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
