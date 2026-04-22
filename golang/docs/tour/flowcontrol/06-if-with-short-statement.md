# If with Short Statement

## Go Concept

The `if` statement can start with a **short statement** to execute before the condition. Variables declared in this short statement are only in scope within the `if` statement (including `else` blocks).

This pattern is useful for initializing a value that's only needed for the condition check or within the if block.

### Go Example

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// v is only in scope within this if statement
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

## C# Equivalent

C# **does not have** a direct equivalent for if-with-short-statement. The closest alternatives are:

1. **Declare variable just before if**: Variable stays in outer scope
2. **Use a separate scope block**: Adds extra braces
3. **Use pattern matching** (C# 7.0+): For specific cases like type checking or null checking

None of these provide the exact same syntax or scoping behavior as Go's if-with-short-statement.

### C# Example

```csharp
using System;

class Program
{
    static double Pow(double x, double n, double lim)
    {
        // Option 1: Declare before (variable stays in function scope)
        double v = Math.Pow(x, n);
        if (v < lim)
        {
            return v;
        }
        return lim;

        // Option 2: Extra scope block
        {
            double v2 = Math.Pow(x, n);
            if (v2 < lim)
            {
                return v2;
            }
        }
        return lim;
    }

    static void Main()
    {
        Console.WriteLine($"{Pow(3, 2, 10)} {Pow(3, 3, 20)}");
    }
}

// C# 7.0+ pattern matching (different use case):
class Example
{
    static void ProcessObject(object obj)
    {
        if (obj is string str && str.Length > 0)
        {
            // str is in scope here
            Console.WriteLine(str);
        }
        // str is NOT in scope here
    }
}
```

## Key Differences

- **Short Statement Syntax**: Go has `if statement; condition {}`; C# has no equivalent
- **Variable Scoping**: Go's variables are scoped to the if block; C# variables stay in function scope unless manually scoped
- **Conciseness**: Go's syntax is more concise for this pattern; C# requires separate declaration
- **Pattern Matching**: C# has pattern matching with `is` and deconstruction, but it's for different use cases
- **Use Case Alignment**: This is **unique to Go** with no direct C# equivalent. The use case is initializing a variable that's only needed for an if condition and its block. In C#, you typically:
  1. Declare the variable before the if (but it remains in outer scope)
  2. Use a separate `{}` block to limit scope (adds clutter)
  3. For specific cases (type checks, null checks), use pattern matching

Go's approach is cleaner for the specific pattern of "compute value, check condition, use value if condition true." C# developers simply declare variables in the broader scope, which is fine but less elegant for this particular pattern.
