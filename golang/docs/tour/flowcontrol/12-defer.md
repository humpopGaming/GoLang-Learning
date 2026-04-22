# Defer

## Go Concept

A `defer` statement defers the execution of a function until the surrounding function returns. The deferred function's arguments are evaluated immediately, but the function call itself doesn't happen until the surrounding function returns.

This is commonly used for cleanup tasks (closing files, unlocking mutexes, etc.) that must happen regardless of how the function exits.

### Go Example

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
// Output:
// hello
// world
```

## C# Equivalent

C# **does not have** a direct equivalent to Go's `defer`. The closest alternatives are:

1. **`using` statement**: For `IDisposable` objects (automatic cleanup)
2. **`try-finally` block**: Ensures cleanup code runs
3. **Manual cleanup**: Explicitly call cleanup before returns

None of these provide the same syntactic convenience as Go's `defer`. The `using` statement is closest in spirit but only works for `IDisposable` types.

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        // No direct defer equivalent
        // Option 1: Manual order (not the same as defer)
        Console.WriteLine("hello");
        Console.WriteLine("world");

        // Option 2: try-finally for guaranteed cleanup
        try
        {
            Console.WriteLine("hello");
        }
        finally
        {
            Console.WriteLine("world");  // Always runs
        }

        // Option 3: using statement for IDisposable (most similar concept)
        using (var resource = new MyResource())
        {
            Console.WriteLine("Using resource");
        }  // Dispose() called automatically here
    }
}

class MyResource : IDisposable
{
    public void Dispose()
    {
        Console.WriteLine("Resource disposed");
    }
}
```

## Key Differences

- **Defer Keyword**: Go has `defer`; C# doesn't
- **Syntax Convenience**: Go can defer any function with one word; C# requires try-finally or using
- **Use Cases**: Go's defer works for any cleanup; C#'s using only works for IDisposable
- **Evaluation Timing**: Go evaluates defer arguments immediately but calls later; C# using/finally blocks evaluate and execute at different times
- **Multiple Defers**: Go can have multiple defers (stack-based); C# would need nested try-finally or multiple using statements
- **Use Case Alignment**: Both ensure cleanup code runs, but **Go's defer is more flexible and convenient**. 

**C# doesn't have a direct equivalent**. The use cases overlap:
- **Go's defer**: Used for any cleanup (close files, unlock mutexes, logging, etc.)
- **C#'s using**: Only for IDisposable objects (files, database connections, etc.)
- **C#'s try-finally**: For any cleanup, but more verbose than Go's defer

**When to use each**:
- Go: `defer file.Close()` — Simple, one line
- C#: `using (var file = ...)` or try-finally block — More boilerplate

Go's defer is a language feature that makes resource cleanup more elegant. C# achieves similar goals with different patterns (using for disposables, try-finally for general cleanup), but these are more verbose and less flexible.
