# Numeric Constants

## Go Concept

Numeric constants are **high-precision values**. An untyped constant takes the type needed by its context.

Untyped numeric constants can be arbitrarily large until they're assigned to a typed variable. They don't overflow until you give them a specific type. This allows writing large numbers and precise calculations without worrying about type limits.

### Go Example

```go
package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big))  // This would overflow and cause an error
}
```

## C# Equivalent

C# **does not have** untyped high-precision constants in the same way. All C# numeric literals have specific types, and constants must fit within those type limits:

- Integer literals default to `int` (32-bit), or the smallest type that fits
- Floating-point literals default to `double`
- You can't create constants larger than `ulong` (64-bit) or `decimal` (128-bit) limits

For arbitrary-precision arithmetic, C# has **`BigInteger`** from `System.Numerics`, but it cannot be used with `const` — you must use `readonly static` and accept runtime initialization.

### C# Example

```csharp
using System;
using System.Numerics;

class Program
{
    // Cannot use const with BigInteger
    static readonly BigInteger Big = BigInteger.Parse("1267650600228229401496703205376");
    static readonly BigInteger Small = Big >> 99;  // 2

    static int NeedInt(int x) => x * 10 + 1;
    static double NeedFloat(double x) => x * 0.1;

    static void Main()
    {
        Console.WriteLine(NeedInt((int)Small));
        Console.WriteLine(NeedFloat((double)Small));
        // Console.WriteLine(NeedInt((int)Big));  // Would overflow
        
        // Working with BigInteger directly
        Console.WriteLine(Big);
        Console.WriteLine(Small);
    }
}

// Alternative for compile-time calculations within type limits:
class Constants
{
    const int Large = 1 << 30;   // OK: fits in int
    // const int TooBig = 1 << 100;  // ERROR: doesn't fit in any const type
}
```

## Key Differences

- **Untyped Constants**: Go has untyped high-precision constants; C# constants must fit in a specific type
- **Arbitrary Precision**: Go constants can be arbitrarily large until assignment; C# limits constants to fixed-size types
- **BigInteger**: C# has `BigInteger` for large numbers, but it can't be `const` (must be `readonly` with runtime initialization)
- **Overflow Checking**: Go checks overflow when assigning to typed variables; C# checks at constant declaration time
- **Context-Dependent Type**: Go constants take type from context; C# literals have fixed types
- **Bit Shifting**: Both support bit shifting, but Go can shift beyond normal type limits in constants
- **Use Case Alignment**: This is a **significant difference without a direct equivalent**. Go's untyped constants allow mathematical expressions with large numbers in constants, automatically fitting to the needed type at usage. C# requires choosing a type upfront or using `BigInteger` at runtime. 

**When they differ**:
- Go: Can write `const Big = 1 << 100` and use it in float context without issue
- C#: Must use `BigInteger` or compute at runtime; cannot have arbitrarily large `const` values

This is a unique Go feature that provides flexibility for mathematical constants. In C#, you work around it by:
1. Using smaller constants that fit in `ulong` or `decimal`
2. Using `BigInteger` with `readonly static` (runtime, not compile-time)
3. Computing values at runtime

For most practical purposes, C#'s approach works fine, but Go's untyped constants are more elegant for mathematical code with large numbers.
