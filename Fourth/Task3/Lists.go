package Task3

import "fmt"

func Lists() {
	fmt.Println("Task 3")

	people := make(map[string]int)

	people["Alice"] = 30
	people["Bob"] = 25
	people["Charlie"] = 35
	people["David"] = 28

	fmt.Println("List now:")
	for name, age := range people {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	nameToRemove := "Bob"
	removePerson(people, nameToRemove)

	fmt.Println("\nList after delete:")
	for name, age := range people {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	fmt.Println()
}
