# Select

## Go Concept

The **`select`** statement lets a goroutine wait on multiple communication operations. It blocks until one of its cases can proceed, then executes that case. If multiple cases are ready, it chooses one at random.

```go
select {
case msg := <-ch1:
    // Use msg from ch1
case msg := <-ch2:
    // Use msg from ch2
case ch3 <- value:
    // Send value to ch3
}
```

`select` is fundamental to Go's concurrency model:
- Coordinate between multiple channels
- Implement timeouts
- Non-blocking operations (with `default`)
- Build complex concurrent patterns

### Go Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Goroutine sends to c1 after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	// Goroutine sends to c2 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Wait for both messages
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
```

### Select with Timeout

```go
func main() {
	c := make(chan string)
	
	go func() {
		time.Sleep(2 * time.Second)
		c <- "result"
	}()
	
	select {
	case res := <-c:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
	// Output: timeout
}
```

## C# Equivalent

C# uses **Task.WhenAny** for similar scenarios - waiting for the first of multiple async operations to complete. For true channel multiplexing, C# can use `Channel<T>` with manual coordination.

### **Paradigm Difference:**

**Go:**
- `select` is a **language construct** designed for channel operations
- Blocks on multiple channels simultaneously
- Integral to CSP model
- Random selection when multiple ready

**C#:**
- **Task.WhenAny** works with any async operations
- **Manual channel coordination** required for true multiplexing
- Part of async/await model, not a special construct
- Predictable ordering based on completion order

### When Paradigms Align:
- Waiting for first of multiple operations: `select` ≈ `Task.WhenAny`
- Timeouts: `time.After()` case ≈ `Task.Delay()` with WhenAny
- Multiple data sources: Go select ≈ C# reading multiple channels

### When Paradigms Differ:
- Go: select blocks cheaply, natural for any coordination
- C#: Manual coordination more verbose, often use Task-based patterns instead

### C# Example (Task.WhenAny)

```csharp
using System;
using System.Threading.Tasks;

class Program
{
    static async Task<string> AfterDelay(int seconds, string message)
    {
        await Task.Delay(seconds * 1000);
        return message;
    }

    static async Task Main()
    {
        var task1 = AfterDelay(1, "one");
        var task2 = AfterDelay(2, "two");

        // Wait for both tasks (like Go's example with 2 iterations)
        for (int i = 0; i < 2; i++)
        {
            var completed = await Task.WhenAny(task1, task2);
            var result = await completed;
            Console.WriteLine($"received {result}");
            
            // Remove completed task from consideration
            if (completed == task1)
                task1 = Task.FromResult("already completed");
            else
                task2 = Task.FromResult("already completed");
        }
    }
}
```

### C# Example (Timeout with Task.WhenAny)

```csharp
using System;
using System.Threading.Tasks;

class Program
{
    static async Task<string> LongOperation()
    {
        await Task.Delay(2000);
        return "result";
    }

    static async Task Main()
    {
        var operationTask = LongOperation();
        var timeoutTask = Task.Delay(1000);
        
        var completed = await Task.WhenAny(operationTask, timeoutTask);
        
        if (completed == timeoutTask)
        {
            Console.WriteLine("timeout");
        }
        else
        {
            Console.WriteLine(await operationTask);
        }
        // Output: timeout
    }
}
```

### C# Example (Manual Channel Multiplexing)

For true channel-style multiplexing in C#:

```csharp
using System;
using System.Threading.Channels;
using System.Threading.Tasks;

class Program
{
    static async Task Main()
    {
        var channel1 = Channel.CreateUnbounded<string>();
        var channel2 = Channel.CreateUnbounded<string>();

        // Producer for channel 1
        _ = Task.Run(async () =>
        {
            await Task.Delay(1000);
            await channel1.Writer.WriteAsync("one");
        });

        // Producer for channel 2
        _ = Task.Run(async () =>
        {
            await Task.Delay(2000);
            await channel2.Writer.WriteAsync("two");
        });

        // Multiplex channels (simulate select)
        for (int i = 0; i < 2; i++)
        {
            var task1 = channel1.Reader.ReadAsync().AsTask();
            var task2 = channel2.Reader.ReadAsync().AsTask();
            
            var completed = await Task.WhenAny(task1, task2);
            
            if (completed == task1)
            {
                Console.WriteLine($"received {await task1}");
            }
            else
            {
                Console.WriteLine($"received {await task2}");
            }
        }
    }
}
```

### C# Alternative: Merge Channels

Another approach is to merge channels into one:

```csharp
static async IAsyncEnumerable<T> Merge<T>(
    ChannelReader<T> reader1, 
    ChannelReader<T> reader2)
{
    var task1 = reader1.WaitToReadAsync().AsTask();
    var task2 = reader2.WaitToReadAsync().AsTask();
    
    while (true)
    {
        var completed = await Task.WhenAny(task1, task2);
        
        if (completed == task1)
        {
            if (await task1)
            {
                while (reader1.TryRead(out var item))
                    yield return item;
                task1 = reader1.WaitToReadAsync().AsTask();
            }
            else
            {
                // Channel 1 completed
                reader1 = null;
            }
        }
        else // completed == task2
        {
            if (await task2)
            {
                while (reader2.TryRead(out var item))
                    yield return item;
                task2 = reader2.WaitToReadAsync().AsTask();
            }
            else
            {
                // Channel 2 completed
                reader2 = null;
            }
        }
        
        if (reader1 == null && reader2 == null)
            break;
    }
}
```

## Key Comparison

| Feature | Go (select) | C# (Task.WhenAny) | C# (Manual Channel Multiplex) |
|---------|-------------|-------------------|-------------------------------|
| **Syntax** | Language construct | Library method | Manual implementation |
| **Purpose** | Multiplex channels | Multiplex any async ops | Multiplex channels |
| **Blocking** | Blocks goroutine cheaply | Awaits async | Awaits async |
| **Random selection** | Yes, if multiple ready | No, deterministic | No, deterministic |
| **Idiomatic?** | Yes, primary pattern | Yes, for mixed async ops | No, verbose |
| **Send operations** | Supported in cases | Not directly supported | Manual WriteAsync |

## Common Patterns

### Pattern 1: Timeout
**Go:**
```go
select {
case res := <-ch:
    // Use res
case <-time.After(timeout):
    // Handle timeout
}
```

**C#:**
```csharp
var completed = await Task.WhenAny(operation, Task.Delay(timeout));
if (completed == operation)
    // Use result
else
    // Handle timeout
```

### Pattern 2: Quit Channel
**Go:**
```go
select {
case v := <-dataCh:
    // Process v
case <-quitCh:
    return
}
```

**C#:**
```csharp
// Use CancellationToken
await operation.AsTask(cancellationToken);
```

### Pattern 3: Non-blocking Receive (see next page)
**Go:**
```go
select {
case v := <-ch:
    // Use v
default:
    // Nothing available
}
```

**C#:**
```csharp
if (channel.Reader.TryRead(out var value))
    // Use value
else
    // Nothing available
```

## When to Use Each

**Use Go select** when:
- Coordinating multiple goroutines via channels
- Building concurrent pipelines
- Implementing timeouts or cancellation with channels
- Need random selection for fairness

**Use C# Task.WhenAny** when:
- Waiting for first of multiple async operations
- Implementing timeouts
- Race scenarios
- Mixed async sources (network, timers, etc.)

**Use C# Channel multiplexing** when:
- Porting Go select code to C#
- True producer-consumer channel coordination
- Building Go-style concurrent systems in C#
