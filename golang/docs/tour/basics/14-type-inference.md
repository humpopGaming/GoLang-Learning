# Type Inference

## Go Concept

When declaring a variable without specifying a type (using `:=` or `var =` without type), the variable's type is **inferred** from the value on the right side.

- For untyped numeric constants, the type depends on precision of the constant
- For typed expressions, the type is the same as the expression's type
- `42` → `int`, `3.142` → `float64`, `0.867 + 0.5i` → `complex128`

### Go Example

```go
package main

import "fmt"

func main() {
	v := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}
```

## C# Equivalent

C# has type inference using the `var` keyword. The compiler infers the type from the right-hand side of the assignment:

- `var x = 42` → `int`
- `var x = 3.142` → `double` (not float!)
- C# has no complex number literals (must use `new Complex()`)

C#'s inference is very similar to Go's, but C# has different defaults for numeric literals (integers default to `int`, floating-point defaults to `double`).

### C# Example

```csharp
using System;
using System.Numerics;

class Program
{
    static void Main()
    {
        var v = 42;              // int
        var f = 3.142;           // double (not float!)
        var g = new Complex(0.867, 0.5);  // Complex
        
        Console.WriteLine($"v is of type {v.GetType()}");
        Console.WriteLine($"f is of type {f.GetType()}");
        Console.WriteLine($"g is of type {g.GetType()}");

        // Explicit float requires suffix
        var f2 = 3.142f;         // float
        var d = 3.142m;          // decimal
    }
}
```

## Key Differences

- **Keyword**: Go uses `:=` or `var =`; C# uses `var`
- **Float Defaults**: Go infers `float64` for decimal literals; C# infers `double` (same thing, different names)
- **Complex Numbers**: Go has complex literals; C# requires explicit construction
- **Literal Suffixes**: C# uses suffixes (`f` for float, `m` for decimal, `L` for long); Go uses type conversions
- **Integer Sizes**: Go's untyped integers can represent any integer; C# integers are always sized types
- **String Interpolation**: C# has built-in string interpolation (`$"{}"`); Go uses `fmt.Printf` or similar
- **Scope**: Go's `:=` only works in functions; C#'s `var` only works in methods (both are local-only)
- **Use Case Alignment**: Both languages reduce boilerplate by inferring types from values. The mechanics are nearly identical - compiler looks at the right side and determines the type. The main practical difference is that C# floating-point literals default to `double` and require suffixes for `float`, while Go defaults to `float64` (which is equivalent to `double`) without suffixes.

Both approaches work well and serve the same purpose: let the compiler figure out obvious types to reduce verbosity without sacrificing type safety.
