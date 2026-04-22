# Switch with No Condition

## Go Concept

A switch without a condition is the same as `switch true`. This form can be used to write long if-else chains in a cleaner way.

This pattern is idiomatic in Go for replacing complex if-else-if chains with something more readable.

### Go Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

## C# Equivalent

C# **does not have** a direct equivalent syntax for switch with no condition. The idiomatic C# approach is:

1. **If-else chain**: Most straightforward
2. **Switch on true** (C# 8.0+): `true switch { ... }` with switch expressions
3. **Traditional if-else**: The classic approach

C# developers typically use if-else chains for this pattern, while Go developers use conditionless switch.

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        var t = DateTime.Now;
        
        // Idiomatic C#: if-else chain
        if (t.Hour < 12)
        {
            Console.WriteLine("Good morning!");
        }
        else if (t.Hour < 17)
        {
            Console.WriteLine("Good afternoon.");
        }
        else
        {
            Console.WriteLine("Good evening.");
        }

        // C# 8.0+ switch expression on a tuple (creative alternative)
        var greeting = (t.Hour) switch
        {
            < 12 => "Good morning!",
            < 17 => "Good afternoon.",
            _ => "Good evening."
        };
        Console.WriteLine(greeting);
    }
}
```

## Key Differences

- **Conditionless Switch**: Go has `switch {}`; C# doesn't have this syntax
- **Idiom**: Go uses conditionless switch for if-else chains; C# uses actual if-else or switch expressions
- **Readability**: Go's switch looks cleaner for multiple conditions; C#'s if-else is more explicit
- **Pattern Matching**: C# switch expressions with patterns can achieve similar results but with different syntax
- **Use Case Alignment**: Both handle multiple conditional branches. Go's conditionless switch provides a cleaner way to write if-else chains. **C# lacks this specific feature**. 

In C#, you choose between:
1. **If-else chains** — Most explicit and widely understood
2. **Switch expressions with patterns** — Modern C# (8.0+) approach, works well for specific cases

Go's `switch {}` is essentially syntactic sugar for `if-else-if` chains, making them look more uniform with regular switches. C# doesn't have this convenience — if-else chains remain as if-else chains.

**This is a Go-specific feature** for cleaner conditional chains. The use case (multi-way branching based on different conditions) is common to both languages, but Go provides special syntax for it while C# uses traditional if-else.
