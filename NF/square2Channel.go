package square2Channel

//Problem: Write a Go program that calculates the sum of squares of numbers in a slice using multiple goroutines. Use channels for communication and a WaitGroup to wait for all goroutines to finish.

import (
	"fmt"
	"sync"
)

func square(num int) int {
	return num * num
}

func worker(numbers <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion when the goroutine exits
	for number := range numbers {
		results <- square(number)
	}
}

func square2Channel() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numGoroutines := 4
	numbersChan := make(chan int, len(numbers))
	resultsChan := make(chan int, len(numbers))
	var wg sync.WaitGroup

	// Launch worker goroutines
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go worker(numbersChan, resultsChan, &wg)
	}

	// Send numbers to the channel
	for _, number := range numbers {
		numbersChan <- number
	}
	close(numbersChan) // Close the channel to signal no more data

	// Wait for all worker goroutines to finish
	go func() {
		wg.Wait()
		close(resultsChan) // Close the results channel after all workers are done
	}()

	// Collect results
	totalSum := 0
	for result := range resultsChan {
		fmt.Println("squares:", result)
		totalSum += result
	}

	fmt.Println("Sum of squares:", totalSum)
}
