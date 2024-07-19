package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task represents a task pending to do.
type Task struct {
	ID int
}

// worker is a logic construction for doing tasks.
func worker(id int, tasks <-chan Task, results chan<- Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task.ID)
		// Simulated work by sleeping for some time
		time.Sleep(time.Second * time.Duration(rand.Intn(3)))
		results <- task // Sending back the processed task to results channel
	}
}

func main() {
	numTasks := 10
	numWorkers := 3
	tasks := make(chan Task, numTasks)
	results := make(chan Task, numTasks)
	var wg sync.WaitGroup
	// Create workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}
	// Add tasks to the tasks channel
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i}
	}
	close(tasks) // Close tasks channel to indicate no more tasks
	// Wait for all workers to finish
	wg.Wait()
	close(results) // Close results channel after all workers are done
	// Collect and print processed tasks
	for result := range results {
		fmt.Println("Processed:", result.ID)
	}
}
