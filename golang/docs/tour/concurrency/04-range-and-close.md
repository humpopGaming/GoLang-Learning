# Range and Close

## Go Concept

Channels can be **closed** to signal that no more values will be sent. Receivers can test whether a channel is closed:

```go
v, ok := <-ch
```

`ok` is `false` if the channel is closed and empty.

The `range` loop receives values from a channel repeatedly until it's closed:

```go
for v := range ch {
    // Use v
}
```

**Important rules:**
- **Only the sender** should close a channel, never the receiver
- Closing is optional; channels don't need to be closed unless the receiver needs to know there are no more values (like terminating a `range` loop)
- Sending on a closed channel causes a panic
- Receiving from a closed channel returns the zero value immediately

### Go Example

```go
package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // Signal no more values
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	
	// Range over channel until it's closed
	for i := range c {
		fmt.Println(i)
	}
	// Output: 0 1 1 2 3 5 8 13 21 34
}
```

### Testing for Closed Channel

```go
func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	close(c)
	
	// Receive until closed
	for {
		v, ok := <-c
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println(v)
	}
}
```

## C# Equivalent

C# **Channel<T>** has `Writer.Complete()` to signal completion (similar to closing). C# also supports **IAsyncEnumerable<T>** and **cancellation tokens** for signaling completion of async operations.

Key differences:
- Go: `close(ch)` is a function call
- C#: `Writer.Complete()` is a method on the writer
- Go: `range` loops automatically stop on close
- C#: `await foreach` stops when channel is completed
- Both: Only producers should signal completion

### C# Example (Channel<T>)

```csharp
using System;
using System.Threading.Channels;
using System.Threading.Tasks;

class Program
{
    static async Task Fibonacci(int n, ChannelWriter<int> writer)
    {
        int x = 0, y = 1;
        for (int i = 0; i < n; i++)
        {
            await writer.WriteAsync(x);
            (x, y) = (y, x + y);
        }
        writer.Complete(); // Signal no more values
    }

    static async Task Main()
    {
        var channel = Channel.CreateBounded<int>(10);
        
        var task = Fibonacci(10, channel.Writer);
        
        // Await foreach over channel until it's completed
        await foreach (var value in channel.Reader.ReadAllAsync())
        {
            Console.WriteLine(value);
        }
        // Output: 0 1 1 2 3 5 8 13 21 34
        
        await task;
    }
}
```

### Testing for Completed Channel

```csharp
using System;
using System.Threading.Channels;
using System.Threading.Tasks;

class Program
{
    static async Task Main()
    {
        var channel = Channel.CreateBounded<int>(2);
        
        await channel.Writer.WriteAsync(1);
        await channel.Writer.WriteAsync(2);
        channel.Writer.Complete();
        
        // Read until completed
        while (await channel.Reader.WaitToReadAsync())
        {
            while (channel.Reader.TryRead(out var value))
            {
                Console.WriteLine(value);
            }
        }
        Console.WriteLine("Channel completed");
    }
}
```

### C# Example (Using IAsyncEnumerable)

C# can also use **async enumerables** to represent streams of data:

```csharp
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

class Program
{
    static async IAsyncEnumerable<int> Fibonacci(int n)
    {
        int x = 0, y = 1;
        for (int i = 0; i < n; i++)
        {
            yield return x;
            await Task.Delay(10); // Simulate async work
            (x, y) = (y, x + y);
        }
        // Completion is automatic when method ends
    }

    static async Task Main()
    {
        await foreach (var value in Fibonacci(10))
        {
            Console.WriteLine(value);
        }
        // Output: 0 1 1 2 3 5 8 13 21 34
    }
}
```

## Key Comparison

| Feature | Go (Channels) | C# (Channel<T>) | C# (IAsyncEnumerable) |
|---------|---------------|-----------------|----------------------|
| **Close/Complete** | `close(ch)` | `Writer.Complete()` | Implicit (method end) |
| **Iterate** | `for v := range ch` | `await foreach (var v in reader)` | `await foreach (var v in enumerable)` |
| **Test if closed** | `v, ok := <-ch` | `await WaitToReadAsync()` | N/A (foreach handles it) |
| **Receive after close** | Zero value, `ok=false` | `WaitToReadAsync()` returns false | Enumeration ends |
| **Send after close** | Panic | `ChannelClosedException` | N/A |
| **Who closes?** | Sender only | Sender only | Automatic |

## Cancellation in C#

C# often uses **CancellationToken** alongside channels for cancellation:

```csharp
static async Task Producer(ChannelWriter<int> writer, CancellationToken ct)
{
    try
    {
        for (int i = 0; !ct.IsCancellationRequested; i++)
        {
            await writer.WriteAsync(i, ct);
            await Task.Delay(100, ct);
        }
    }
    finally
    {
        writer.Complete();
    }
}

static async Task Main()
{
    var cts = new CancellationTokenSource();
    var channel = Channel.CreateUnbounded<int>();
    
    var producer = Producer(channel.Writer, cts.Token);
    
    // Cancel after 1 second
    cts.CancelAfter(1000);
    
    await foreach (var value in channel.Reader.ReadAllAsync())
    {
        Console.WriteLine(value);
    }
    
    await producer;
}
```

## When to Close/Complete

**Go:**
- Close when receiver needs to know no more values are coming
- Required for `range` loops to terminate
- Optional for fixed communication patterns

**C#:**
- Complete when producer is done
- Required for `await foreach` to terminate
- Use CancellationToken for external cancellation
- Consider IAsyncEnumerable for simpler producer patterns

## Common Patterns

**Fan-out (one producer, many consumers):**
- Go: One goroutine closes after sending all values
- C#: One task completes writer after sending all values

**Fan-in (many producers, one consumer):**
- Go: Use sync.WaitGroup, close after all producers finish
- C#: Use Task.WhenAll, complete after all producers finish

**Pipeline:**
- Go: Each stage closes its output channel when input is closed
- C#: Each stage completes its writer when input is completed
