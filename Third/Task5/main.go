package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[:]

	fmt.Println("First slice:", slice)

	slice = append(slice, 6)
	fmt.Println("After entered element 6:", slice)

	indexToRemove := 2
	slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
	fmt.Println("After deleted element with index 2:", slice)

	slice = append(slice, 7)
	fmt.Println("After entered element:", slice)
}
