# Forever

## Go Concept

If you omit the loop condition entirely, you create an **infinite loop**. This is commonly written as `for {}` or `for { }` in Go.

Infinite loops are useful for servers, event loops, and situations where you break out based on internal logic rather than a loop condition.

### Go Example

```go
package main

func main() {
	// Infinite loop
	for {
		// This loops forever until break or return
		// Typically used in servers or event loops
	}
}
```

## C# Equivalent

C# has several ways to create infinite loops:

1. **`while (true)`** — Most explicit and commonly used
2. **`for (;;)`** — C-style infinite loop
3. **`do { } while (true);`** — Less common

C# developers typically use `while (true)` because it's most explicit and readable.

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        // Most idiomatic C# infinite loop
        while (true)
        {
            // Loops forever until break or return
        }

        // Alternative: for style (less common)
        for (;;)
        {
            // Also infinite
        }
    }
}
```

## Key Differences

- **Syntax**: Go uses `for {}`; C# typically uses `while (true)`
- **Clarity**: C#'s `while (true)` is more explicit; Go's `for {}` is more concise
- **Alternatives**: C# has multiple ways; Go has one way
- **Common Usage**: In Go, `for {}` is standard; in C#, `while (true)` is standard
- **Use Case Alignment**: Both create infinite loops for the same purposes (servers, event processors, etc.). The only difference is syntax and explicitness. **C#'s `while (true)` immediately signals an intentional infinite loop**, while Go's `for {}` might look like a typo to newcomers. Both require `break` or `return` to exit.

Both serve identical use cases. C# is more explicit, Go is more concise. Personal/team preference and language idioms determine which is "better."
