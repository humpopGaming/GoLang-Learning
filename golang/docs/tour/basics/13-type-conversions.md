# Type Conversions

## Go Concept

Go requires **explicit type conversions**. The expression `T(v)` converts the value `v` to type `T`. Unlike C, there are no implicit type conversions in Go — you must explicitly convert between types, even between numeric types.

This explicitness prevents bugs from unexpected type coercions and makes the code's intent clear.

### Go Example

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
```

## C# Equivalent

C# has both **implicit** and **explicit** conversions:

- **Implicit**: Safe conversions that don't lose data (e.g., `int` to `long`, `float` to `double`)
- **Explicit (casts)**: Conversions that might lose data, using `(Type)value` or `Convert` class syntax
- **Parse methods**: For converting strings to numbers

C# is more lenient than Go, allowing implicit conversions when safe. This can be convenient but may hide bugs.

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        int x = 3, y = 4;
        
        // Implicit conversion: int to double is safe
        double f = Math.Sqrt(x * x + y * y);
        
        // Explicit cast: double to uint (may lose data)
        uint z = (uint)f;
        
        Console.WriteLine($"{x} {y} {z}");

        // Alternative: Convert class
        uint z2 = Convert.ToUInt32(f);
        
        // Parse from string
        int parsed = int.Parse("42");
    }
}
```

## Key Differences

- **Explicit vs Implicit**: Go requires all type conversions to be explicit; C# allows implicit conversions when "safe"
- **Syntax**: Go uses `Type(value)`; C# uses `(Type)value` for casts
- **Safety**: Go's approach prevents accidental conversions; C#'s implicit conversions are convenient but can hide issues
- **Numeric Promotions**: C# automatically promotes smaller types in expressions; Go requires explicit conversion
- **String Conversions**: Both require explicit conversion for string ↔ numeric, but C# has more options (Parse, TryParse, Convert class)
- **Compile-Time Checks**: Go catches all type mismatches at compile time; C# allows some runtime-checked conversions
- **Use Case Alignment**: Both convert between types, but with different philosophies. Go's explicit-only approach makes code more verbose but clearer about what's happening. C#'s implicit conversions reduce boilerplate for "obvious" safe conversions like `int` to `long`.

**Example where they differ**:
```go
// Go: ERROR - must explicitly convert
var i int = 42
var f float64 = i  // ERROR!
var f float64 = float64(i)  // OK
```

```csharp
// C#: Works - implicit conversion
int i = 42;
double f = i;  // OK - implicit conversion
```

This is a philosophical difference: Go values explicitness and clarity; C# values convenience for safe operations. Neither is wrong, but developers switching languages must be aware of this difference.
