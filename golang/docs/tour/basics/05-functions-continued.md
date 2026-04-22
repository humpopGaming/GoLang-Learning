# Functions Continued

## Go Concept

When consecutive function parameters share the same type, you can omit the type from all but the last parameter. This shorthand reduces repetition and improves readability.

For example, `x int, y int` can be shortened to `x, y int`.

### Go Example

```go
package main

import "fmt"

// Long form
func addLong(x int, y int) int {
	return x + y
}

// Short form (preferred when types match)
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

## C# Equivalent

C# **does not have** a direct equivalent for this type shorthand. Each parameter must explicitly declare its type, even when consecutive parameters share the same type.

However, C# does support parameter arrays (`params`) which allow a variable number of same-typed arguments, though this serves a different purpose than Go's shorthand.

### C# Example

```csharp
using System;

class Program
{
    // C# requires explicit type for each parameter
    static int Add(int x, int y)
    {
        return x + y;
    }

    // params keyword allows variable arguments (different concept)
    static int Sum(params int[] numbers)
    {
        int total = 0;
        foreach (int num in numbers)
        {
            total += num;
        }
        return total;
    }

    static void Main()
    {
        Console.WriteLine(Add(42, 13));
        Console.WriteLine(Sum(1, 2, 3, 4, 5));  // Variable arguments
    }
}
```

## Key Differences

- **Type Shorthand**: Go allows `x, y int`; C# requires `int x, int y` for every parameter
- **Verbosity**: C# is more verbose when multiple parameters share a type
- **Clarity**: C# is more explicit; Go favors conciseness
- **Variable Arguments**: C# has `params` for variable-length arguments; Go uses variadic functions with `...Type` (covered later)
- **Use Case Alignment**: While Go's shorthand reduces typing for common cases (especially with multiple same-typed parameters), C#'s explicit approach ensures each parameter's type is always visible. The use cases are the same (passing multiple parameters), but Go optimizes for the common case where types match. In C#, you simply write out each parameter fully.

This is a **syntax convenience in Go with no direct C# equivalent**. C# developers write out each type, which is more typing but arguably clearer at a glance.
