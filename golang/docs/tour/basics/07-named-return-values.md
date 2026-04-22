# Named Return Values

## Go Concept

Go's return values can be **named**. Named return values are treated as variables defined at the top of the function and can be referenced and modified throughout the function body.

A `return` statement without arguments returns the current values of the named return values. This is called a "naked return" and should be used sparingly for clarity.

### Go Example

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return  // naked return returns x and y
}

func divmod(dividend, divisor int) (quotient, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor
	return  // returns quotient and remainder
}

func main() {
	fmt.Println(split(17))
	fmt.Println(divmod(17, 5))
}
```

## C# Equivalent

C# **does not have** named return values in the same way Go does. The closest equivalents are:

1. **Named tuple returns** (C# 7.0+): You can name tuple elements, but they're not variables in the function body
2. **Local variables**: You declare local variables and explicitly return them
3. **Out parameters**: An older pattern that's somewhat similar but works differently

None of these provide the same behavior as Go's named returns, where the names act as both documentation and as actual variables you can use in the function.

### C# Example

```csharp
using System;

class Program
{
    // Named tuple return (closest to Go's concept)
    static (int x, int y) Split(int sum)
    {
        int x = sum * 4 / 9;
        int y = sum - x;
        return (x, y);  // Must explicitly return
    }

    // Alternative: using named tuple in signature
    static (int quotient, int remainder) DivMod(int dividend, int divisor)
    {
        // These are local variables, not the return names
        int quotient = dividend / divisor;
        int remainder = dividend % divisor;
        return (quotient, remainder);  // Must explicitly return
    }

    // Old style with out parameters (different concept)
    static void DivModOut(int dividend, int divisor, out int quotient, out int remainder)
    {
        quotient = dividend / divisor;
        remainder = dividend % divisor;
        // No return statement needed for out parameters
    }

    static void Main()
    {
        Console.WriteLine(Split(17));
        Console.WriteLine(DivMod(17, 5));
        
        DivModOut(17, 5, out int q, out int r);
        Console.WriteLine($"({q}, {r})");
    }
}
```

## Key Differences

- **Named Returns as Variables**: Go's named returns are actual variables in the function; C# tuple names are only for the caller
- **Naked Returns**: Go supports returning without specifying values; C# always requires explicit return values
- **Documentation**: Go's named returns document what's being returned in the signature; C# tuple names do this too but don't create variables
- **Out Parameters**: C# has `out` which passes results through parameters rather than returns; Go doesn't have this concept
- **Explicitness**: C# always requires explicit `return (x, y)`; Go can use naked `return` with named returns
- **Use Case Alignment**: Both approaches aim to document what's being returned, but Go's named returns serve double duty as both documentation and working variables. This is **unique to Go** - C# requires you to explicitly return values. While C#'s named tuples provide similar documentation value, they don't create variables you can use in the function body. For simple functions, this difference is minor; for complex functions, Go's approach can reduce repetition but may harm clarity (thus Go's style guide recommends using naked returns only in short functions).
