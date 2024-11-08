package Task1

import (
	"fmt"
)

func People() {
	fmt.Println("Task 1")

	people := make(map[string]int)

	people["Alice"] = 30
	people["Bob"] = 25
	people["Charlie"] = 35

	addPerson := func(name string, age int) {
		people[name] = age
	}

	addPerson("Daniil", 20)

	fmt.Println("List of people:")
	for name, age := range people {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	fmt.Println()
}
