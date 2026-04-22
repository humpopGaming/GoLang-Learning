# For is Go's "While"

## Go Concept

Go doesn't have a `while` keyword. Instead, you drop the semicolons from the `for` loop to create a while-style loop. The `for` loop with just a condition is Go's "while loop".

This simplifies the language by having one looping construct that handles all cases.

### Go Example

```go
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {  // This is Go's "while"
		sum += sum
	}
	fmt.Println(sum)
}
```

## C# Equivalent

C# has a dedicated **`while` loop** for this pattern. The `while` loop has only a condition and no init or post statements.

C#'s approach is more explicit — different keywords for different loop types. Go's approach is simpler — one keyword for all loops.

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        int sum = 1;
        while (sum < 1000)  // C# has explicit while keyword
        {
            sum += sum;
        }
        Console.WriteLine(sum);
    }
}
```

## Key Differences

- **While Keyword**: C# has `while`; Go uses `for` without semicolons
- **Explicitness**: C#'s separate keyword makes intent clearer; Go's unified approach is simpler
- **Language Complexity**: Go has fewer keywords; C# has more specialized constructs
- **Readability**: `while` immediately signals "condition-only loop"; `for condition` requires knowing Go's convention
- **Use Case Alignment**: Both provide condition-only loops that iterate while a condition is true. The functionality is identical, only the syntax differs. C# makes the distinction between different loop types explicit with different keywords (`for`, `while`, `do-while`), while Go uses `for` for everything, relying on what you include or omit to determine the loop behavior.

This is a language design philosophy difference: C# favors explicit, specialized constructs; Go favors simplicity and fewer keywords. Neither is wrong — it's a trade-off between expressiveness and simplicity.
