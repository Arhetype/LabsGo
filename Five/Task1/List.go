package Task1

import "fmt"

func List() {
	fmt.Println("Task 1")

	person := Person{name: "Alice", age: 30}

	fmt.Println(person.Info())
	fmt.Println()
}
