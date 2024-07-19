# Worker pool

Is a pattern to achieve concurrency using fixed number of workers to execute multiple amounts of tasks on a queue. In go ecosystem, we use goroutines to spawn the worker and implement the queue using channels. The defined amounts of workers will pull the task from the queue, and finish up the task, and when the task has been done, the worker will keep pulling the new one until queue is empty.

This approach helps control resource consumption, parallelize work, and efficiently utilize available resources. Keep in mind that the minimal size of a goroutine is 2KB.

```go
// T is a type alias to accept any type
type T = interface{}
// WorkerPool is a contract for worker pool implementation.
// First, Run() will dispatch the worker pool.
// Second, AddTask() will add a task (function to be processed) to the worker pool.
type WorkerPool interface {
	Run()
	AddTask(task func())
}
```

The `sync.Pool` allows to reuse objects (such as worker goroutines) rather than creating new ones, thus reducing the overhead of goroutine creation and improving performance.

### Dynamic Resizing

this pattern allows the worker pool to increase or decrease the number of workers (goroutines) dinamically, optimizing resource utilization based on dmeand.