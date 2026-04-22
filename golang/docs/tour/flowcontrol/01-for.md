# For

## Go Concept

Go has only one looping construct: the **`for` loop**. The basic `for` loop has three components separated by semicolons:

- **Init statement**: Executed before the first iteration (often a variable declaration)
- **Condition expression**: Evaluated before every iteration (loop continues while true)
- **Post statement**: Executed at the end of every iteration (often an increment)

Unlike C, there are no parentheses around the three components, but the braces `{}` are always required.

### Go Example

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

## C# Equivalent

C# has the **`for` loop** with nearly identical syntax and semantics:

- Same three components: init, condition, post
- C# requires parentheses around the components
- Braces are required for multi-statement blocks (optional for single statements, though not recommended)

The concepts are almost identical, with only minor syntactic differences.

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        int sum = 0;
        for (int i = 0; i < 10; i++)
        {
            sum += i;
        }
        Console.WriteLine(sum);
    }
}
```

## Key Differences

- **Parentheses**: Go doesn't use parentheses around the `for` components; C# requires them
- **Braces**: Go always requires braces; C# allows omitting them for single statements (not recommended)
- **Variable Scope**: Go's `i := 0` scopes `i` to the for loop; C# does the same with `for (int i = 0...)`
- **Only Loop Type**: Go only has `for` (no while or do-while); C# has `for`, `while`, and `do-while`
- **Use Case Alignment**: Both provide iteration with init, condition, and post-iteration logic. The use cases are identical. The only differences are syntactic (parentheses and brace requirements). C# developers will find Go's for loop immediately familiar, just without the parentheses.
