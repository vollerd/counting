package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numGoroutines = 1000
	numIncrements = 10000
)

func mutexCounter() int {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numIncrements; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return counter
}

func segmentedCounter() int {
	segments := make([]int, numGoroutines)
	var wg sync.WaitGroup

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()
			for j := 0; j < numIncrements; j++ {
				segments[index]++
			}
		}(i)
	}
	wg.Wait()

	total := 0
	for _, segment := range segments {
		total += segment
	}
	return total
}

func main() {
	start := time.Now()
	mutexResult := mutexCounter()
	mutexDuration := time.Since(start)

	start = time.Now()
	segmentedResult := segmentedCounter()
	segmentedDuration := time.Since(start)

	fmt.Printf("Mutex-based counter result: %d\n", mutexResult)
	fmt.Printf("Mutex-based counter time: %v\n", mutexDuration)
	fmt.Printf("Segmented counter result: %d\n", segmentedResult)
	fmt.Printf("Segmented counter time: %v\n", segmentedDuration)
}
