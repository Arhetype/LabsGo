package main

import (
	"fmt"
	"sync"
)

func fibonacci(n int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func printFibonacci(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go fibonacci(10, &wg, ch)
	go printFibonacci(&wg, ch)

	wg.Wait()
}
