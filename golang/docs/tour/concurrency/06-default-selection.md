# Default Selection

## Go Concept

The **`default`** case in a `select` statement runs if no other case is ready. This makes `select` **non-blocking**:

```go
select {
case v := <-ch:
    // Use v
default:
    // No value ready, do something else
}
```

Use `default` to:
- Implement non-blocking sends and receives
- Poll channels without blocking
- Implement try-send and try-receive patterns
- Avoid deadlocks when channels might not be ready

### Go Example: Non-blocking Receive

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 42

	// Non-blocking receive
	select {
	case v := <-ch:
		fmt.Println("received", v)
	default:
		fmt.Println("no value ready")
	}
	// Output: received 42

	// Channel now empty
	select {
	case v := <-ch:
		fmt.Println("received", v)
	default:
		fmt.Println("no value ready")
	}
	// Output: no value ready
}
```

### Go Example: Non-blocking Send

```go
func main() {
	ch := make(chan int, 1)
	ch <- 1 // Buffer full

	// Try to send without blocking
	select {
	case ch <- 2:
		fmt.Println("sent 2")
	default:
		fmt.Println("channel full, cannot send")
	}
	// Output: channel full, cannot send
}
```

### Go Example: Polling Pattern

```go
func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```

## C# Equivalent

C# **Channel<T>** provides **TryRead** and **TryWrite** methods for non-blocking operations. These are the direct equivalents to Go's `select` with `default`.

Key differences:
- Go: `select` with `default` is a language construct
- C#: TryRead/TryWrite are methods on the channel
- Both: Enable non-blocking operations
- C# also has async versions: `WaitToReadAsync(timeout)`, `WaitToWriteAsync(timeout)`

### C# Example: Non-blocking Receive

```csharp
using System;
using System.Threading.Channels;

class Program
{
    static void Main()
    {
        var channel = Channel.CreateBounded<int>(1);
        channel.Writer.TryWrite(42);

        // Non-blocking receive
        if (channel.Reader.TryRead(out var value))
        {
            Console.WriteLine($"received {value}");
        }
        else
        {
            Console.WriteLine("no value ready");
        }
        // Output: received 42

        // Channel now empty
        if (channel.Reader.TryRead(out value))
        {
            Console.WriteLine($"received {value}");
        }
        else
        {
            Console.WriteLine("no value ready");
        }
        // Output: no value ready
    }
}
```

### C# Example: Non-blocking Send

```csharp
using System;
using System.Threading.Channels;

class Program
{
    static void Main()
    {
        var channel = Channel.CreateBounded<int>(1);
        channel.Writer.TryWrite(1); // Buffer full

        // Try to send without blocking
        if (channel.Writer.TryWrite(2))
        {
            Console.WriteLine("sent 2");
        }
        else
        {
            Console.WriteLine("channel full, cannot send");
        }
        // Output: channel full, cannot send
    }
}
```

### C# Example: Polling Pattern with Timeout

```csharp
using System;
using System.Threading;
using System.Threading.Channels;
using System.Threading.Tasks;

class Program
{
    static async Task Main()
    {
        var tickChannel = Channel.CreateUnbounded<bool>();
        var boomChannel = Channel.CreateUnbounded<bool>();
        
        // Tick producer
        _ = Task.Run(async () =>
        {
            while (true)
            {
                await Task.Delay(100);
                tickChannel.Writer.TryWrite(true);
            }
        });
        
        // Boom producer
        _ = Task.Run(async () =>
        {
            await Task.Delay(500);
            await boomChannel.Writer.WriteAsync(true);
        });

        while (true)
        {
            if (tickChannel.Reader.TryRead(out _))
            {
                Console.WriteLine("tick.");
            }
            else if (boomChannel.Reader.TryRead(out _))
            {
                Console.WriteLine("BOOM!");
                return;
            }
            else
            {
                Console.WriteLine("    .");
                await Task.Delay(50);
            }
        }
    }
}
```

### C# Example: Async Timeout Pattern

C# often uses **timeouts** with async operations instead of polling:

```csharp
using System;
using System.Threading;
using System.Threading.Channels;
using System.Threading.Tasks;

class Program
{
    static async Task Main()
    {
        var channel = Channel.CreateUnbounded<int>();
        
        // Try to read with timeout
        using var cts = new CancellationTokenSource(TimeSpan.FromSeconds(1));
        
        try
        {
            if (await channel.Reader.WaitToReadAsync(cts.Token))
            {
                var value = await channel.Reader.ReadAsync();
                Console.WriteLine($"received {value}");
            }
        }
        catch (OperationCanceledException)
        {
            Console.WriteLine("timeout - no value received");
        }
        // Output: timeout - no value received
    }
}
```

## Key Comparison

| Feature | Go (select with default) | C# (TryRead/TryWrite) |
|---------|--------------------------|------------------------|
| **Non-blocking receive** | `select { case v := <-ch: ...; default: ... }` | `if (TryRead(out var v)) { ... } else { ... }` |
| **Non-blocking send** | `select { case ch <- v: ...; default: ... }` | `if (TryWrite(v)) { ... } else { ... }` |
| **Syntax** | Language construct | Method calls |
| **Return value** | Case selected | Boolean success |
| **Multi-channel** | Can check multiple channels | Need separate calls |
| **Timeout** | Combine with `time.After()` | Use async with CancellationToken |

## Comparison to Traditional C# Patterns

| Pattern | Go | C# |
|---------|----|----|
| **Try receive** | `select { case v := <-ch: ...; default: ... }` | `if (TryRead(out var v)) { ... }` |
| **Try send** | `select { case ch <- v: ...; default: ... }` | `if (TryWrite(v)) { ... }` |
| **Poll multiple** | `select` with multiple cases + default | Multiple TryRead calls |
| **Async timeout** | `select { case <-ch: ...; case <-time.After(t): ... }` | `await ReadAsync(cancellationToken)` |

## BlockingCollection Comparison

C#'s older **BlockingCollection<T>** also has try methods:

```csharp
using System.Collections.Concurrent;

var collection = new BlockingCollection<int>(1);
collection.TryAdd(1); // true
collection.TryAdd(2); // false - full

if (collection.TryTake(out var item))
    Console.WriteLine($"got {item}");
else
    Console.WriteLine("empty");
```

## Common Use Cases

### Use Case 1: Try-Send Pattern
**Go:**
```go
select {
case ch <- value:
    fmt.Println("sent")
default:
    fmt.Println("channel full, dropping value")
}
```

**C#:**
```csharp
if (channel.Writer.TryWrite(value))
    Console.WriteLine("sent");
else
    Console.WriteLine("channel full, dropping value");
```

### Use Case 2: Timeout on Receive
**Go:**
```go
select {
case v := <-ch:
    fmt.Println("got", v)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
}
```

**C#:**
```csharp
using var cts = new CancellationTokenSource(TimeSpan.FromSeconds(1));
try {
    var value = await channel.Reader.ReadAsync(cts.Token);
    Console.WriteLine($"got {value}");
} catch (OperationCanceledException) {
    Console.WriteLine("timeout");
}
```

### Use Case 3: Check Without Blocking
**Go:**
```go
select {
case v := <-ch:
    // Process v immediately
default:
    // Continue doing other work
}
```

**C#:**
```csharp
if (channel.Reader.TryRead(out var value))
{
    // Process value immediately
}
else
{
    // Continue doing other work
}
```

## Performance Considerations

**Go:**
- `select` with `default` is very fast
- No allocation
- Safe to use in hot loops
- Cheap even with many cases

**C#:**
- TryRead/TryWrite are also fast
- Minimal allocation
- Safe in hot loops
- Async versions allocate tasks

## When to Use

**Use Go select with default** when:
- You want to check a channel without blocking
- Implementing rate limiting or backpressure
- Building responsive systems that can't afford to block
- Coordinating between multiple channels

**Use C# TryRead/TryWrite** when:
- You want to check a channel without awaiting
- Implementing non-blocking patterns
- Performance-critical synchronous code
- Polling scenarios

**Use C# async with timeout** when:
- You can afford to wait a bit but not forever
- Building responsive UI or network code
- Need cancellation support
- Want to avoid busy-waiting
