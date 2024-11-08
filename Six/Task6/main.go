package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func worker(id int, tasks <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		reversed := reverseString(task)
		results <- fmt.Sprintf("Worker %d: %s", id, reversed)
	}
}

func main() {
	var numWorkers int
	fmt.Print("Enter the number of workers: ")
	_, err2 := fmt.Scanln(&numWorkers)
	if err2 != nil {
		return
	}

	tasks := make(chan string, 10)
	results := make(chan string, 10)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, tasks, results, &wg)
	}

	file, err := os.Open("C:\\code\\StudyD\\PP\\Six\\Task6\\input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
