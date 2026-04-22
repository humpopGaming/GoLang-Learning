# Exercise: Loops and Functions

## Go Concept

This exercise demonstrates using loops and functions together to implement Newton's method for finding square roots. It reinforces iteration patterns and function calls.

The key learning is using a `for` loop to iteratively refine a calculation until it reaches desired accuracy.

### Go Example

```go
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("Compare to:", math.Sqrt(2))
}
```

## C# Equivalent

C# can implement Newton's method identically using similar loop constructs. The algorithm and approach are the same.

This demonstrates that algorithmic approaches translate directly between the languages — the main differences are syntax.

### C# Example

```csharp
using System;

class Program
{
    static double Sqrt(double x)
    {
        double z = 1.0;
        for (int i = 0; i < 10; i++)
        {
            z -= (z * z - x) / (2 * z);
        }
        return z;
    }

    static void Main()
    {
        Console.WriteLine(Sqrt(2));
        Console.WriteLine($"Compare to: {Math.Sqrt(2)}");
    }
}
```

## Key Differences

- **Algorithm**: Identical in both languages
- **Loop Syntax**: Go without parentheses; C# with parentheses
- **Math Functions**: `math.Sqrt` vs `Math.Sqrt` (same functionality)
- **Use Case Alignment**: This exercise shows that **algorithmic patterns translate directly** between Go and C#. Iterative refinement, numerical methods, and similar computational patterns work the same way. The syntax differs slightly, but the logic and structure are identical.

This is less about language-specific features and more about demonstrating that **computational algorithms are language-independent** — they work the same way in Go and C# with only syntactic differences.
