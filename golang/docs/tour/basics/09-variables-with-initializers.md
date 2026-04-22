# Variables with Initializers

## Go Concept

A `var` declaration can include **initializers**, one per variable. When an initializer is present, the type can be omitted — the variable will take the type of the initializer (type inference).

This allows for more concise variable declarations when the type is obvious from the value.

### Go Example

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
	// Type inferred from initializer
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

## C# Equivalent

C# also supports variable initialization at declaration. C# offers multiple options:

1. **Explicit type with initializer**: `int i = 1;`
2. **Type inference with `var`**: `var i = 1;` (type inferred from right side)
3. **Multiple initialization**: Must initialize separately or use separate statements

C#'s `var` keyword is used for type inference, which can be confusing since Go's `var` is used for explicit type declaration. The concepts align but the keywords differ.

### C# Example

```csharp
using System;

class Program
{
    // Class-level with initializers
    static int i = 1, j = 2;

    static void Main()
    {
        // Explicit types
        bool c = true;
        bool python = false;
        string java = "no!";

        // Type inference with var
        var c2 = true;          // bool inferred
        var python2 = false;    // bool inferred
        var java2 = "no!";      // string inferred

        Console.WriteLine($"{i} {j} {c} {python} {java}");
    }
}
```

## Key Differences

- **Type Inference Keyword**: Go uses `var name = value` for inference; C# uses `var name = value` for inference
- **Confusing Keyword Usage**: Go's `var` means "variable declaration"; C#'s `var` specifically means "infer the type"
- **Multiple Variables**: Go can declare and initialize multiple variables of different types in one line; C# typically does them separately
- **Scope Requirements**: Go's `var` with inference works at package and function level; C#'s `var` only works at local scope (not class-level fields)
- **Explicit Type with Initializer**: Both support `var name type = value`, but it's redundant in both languages when you have an initializer
- **Use Case Alignment**: Both languages support type inference from initializers, reducing boilerplate. However, the keyword usage is opposite: Go's `var` is general-purpose (with or without inference), while C#'s `var` specifically signals type inference. In practice, Go developers use `:=` for local variables (covered next), making `var` less common. C# developers use `var` extensively for local variables to reduce verbosity.

The concepts align well, but the keywords can be confusing when switching between languages.
