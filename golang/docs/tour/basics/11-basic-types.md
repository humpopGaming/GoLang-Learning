# Basic Types

## Go Concept

Go's basic types include:

- **bool**: Boolean values (`true` or `false`)
- **string**: String values
- **int**, **int8**, **int16**, **int32**, **int64**: Signed integers of various sizes
- **uint**, **uint8**, **uint16**, **uint32**, **uint64**: Unsigned integers
- **byte**: Alias for uint8
- **rune**: Alias for int32, represents a Unicode code point
- **float32**, **float64**: Floating-point numbers
- **complex64**, **complex128**: Complex numbers

The `int`, `uint`, and `uintptr` types are platform-dependent (32 or 64 bits depending on the system).

### Go Example

```go
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

## C# Equivalent

C# has similar basic types with some naming and behavioral differences:

- **bool**: Same as Go
- **string**: Same as Go (though implemented differently)
- **sbyte**, **short**, **int**, **long**: Signed integers (8, 16, 32, 64 bits)
- **byte**, **ushort**, **uint**, **ulong**: Unsigned integers (8, 16, 32, 64 bits)
- **char**: 16-bit Unicode character (similar to Go's rune but different)
- **float**, **double**, **decimal**: Floating-point and decimal types
- **Complex** (in System.Numerics): Complex numbers (not a primitive type)

Key differences: C# has fixed sizes for all integer types (no platform-dependent int), and C# includes `decimal` for precise decimal arithmetic.

### C# Example

```csharp
using System;
using System.Numerics;

class Program
{
    static bool ToBe = false;
    static ulong MaxInt = ulong.MaxValue;  // C# provides MaxValue constants
    static Complex z = Complex.Sqrt(new Complex(-5, 12));

    static void Main()
    {
        Console.WriteLine($"Type: {ToBe.GetType()} Value: {ToBe}");
        Console.WriteLine($"Type: {MaxInt.GetType()} Value: {MaxInt}");
        Console.WriteLine($"Type: {z.GetType()} Value: {z}");
    }
}
```

## Key Differences

- **Integer Sizes**: Go has platform-dependent `int`/`uint`; C# has fixed-size `int` (always 32-bit) and `IntPtr`/`UIntPtr` for platform-specific
- **Character Types**: Go uses `rune` (int32) for Unicode code points; C# uses `char` (16-bit, UTF-16 code unit)
- **Byte Alias**: Go's `byte` is alias for `uint8`; C# `byte` is directly uint8 with no sbyte alias
- **Complex Numbers**: Go has primitive `complex64`/`complex128`; C# has struct type `Complex` in a library
- **Decimal Type**: C# has `decimal` for financial calculations; Go uses packages for arbitrary precision
- **Default Sizes**: Go's `int` is platform-dependent; C#'s `int` is always 32 bits
- **Use Case Alignment**: Both provide a rich set of numeric types for different precision and performance needs. Go's approach with platform-dependent `int` is designed for performance (native word size), while C#'s fixed sizes provide predictability. For most use cases, the types map directly (Go's `int32` = C#'s `int`, Go's `float64` = C#'s `double`). The main exception is complex numbers, where Go has first-class support and C# requires a library.
