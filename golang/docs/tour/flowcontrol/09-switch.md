# Switch

## Go Concept

Go's `switch` statement is a shorter way to write a sequence of if-else statements. It evaluates cases from top to bottom, stopping when a case succeeds.

Key features:
- No parentheses around the switch expression
- **No fall-through by default** (unlike C/C#) — Go runs only the selected case
- Cases can be expressions (not just constants)
- No `break` needed (it's implicit)

### Go Example

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
```

## C# Equivalent

C# `switch` statements are similar but have different default behavior:

- Requires parentheses around the switch expression
- **Fall-through requires explicit `goto case`** (default is no fall-through like Go)
- Traditionally required constant case values, but C# 7.0+ added pattern matching
- Requires `break` at the end of each case (or another jump statement)

Modern C# (8.0+) also has **switch expressions** which are more concise and expression-based.

### C# Example

```csharp
using System;
using System.Runtime.InteropServices;

class Program
{
    static void Main()
    {
        Console.Write("C# runs on ");
        
        // Determine OS
        string os;
        if (RuntimeInformation.IsOSPlatform(OSPlatform.OSX))
            os = "darwin";
        else if (RuntimeInformation.IsOSPlatform(OSPlatform.Linux))
            os = "linux";
        else if (RuntimeInformation.IsOSPlatform(OSPlatform.Windows))
            os = "windows";
        else
            os = "unknown";

        switch (os)
        {
            case "darwin":
                Console.WriteLine("OS X.");
                break;  // Explicit break required
            case "linux":
                Console.WriteLine("Linux.");
                break;
            default:
                Console.WriteLine($"{os}.");
                break;
        }

        // C# 8.0+ switch expression (more concise):
        var message = os switch
        {
            "darwin" => "OS X.",
            "linux" => "Linux.",
            _ => $"{os}."
        };
        Console.WriteLine(message);
    }
}
```

## Key Differences

- **Parentheses**: Go doesn't use them; C# requires them
- **Break Statement**: Go doesn't need `break` (no fall-through); C# requires `break`
- **Fall-Through**: Go's default is no fall-through (use `fallthrough` keyword for it); C# also defaults to no fall-through but requires `break`
- **Short Statement**: Go supports switch short statement like `switch os := getOS(); os {}`; C# doesn't
- **Expression Cases**: Go always supported expression cases; C# added pattern matching in 7.0+
- **Switch Expressions**: C# 8.0+ has concise switch expressions; Go doesn't have this syntax
- **Use Case Alignment**: Both provide multi-way branching. Go's syntax is cleaner (no parentheses, no breaks). C#'s traditional switch requires more boilerplate but modern C# switch expressions are very concise. The use cases align perfectly — both handle multi-way branches based on a value.

**Modern C# (8.0+) switch expressions** are actually more concise than Go for simple value mappings, while Go's traditional switch is cleaner than C#'s traditional switch for statement-based cases.
