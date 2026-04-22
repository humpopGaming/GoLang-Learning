# If and Else

## Go Concept

Variables declared in an `if` short statement are also available in any corresponding `else` blocks.

This extends the scoping of the if-short-statement to include the else branch, making it convenient for variables that need to be checked and used in both branches.

### Go Example

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)  // v is still in scope
	}
	// v is out of scope here
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

C# doesn't have if-short-statements, so the else scoping doesn't exist either. C# developers use:

1. **Variable in outer scope**: Accessible in both if and else
2. **Pattern matching**: For specific type/null checking scenarios

The pattern matching in C# 7.0+ provides similar scoping for matched values, but only for type checks and deconstruction.

### C# Example

```csharp
using System;

class Program
{
    static double Pow(double x, double n, double lim)
    {
        // Declare in outer scope (accessible in both branches)
        double v = Math.Pow(x, n);
        if (v < lim)
        {
            return v;
        }
        else
        {
            Console.WriteLine($"{v} >= {lim}");
        }
        return lim;
    }

    static void Main()
    {
        Console.WriteLine(Pow(3, 2, 10));
        Console.WriteLine(Pow(3, 3, 20));
    }
}

// C# pattern matching example (different use case):
class Example
{
    static void ProcessValue(object obj)
    {
        if (obj is int value && value > 0)
        {
            Console.WriteLine($"Positive: {value}");
        }
        else if (obj is int negValue && negValue <= 0)
        {
            Console.WriteLine($"Non-positive: {negValue}");
        }
        // value and negValue are out of scope here
    }
}
```

## Key Differences

- **Short Statement**: Go has it; C# doesn't
- **Variable Scoping**: Go scopes to if+else blocks; C# typically uses function scope
- **Convenience**: Go's scoping is convenient for if-else chains with temporary values; C# requires variable in outer scope
- **Pattern Matching**: C# pattern matching provides per-branch scoping for matched values, but only for type/null checks
- **Use Case Alignment**: Go's feature is designed for temporary values needed in conditional branches. **C# has no direct equivalent**. The closest C# patterns are:
  1. Declare variable in function scope (it persists beyond if-else)
  2. Use pattern matching for type/value checks (limited use cases)

This is another **Go-specific feature** that reduces variable scope and improves clarity. In C#, variables typically have wider scope, which can be less clean but is what the language provides.

The practical impact is minor — both languages support if-else logic. Go just provides cleaner scoping for variables that are only needed within the conditional blocks.
