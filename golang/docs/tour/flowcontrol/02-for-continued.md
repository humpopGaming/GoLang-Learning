# For Continued

## Go Concept

The init and post statements in a `for` loop are **optional**. You can omit them while keeping the semicolons, or omit them entirely for a while-style loop.

This flexibility means you can use `for` for different iteration patterns without needing separate loop types.

### Go Example

```go
package main

import "fmt"

func main() {
	sum := 1
	// Init and post omitted (semicolons dropped)
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

## C# Equivalent

C# `for` loop also allows omitting init and post statements, but you typically use a `while` loop instead when you don't need them:

- `for (;condition;)` is valid but unusual
- `while (condition)` is the idiomatic choice

C# has explicit loop types for different patterns, while Go uses `for` for everything.

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        int sum = 1;
        
        // Idiomatic C#: use while loop
        while (sum < 1000)
        {
            sum += sum;
        }
        Console.WriteLine(sum);

        // Valid but unusual: for loop with only condition
        int sum2 = 1;
        for (; sum2 < 1000;)
        {
            sum2 += sum2;
        }
    }
}
```

## Key Differences

- **Loop Types**: Go uses `for` for all loops; C# has dedicated `while` loop for this pattern
- **Idiomacy**: Go's `for condition {}` is standard; C#'s `for (; condition;)` is unusual
- **Readability**: C#'s `while` is clearer for condition-only loops; Go's `for` is consistent but less descriptive
- **Flexibility vs Explicitness**: Go's single loop type is simple; C#'s multiple types are more explicit about intent
- **Use Case Alignment**: Both support condition-only looping (like while loops). C# has a dedicated `while` keyword that makes this intent clearer, while Go achieves the same functionality by allowing you to omit parts of the `for` loop. The use case is identical — checking a condition before each iteration without init or post logic. C# developers use `while`; Go developers use `for` without init/post.
