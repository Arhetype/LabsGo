package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers(ch chan<- int) {
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		ch <- num
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func checkEvenOdd(ch <-chan int, resultCh chan<- string) {
	for num := range ch {
		if num%2 == 0 {
			resultCh <- fmt.Sprintf("%d is even", num)
		} else {
			resultCh <- fmt.Sprintf("%d is odd", num)
		}
	}
}

func main() {
	numberCh := make(chan int)
	resultCh := make(chan string)

	go generateRandomNumbers(numberCh)
	go checkEvenOdd(numberCh, resultCh)

	for i := 0; i < 10; i++ {
		select {
		case result := <-resultCh:
			fmt.Println(result)
		}
	}
}
