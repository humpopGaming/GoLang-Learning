# Zero Values

## Go Concept

Variables declared without an explicit initial value are given their **zero value**. The zero value is:

- `0` for numeric types
- `false` for boolean type
- `""` (empty string) for strings
- `nil` for pointers, slices, maps, channels, functions, and interfaces

This ensures variables are always in a defined state and never contain garbage values.

### Go Example

```go
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	// Output: 0 0 false ""
}
```

## C# Equivalent

C# has **default values** which serve the same purpose as Go's zero values:

- `0` for numeric types (int, float, double, etc.)
- `false` for bool
- `'\0'` (null character) for char
- `null` for reference types (strings, classes, arrays)
- Default for structs (all fields set to their defaults)

C# distinguishes between **value types** (structs, primitives) and **reference types** (classes, strings, arrays). Reference types default to `null`, similar to Go's `nil`.

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        int i;
        double f;
        bool b;
        string s;
        
        // C# requires initialization before use in local variables
        // So we'll use default keyword or class fields
        i = default;
        f = default;
        b = default;
        s = default;
        
        Console.WriteLine($"{i} {f} {b} \"{s}\"");
        // Output: 0 0 False (empty string shows as nothing)
    }
}

// Or using class fields (which are auto-initialized):
class Defaults
{
    static int i;
    static double f;
    static bool b;
    static string s;
    
    static void ShowDefaults()
    {
        Console.WriteLine($"{i} {f} {b} \"{s}\"");
    }
}
```

## Key Differences

- **Automatic vs Explicit**: Go always initializes to zero values automatically; C# requires initialization for local variables before use (but class fields are auto-initialized)
- **Null vs Nil**: C# uses `null` for reference types; Go uses `nil` for pointers, slices, maps, channels, functions, and interfaces
- **Default Keyword**: C# has `default(T)` or `default` keyword to explicitly get default values; Go doesn't need this
- **String Zero Value**: Go's string zero value is `""` (empty string); C#'s string default is `null` (not the same as empty string)
- **Compiler Checks**: C# compiler prevents using uninitialized local variables; Go allows it (they'll have zero values)
- **Value vs Reference Types**: C# makes a strong distinction; Go has a simpler model
- **Use Case Alignment**: Both languages ensure variables have predictable initial states. Go's approach is simpler and more consistent (everything has a zero value). C#'s distinction between value types and reference types with `null` provides more control but is more complex. **Important difference**: C# strings default to `null`, which can cause `NullReferenceException`. Go strings default to `""`, which is always safe to use. This is a common source of bugs when moving from Go to C#.
