# Goroutines

## Go Concept

A **goroutine** is a lightweight thread managed by the Go runtime. Goroutines enable concurrent execution of functions. To start a goroutine, use the `go` keyword before a function call.

Goroutines are extremely cheap compared to OS threads:
- They start with a small stack (a few KB) that grows as needed
- Thousands or even millions of goroutines can run in a single program
- The Go runtime multiplexes goroutines onto a small number of OS threads

Goroutines are part of Go's **CSP (Communicating Sequential Processes)** concurrency model, where concurrent processes communicate through channels rather than sharing memory.

### Go Example

```go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// Start a goroutine
	go say("world")
	
	// Run in the main goroutine
	say("hello")
}
```

## C# Equivalent

C# uses **Tasks and async/await** for asynchronous and concurrent operations. C# also has **Thread** for lower-level thread management. However, the paradigms are fundamentally different:

### **Paradigm Difference: CSP vs Async/Await**

**Go (CSP Model):**
- Goroutines are independent, concurrent execution units
- Communication happens through **channels** (explicit message passing)
- "Don't communicate by sharing memory; share memory by communicating"
- Functions don't need special syntax to be concurrent
- Blocking is cheap - goroutines are so lightweight you can block freely

**C# (Async/Await Model):**
- Tasks represent asynchronous operations that will complete in the future
- Methods must be marked `async` to enable await
- Focus is on **non-blocking I/O** rather than parallelism
- Sharing memory with locks is the primary concurrency mechanism
- Blocking is expensive - should be avoided in async code

### When Paradigms Align:
- Starting background work: `go func()` ≈ `Task.Run()`
- I/O-bound operations: Go's blocking goroutines ≈ C# async/await
- CPU-bound parallel work: goroutines ≈ `Parallel.For` or `Task.WhenAll`

### When Paradigms Differ:
- Go: Goroutines + channels for coordination
- C#: async/await for I/O, explicit locking for synchronization

### C# Example

```csharp
using System;
using System.Threading;
using System.Threading.Tasks;

class Program
{
    static void Say(string s)
    {
        for (int i = 0; i < 5; i++)
        {
            Thread.Sleep(100);
            Console.WriteLine(s);
        }
    }

    static async Task Main()
    {
        // Start a background task (similar to goroutine)
        var task = Task.Run(() => Say("world"));
        
        // Run in the main thread
        Say("hello");
        
        // Wait for the background task to complete
        await task;
    }
}
```

## Key Comparison

| Feature | Go (Goroutines) | C# (Tasks/Threads) |
|---------|-----------------|-------------------|
| **Concurrency model** | CSP (message passing) | Async/await (callbacks/continuations) |
| **Syntax** | `go func()` | `Task.Run()` or `async/await` |
| **Weight** | Ultra-lightweight (KB) | Heavier (Thread Pool or OS thread) |
| **Scalability** | Millions of goroutines | Hundreds to thousands of tasks |
| **Blocking** | Cheap, encouraged | Expensive, should avoid in async |
| **Communication** | Channels (explicit) | Shared memory + locks (implicit) |
| **Function marking** | Not needed | `async` keyword required |
| **Waiting** | Channels, WaitGroup | `await`, `Task.Wait()`, events |

## When to Use Each

**Use Go's goroutines** when:
- You need massive concurrency (web servers, network services)
- You want simple concurrent code without callback hell
- Communication between concurrent operations is important

**Use C# async/await** when:
- You have I/O-bound operations (database, HTTP, file I/O)
- You want to keep the UI thread responsive
- You're working in an existing async ecosystem

**Use C# Task.Run/Threads** when:
- You need CPU-bound parallel processing
- You need precise thread control
- You're doing background work separate from async I/O
