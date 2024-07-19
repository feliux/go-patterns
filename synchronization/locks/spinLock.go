package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// SpinLock represents a simple spin lock.
type SpinLock struct {
	flag int32
}

// Lock acquires the lock.
func (s *SpinLock) Lock() {
	for !s.tryLock() {
		runtime.Gosched() // Yield the processor to other goroutines
	}
}

// Unlock releases the lock.
func (s *SpinLock) Unlock() {
	atomic.StoreInt32(&s.flag, 0)

}

// tryLock attempts to acquire the lock.
func (s *SpinLock) tryLock() bool {
	return atomic.CompareAndSwapInt32(&s.flag, 0, 1)
}

func main() {
	var (
		counter   int
		spinLock  SpinLock
		waitGroup sync.WaitGroup
	)
	const numGoroutines = 10
	waitGroup.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer waitGroup.Done()
			spinLock.Lock()
			defer spinLock.Unlock()
			// Simulate some work
			time.Sleep(time.Millisecond)
			// Increment the counter safely
			counter++
		}()
	}
	waitGroup.Wait()
	fmt.Println("Final Counter Value:", counter)
}
