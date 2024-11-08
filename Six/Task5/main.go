package main

import (
	"fmt"
	"sync"
)

type Request struct {
	Operation string
	Operand1  float64
	Operand2  float64
	Result    chan float64
}

func calculator(requests <-chan Request, wg *sync.WaitGroup) {
	defer wg.Done()
	for req := range requests {
		var result float64
		switch req.Operation {
		case "+":
			result = req.Operand1 + req.Operand2
		case "-":
			result = req.Operand1 - req.Operand2
		case "*":
			result = req.Operand1 * req.Operand2
		case "/":
			if req.Operand2 != 0 {
				result = req.Operand1 / req.Operand2
			} else {
				fmt.Println("Error: Division by zero")
				result = 0
			}
		default:
			fmt.Println("Error: Unknown operation")
			result = 0
		}
		req.Result <- result
	}
}

func main() {
	requests := make(chan Request)
	var wg sync.WaitGroup

	wg.Add(1)
	go calculator(requests, &wg)

	for i := 0; i < 5; i++ {
		resultChan := make(chan float64)
		req := Request{
			Operation: "+",
			Operand1:  float64(i),
			Operand2:  float64(i + 1),
			Result:    resultChan,
		}
		requests <- req

		result := <-resultChan
		fmt.Printf("Result of %f + %f = %f\n", req.Operand1, req.Operand2, result)
	}

	close(requests)
	wg.Wait()
}
