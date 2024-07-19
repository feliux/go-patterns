package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	mutex   sync.Mutex
	wg      sync.WaitGroup
	cond    *sync.Cond
	started bool
)

// worker is a logic construction for doing work.
func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	for !started {
		cond.Wait() // wait for the condition to become true
	}
	// do work
	fmt.Println("Doing some work...")
	mutex.Unlock()
}

func main() {
	// mutex.Lock()
	cond := sync.NewCond(&mutex)
	wg.Add(1)
	go worker(&wg)
	started = true
	// wait a time to see the cond signal
	time.Sleep(time.Second * time.Duration(rand.Intn(2)))
	// signal that the condition is now true
	cond.Signal()
	// cond.Signal() // signal waiting goroutines. No garantees about the order
	// cond.Broadcast() // signal all waiting goroutines
	// mutex.Unlock()
	wg.Wait()
}
