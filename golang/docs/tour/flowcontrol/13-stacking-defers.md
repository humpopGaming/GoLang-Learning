# Stacking Defers

## Go Concept

Deferred function calls are pushed onto a **stack**. When the surrounding function returns, deferred calls are executed in **last-in-first-out (LIFO) order**.

This stacking behavior is particularly useful for properly nesting cleanup operations (like closing nested resources in the reverse order they were opened).

### Go Example

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
// Output:
// counting
// done
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1
// 0
```

## C# Equivalent

Since C# doesn't have `defer`, it also doesn't have defer stacking. However, the **LIFO cleanup pattern** can be achieved with:

1. **Nested `using` statements**: Automatic disposal in reverse order
2. **Nested `try-finally` blocks**: Manual cleanup in reverse order
3. **Stack-based manual tracking**: Explicitly managing a stack of cleanup actions

The `using` statement naturally provides LIFO cleanup for disposable objects because of nesting scope.

### C# Example

```csharp
using System;
using System.Collections.Generic;

class Program
{
    static void Main()
    {
        Console.WriteLine("counting");

        // Manual defer-like behavior with stack
        var deferredActions = new Stack<Action>();
        
        for (int i = 0; i < 10; i++)
        {
            int captured = i;  // Capture for closure
            deferredActions.Push(() => Console.WriteLine(captured));
        }

        Console.WriteLine("done");

        // Execute deferred actions in LIFO order
        while (deferredActions.Count > 0)
        {
            deferredActions.Pop()();
        }

        // Nested using example (LIFO disposal):
        using (var resource1 = new Resource("First"))
        using (var resource2 = new Resource("Second"))
        using (var resource3 = new Resource("Third"))
        {
            Console.WriteLine("Using resources");
        }
        // Disposal happens in reverse: Third, Second, First
    }
}

class Resource : IDisposable
{
    private string name;
    public Resource(string name) { this.name = name; }
    public void Dispose() => Console.WriteLine($"{name} disposed");
}
```

## Key Differences

- **Built-in Stack**: Go's defer has automatic LIFO stacking; C# requires manual implementation or nested using
- **Convenience**: Go's defer is simple and automatic; C# requires design patterns
- **Nested Resources**: C# nested using statements provide LIFO disposal naturally; similar to Go's defer stacking
- **Order Guarantee**: Both guarantee reverse order cleanup (LIFO)
- **Use Case Alignment**: The LIFO cleanup pattern is important in both languages for resource management (open database, start transaction, acquire lock → release lock, commit transaction, close database). 

**Go provides this pattern built-in with defer stacking**. 

**C# achieves it through**:
1. **Nested using statements** — For IDisposable objects, automatically LIFO
2. **Manual stack of actions** — For general cleanup, requires explicit code
3. **Nested try-finally** — Verbose but works for any cleanup

Example comparison:
```go
// Go: Simple and clear
defer close1()
defer close2()
defer close3()
// Executes: close3, close2, close1
```

```csharp
// C#: Nested using (for disposables)
using (var r1 = ...)
using (var r2 = ...)
using (var r3 = ...)
{
}
// Disposes: r3, r2, r1

// C#: Manual stack (for non-disposables)
var stack = new Stack<Action>();
stack.Push(close1);
stack.Push(close2);
stack.Push(close3);
while (stack.Count > 0) stack.Pop()();
```

**Go's defer stacking is more elegant**, but C# can achieve the same LIFO cleanup behavior with more verbose code. The nested `using` pattern in C# is actually quite clean for disposable resources.
