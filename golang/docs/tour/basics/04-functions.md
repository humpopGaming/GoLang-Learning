# Functions

## Go Concept

A function in Go is declared with the `func` keyword, followed by the function name, parameters in parentheses, return type, and the function body in braces.

The type of each parameter comes **after** the parameter name (opposite of C-style languages). This "name-first" syntax becomes more readable with complex types and multiple return values.

### Go Example

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

## C# Equivalent

C# function (method) syntax is similar but follows C-style conventions with the type coming **before** the parameter name:

- Both use braces `{}` for function bodies
- Both require explicit return types (though C# has `var` for local variables)
- Both support multiple parameters
- C# methods typically belong to classes; Go functions can be package-level

The key syntactic difference is the parameter order: C# uses `type name`, Go uses `name type`.

### C# Example

```csharp
using System;

class Program
{
    // Method belongs to a class in C#
    static int Add(int x, int y)
    {
        return x + y;
    }

    static void Main()
    {
        Console.WriteLine(Add(42, 13));
    }
}

// C# 9.0+ with top-level statements and local functions:
// int Add(int x, int y) => x + y;
// Console.WriteLine(Add(42, 13));
```

## Key Differences

- **Parameter Syntax**: Go uses `name type`; C# uses `type name`
- **Location**: Go functions can be package-level; C# methods must be in classes (except local functions)
- **Access Modifiers**: C# requires `static`, `public`, etc.; Go uses capitalization for export
- **Expression Bodies**: C# supports `=>` expression-bodied members; Go always uses full syntax
- **Top-Level Functions**: Go natively supports package-level functions; C# historically required classes but now supports top-level statements
- **Use Case Alignment**: Both define reusable blocks of code with parameters and return values. The syntax differs, but the concept is identical. C#'s requirement for classes is more restrictive, though modern C# has relaxed this with local functions and top-level statements.
