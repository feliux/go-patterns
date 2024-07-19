# Fan-out Fan-in

Is a concurrent designt pattern used to distribute work among multiple goroutines (fan-out) and then aggregate the results from these goroutines (fan-in). It is a way to diverge and converge data into a single data stream from multiple streams or from one stream to multiple streams or pipelines.

This pattern is commonly employed in go to handle concurrent processing and efficiently utilize available resources. It is particularly useful when you have a time consuming task that can be divided into a smaller subtasks that can be executed concurrently.

In the fan-out stage, a single task is divided into a multiple smaller subtasks, which are then executed concurrently. each subtask can be assigned to a separate goroutine to run concurrently. this stage distributes the workload across multiple goroutines.

In the fan-in stage, the results or outputs from all the concurrently executed subtasks are collected and combined into a single result. This stage waits for all the subtasks to complete and aggregates their results. this stage can also handle synchronization and coordination between the goroutines to ensure that all results are collected before proceeding.

**Fan-Out**

```go
// Split a channel into n channels that receive messages in a round-robin fashion.
func Split(ch <-chan int, n int) []<-chan int {
	cs := make([]chan int)
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}

	// Distributes the work in a round robin fashion among the stated number
	// of channels until the main channel has been closed. In that case, close
	// all channels and return.
	distributeToChannels := func(ch <-chan int, cs []chan<- int) {
		// Close every channel when the execution ends.
		defer func(cs []chan<- int) {
			for _, c := range cs {
				close(c)
			}
		}(cs)

		for {
			for _, c := range cs {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}

					c <- val
				}
			}
		}
	}

	go distributeToChannels(ch, cs)

	return cs
}
```

**Fan-In**

```go
// Merge different channels in one channel
func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	// Start an send goroutine for each input channel in cs. send
	// copies values from c to out until c is closed, then calls wg.Done.
	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	// Start a goroutine to close out once all the send goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
```
