# Buffered Channels

## Go Concept

Channels can be **buffered**. Provide the buffer length as the second argument to `make` to create a buffered channel:

```go
ch := make(chan int, 100)
```

**Buffered channels** allow sending without an immediate receiver:
- Sends to a buffered channel block only when the buffer is **full**
- Receives from a buffered channel block only when the buffer is **empty**
- Unbuffered channels (buffer size 0) block on every send until received

Buffered channels are useful for:
- Rate limiting and throttling
- Decoupling producers from consumers
- Avoiding unnecessary blocking when you know capacity limits

### Go Example

```go
package main

import "fmt"

func main() {
	// Create a buffered channel with capacity 2
	ch := make(chan int, 2)
	
	// These sends don't block because buffer has space
	ch <- 1
	ch <- 2
	
	// This would block (buffer full) if uncommented:
	// ch <- 3
	
	// Receive values
	fmt.Println(<-ch) // Output: 1
	fmt.Println(<-ch) // Output: 2
}
```

### Practical Example: Worker Pool

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)      // Buffered: can queue all jobs
	results := make(chan int, numJobs)   // Buffered: can collect all results

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs (doesn't block because buffer has space)
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
```

## C# Equivalent

C# **Channel<T>** supports **bounded channels** which are directly equivalent to Go's buffered channels. C# also has **BlockingCollection<T>** with bounded capacity.

Key differences:
- Go's buffer size is fixed at creation
- C#'s BoundedChannelOptions provides more control (drop oldest, drop newest, wait)
- Both support blocking sends when full and blocking receives when empty
- C# uses async/await instead of blocking

### C# Example (Bounded Channel)

```csharp
using System;
using System.Threading.Channels;
using System.Threading.Tasks;

class Program
{
    static async Task Main()
    {
        // Create a bounded channel with capacity 2
        var channel = Channel.CreateBounded<int>(2);
        
        // These sends don't block because buffer has space
        await channel.Writer.WriteAsync(1);
        await channel.Writer.WriteAsync(2);
        
        // This would block (buffer full):
        // await channel.Writer.WriteAsync(3);
        
        // Receive values
        Console.WriteLine(await channel.Reader.ReadAsync()); // Output: 1
        Console.WriteLine(await channel.Reader.ReadAsync()); // Output: 2
    }
}
```

### Practical Example: Worker Pool in C#

```csharp
using System;
using System.Threading;
using System.Threading.Channels;
using System.Threading.Tasks;

class Program
{
    static async Task Worker(int id, ChannelReader<int> jobs, ChannelWriter<int> results)
    {
        await foreach (var job in jobs.ReadAllAsync())
        {
            Console.WriteLine($"Worker {id} started job {job}");
            await Task.Delay(1000); // Simulate work
            Console.WriteLine($"Worker {id} finished job {job}");
            await results.WriteAsync(job * 2);
        }
    }

    static async Task Main()
    {
        const int numJobs = 5;
        var jobs = Channel.CreateBounded<int>(numJobs);      // Buffered
        var results = Channel.CreateBounded<int>(numJobs);   // Buffered

        // Start 3 workers
        var workers = new Task[3];
        for (int w = 0; w < 3; w++)
        {
            int workerId = w + 1;
            workers[w] = Worker(workerId, jobs.Reader, results.Writer);
        }

        // Send 5 jobs (doesn't block because buffer has space)
        for (int j = 1; j <= numJobs; j++)
        {
            await jobs.Writer.WriteAsync(j);
        }
        jobs.Writer.Complete();

        // Wait for all workers to finish
        await Task.WhenAll(workers);
        results.Writer.Complete();

        // Collect results
        await foreach (var result in results.Reader.ReadAllAsync())
        {
            Console.WriteLine($"Result: {result}");
        }
    }
}
```

### C# Example (BlockingCollection - Traditional Approach)

```csharp
using System;
using System.Collections.Concurrent;
using System.Threading;
using System.Threading.Tasks;

class Program
{
    static void Main()
    {
        // Bounded blocking collection with capacity 2
        var queue = new BlockingCollection<int>(boundedCapacity: 2);
        
        // Add items (blocks if full)
        queue.Add(1);
        queue.Add(2);
        
        // This would block:
        // queue.Add(3);
        
        // Take items (blocks if empty)
        Console.WriteLine(queue.Take()); // Output: 1
        Console.WriteLine(queue.Take()); // Output: 2
    }
}
```

## Key Comparison

| Feature | Go (Buffered Channels) | C# (Bounded Channel<T>) | C# (BlockingCollection<T>) |
|---------|------------------------|-------------------------|----------------------------|
| **Creation** | `make(chan T, n)` | `Channel.CreateBounded<T>(n)` | `new BlockingCollection<T>(n)` |
| **Send (blocks when full)** | `ch <- v` | `await WriteAsync(v)` | `Add(v)` |
| **Receive (blocks when empty)** | `v := <-ch` | `await ReadAsync()` | `Take()` |
| **Capacity** | Fixed at creation | Fixed at creation | Fixed at creation |
| **Behavior when full** | Block sender | Configurable (wait/drop) | Block sender |
| **Async support** | N/A (blocking is cheap) | Native async/await | Via TryAdd with timeout |

## When to Use Buffered Channels

**Go:**
- When you want to decouple send and receive timing
- To implement rate limiting (buffer = rate limit)
- For worker pools with known job capacity
- When you know the maximum items in flight

**C#:**
- Producer-consumer with known capacity
- Backpressure handling
- Avoiding unbounded memory growth
- Throttling async operations

## Buffer Size Guidelines

**Small buffers (1-10):**
- Message passing between few goroutines
- Simple coordination

**Medium buffers (10-1000):**
- Worker pools
- Request queues
- Batch processing

**Large buffers (1000+):**
- High-throughput pipelines
- Decoupling fast producers from slow consumers
- Consider if unbounded queue might be better
