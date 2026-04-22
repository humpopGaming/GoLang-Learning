# Channels

## Go Concept

**Channels** are typed conduits through which you can send and receive values with the channel operator `<-`. Channels are the primary way goroutines communicate with each other and synchronize their execution.

```go
ch <- v    // Send v to channel ch
v := <-ch  // Receive from ch and assign to v
```

Key properties:
- Channels are **blocking** by default: a send blocks until another goroutine receives, and a receive blocks until data is available
- Channels must be created using `make()` before use
- Channels provide **synchronization** for free - no explicit locks needed
- Channels embody Go's philosophy: "Don't communicate by sharing memory; share memory by communicating"

### Go Example

```go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // Send sum to channel
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // Receive from channel (blocks until data available)

	fmt.Println(x, y, x+y) // Output: -5 17 12 (order may vary)
}
```

## C# Equivalent

C# has **Channel<T>** (System.Threading.Channels, .NET Core 2.1+) which is directly inspired by Go's channels. However, C# also has older patterns like **BlockingCollection<T>** and manual synchronization primitives.

### **Concurrency Model Difference:**

**Go:**
- Channels are the **idiomatic** way to communicate between goroutines
- Blocking on channels is natural and efficient
- CSP model: communicate through message passing

**C#:**
- Multiple options: Channel<T>, BlockingCollection<T>, concurrent collections
- Channels are less common; most code uses async/await or locks
- Async model: coordinate through Tasks and shared state

### When Paradigms Align:
- Producer-consumer patterns work similarly
- Both can pass messages between concurrent operations
- Both support bounded/unbounded queues

### When Paradigms Differ:
- Go: Channels are primary; used everywhere
- C#: Channels are specialized; async/await is primary
- Go: Blocking is cheap
- C#: Blocking async code causes issues (deadlocks, thread pool starvation)

### C# Example (Using Channel<T>)

```csharp
using System;
using System.Linq;
using System.Threading.Channels;
using System.Threading.Tasks;

class Program
{
    static async Task Sum(int[] array, Channel<int> channel)
    {
        int sum = array.Sum();
        await channel.Writer.WriteAsync(sum); // Send to channel
    }

    static async Task Main()
    {
        int[] s = { 7, 2, 8, -9, 4, 0 };
        
        var channel = Channel.CreateUnbounded<int>();
        
        // Start two concurrent tasks
        var task1 = Sum(s[..(s.Length / 2)], channel);
        var task2 = Sum(s[(s.Length / 2)..], channel);
        
        // Receive from channel (awaits until data available)
        int x = await channel.Reader.ReadAsync();
        int y = await channel.Reader.ReadAsync();
        
        Console.WriteLine($"{x} {y} {x + y}"); // Output: -5 17 12 (order may vary)
        
        await Task.WhenAll(task1, task2);
    }
}
```

### C# Example (Traditional Approach with Task)

```csharp
using System;
using System.Linq;
using System.Threading.Tasks;

class Program
{
    static int Sum(int[] array)
    {
        return array.Sum();
    }

    static async Task Main()
    {
        int[] s = { 7, 2, 8, -9, 4, 0 };
        
        // Start two concurrent tasks that return values
        var task1 = Task.Run(() => Sum(s[..(s.Length / 2)]));
        var task2 = Task.Run(() => Sum(s[(s.Length / 2)..]));
        
        // Await results
        int[] results = await Task.WhenAll(task1, task2);
        
        Console.WriteLine($"{results[0]} {results[1]} {results[0] + results[1]}");
    }
}
```

## Key Comparison

| Feature | Go (Channels) | C# (Channel<T>) | C# (Traditional) |
|---------|---------------|-----------------|------------------|
| **Idiomatic?** | Yes, primary pattern | No, specialized | Yes (Task-based) |
| **Send** | `ch <- v` (blocks) | `await Writer.WriteAsync(v)` | Return from Task |
| **Receive** | `v := <-ch` (blocks) | `await Reader.ReadAsync()` | `await task` |
| **Creation** | `make(chan T)` | `Channel.CreateUnbounded<T>()` | N/A |
| **Blocking behavior** | Synchronous, cheap | Asynchronous (await) | Asynchronous (await) |
| **Direction** | Bidirectional by default | Separate reader/writer | N/A |
| **Use cases** | All goroutine communication | Specialized scenarios | General async work |

## Comparison to Other C# Patterns

| Pattern | Description | Similarity to Go Channels |
|---------|-------------|--------------------------|
| **Channel<T>** | Modern bounded/unbounded queue | Very similar, inspired by Go |
| **BlockingCollection<T>** | Thread-safe collection with blocking | Similar but older, thread-based |
| **Task<T>** | Async operation with single result | Different: represents future value, not communication |
| **ConcurrentQueue<T>** | Lock-free queue | Similar but non-blocking, requires manual sync |
| **Manual locks** | Monitor, Semaphore, lock | Opposite philosophy: shared memory |

## Choosing Between Patterns

**Use Go channels** when:
- Coordinating between goroutines
- Implementing pipelines or workflows
- Need synchronization with communication
- Want clean, expressive concurrent code

**Use C# Channel<T>** when:
- Building producer-consumer patterns
- Need backpressure or bounded buffering
- Porting Go-style concurrent code to C#

**Use C# Task/async/await** when:
- Doing I/O operations
- Single result from async work
- Working with existing async APIs
- Need UI responsiveness
