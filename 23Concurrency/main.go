// Concurrency (Goroutines & Channels)
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 1. Basic Goroutine Example
func sayHello(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Hello %s! (iteration %d)\n", name, i+1)
		time.Sleep(100 * time.Millisecond)
	}
}

// 2. Function with channels
func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Printf("Consuming: %d\n", value)
		time.Sleep(300 * time.Millisecond)
	}
}

// 3. Select statement example
func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine for channel 1
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	// Goroutine for channel 2
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	// Select statement waits for multiple channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout!")
		}
	}
}

// 4. Worker Pool pattern
type Job struct {
	ID   int
	Work int
}

type Result struct {
	JobID int
	Sum   int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		// Simulate work
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		result := Result{
			JobID: job.ID,
			Sum:   job.Work + job.Work,
		}
		results <- result
	}
}

func workerPoolExample() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	for i := 1; i <= numJobs; i++ {
		jobs <- Job{ID: i, Work: i * 10}
	}
	close(jobs)

	// Wait for workers to finish
	wg.Wait()
	close(results)

	// Collect results
	fmt.Println("Results:")
	for result := range results {
		fmt.Printf("Job %d: Sum = %d\n", result.JobID, result.Sum)
	}
}

// 5. Buffered vs Unbuffered channels
func bufferedChannels() {
	fmt.Println("\n=== Buffered vs Unbuffered Channels ===")

	// Unbuffered channel (synchronous)
	unbuffered := make(chan int)
	go func() {
		fmt.Println("Sending to unbuffered channel...")
		unbuffered <- 42
		fmt.Println("Sent to unbuffered channel!")
	}()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Receiving from unbuffered channel...")
	value := <-unbuffered
	fmt.Printf("Received: %d\n", value)

	// Buffered channel (asynchronous)
	buffered := make(chan int, 3)
	fmt.Println("\nSending to buffered channel...")
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Println("All values sent to buffered channel!")

	fmt.Println("Receiving from buffered channel...")
	fmt.Printf("Received: %d\n", <-buffered)
	fmt.Printf("Received: %d\n", <-buffered)
	fmt.Printf("Received: %d\n", <-buffered)
}

// 6. Channel directions
func sender(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)
}

func receiver(ch <-chan int) {
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
}

func channelDirections() {
	fmt.Println("\n=== Channel Directions ===")
	ch := make(chan int)
	go sender(ch)
	receiver(ch)
}

func main() {
	fmt.Println("=== Go Concurrency Examples ===")

	// 1. Basic Goroutines
	fmt.Println("\n1. Basic Goroutines:")
	go sayHello("Alice")
	go sayHello("Bob")
	time.Sleep(2 * time.Second) // Wait for goroutines to finish

	// 2. Producer-Consumer with channels
	fmt.Println("\n2. Producer-Consumer Pattern:")
	ch := make(chan int)
	go producer(ch)
	consumer(ch)

	// 3. Select statement
	fmt.Println("\n3. Select Statement:")
	selectExample()

	// 4. Worker Pool
	fmt.Println("\n4. Worker Pool:")
	workerPoolExample()

	// 5. Buffered vs Unbuffered channels
	bufferedChannels()

	// 6. Channel directions
	channelDirections()

	fmt.Println("\n=== Concurrency Examples Complete ===")
}

/*
Key Concepts:

1. GOROUTINES:
- Lightweight threads managed by Go runtime
- Started with 'go' keyword
- Run concurrently with other goroutines
- Much cheaper than OS threads

2. CHANNELS:
- Typed conduits for data between goroutines
- Communication by sharing memory (vs sharing memory by communicating)
- Can be buffered or unbuffered
- Safe for concurrent use

3. SELECT STATEMENT:
- Wait on multiple channel operations
- Executes first ready case
- Can include timeout with time.After()
- Essential for handling multiple concurrent operations

4. WORKER POOL PATTERN:
- Fixed number of worker goroutines
- Jobs distributed via channel
- Results collected via another channel
- Efficient resource utilization

5. CHANNEL DIRECTIONS:
- chan<- : send-only channel
- <-chan : receive-only channel
- Compile-time safety for channel usage

6. SYNCHRONIZATION:
- sync.WaitGroup for waiting multiple goroutines
- Mutex for protecting shared state
- Atomic operations for simple counters

Best Practices:
- Don't communicate by sharing memory, share memory by communicating
- Use buffered channels when you know the capacity needed
- Always close channels when done sending
- Use select for timeout and cancellation
- Prefer channels over mutexes when possible
*/
