package Task2

import "fmt"

func Lists() {
	fmt.Println("Task 2")

	person := Person{name: "Alice", age: 30}
	fmt.Println(person.Info())
	person.Birthday()
	fmt.Println(person.Info())

	fmt.Println()
}
