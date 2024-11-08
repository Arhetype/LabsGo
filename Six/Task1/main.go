package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func factorial(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	time.Sleep(1 * time.Second)
	fmt.Println("fact:", fact)
}

func randomNumbers(count int, wg *sync.WaitGroup) {
	defer wg.Done()
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(100)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("random numbers", numbers)
}

func sumSeries(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	time.Sleep(3 * time.Second)
	fmt.Println("sum:", sum)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go factorial(5, &wg)
	go randomNumbers(5, &wg)
	go sumSeries(10, &wg)

	wg.Wait()
}
