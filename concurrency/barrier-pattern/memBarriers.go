package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

func main() {
	var iters int = 5
	var counter int32
	wg.Add(2)
	// Goroutine 1: Atomic increment
	go func(iters int) {
		for i := 0; i < iters; i++ {
			atomic.AddInt32(&counter, 1)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(3))) // Simulating some work
		}
		wg.Done()
	}(iters)

	// Goroutine 2: Atomic increment
	go func(iters int) {
		for i := 0; i < iters; i++ {
			atomic.AddInt32(&counter, 1)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(3))) // Simulating some work
		}
		wg.Done()
	}(iters)

	// Wait for goroutines to complete
	// time.Sleep(100 * time.Millisecond)
	wg.Wait()
	// Print the final value of the counter
	finalValue := atomic.LoadInt32(&counter)
	fmt.Println("Final Counter Value:", finalValue)
}
