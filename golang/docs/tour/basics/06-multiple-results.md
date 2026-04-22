# Multiple Results

## Go Concept

A Go function can return **multiple values**. This is one of Go's most distinctive features and is commonly used for error handling (returning both a result and an error).

Multiple return values are declared by listing them in parentheses after the function parameters. You can return values separated by commas.

### Go Example

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func divmod(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	q, r := divmod(17, 5)
	fmt.Println("Quotient:", q, "Remainder:", r)
}
```

## C# Equivalent

C# historically did not have built-in multiple return values. Modern C# (7.0+) introduced **tuples** which provide similar functionality:

- **ValueTuple** syntax: `(string, string)` for return type, `(y, x)` for return value
- **Named tuples**: `(int quotient, int remainder)` for better readability
- **Deconstruction**: `var (a, b) = swap("hello", "world");`

Before C# 7.0, developers used `out` parameters, classes, or the older `Tuple<T1, T2>` class. The new tuple syntax is much closer to Go's approach.

### C# Example

```csharp
using System;

class Program
{
    // C# 7.0+ tuple syntax
    static (string, string) Swap(string x, string y)
    {
        return (y, x);
    }

    // Named tuple elements (recommended)
    static (int quotient, int remainder) DivMod(int x, int y)
    {
        return (x / y, x % y);
    }

    static void Main()
    {
        // Deconstruction
        var (a, b) = Swap("hello", "world");
        Console.WriteLine($"{a} {b}");

        // Named access
        var result = DivMod(17, 5);
        Console.WriteLine($"Quotient: {result.quotient} Remainder: {result.remainder}");

        // Or deconstruct directly
        var (q, r) = DivMod(17, 5);
        Console.WriteLine($"Quotient: {q} Remainder: {r}");
    }
}
```

## Key Differences

- **Native Support**: Go had multiple returns from day one; C# added tuples in version 7.0 (2017)
- **Syntax**: Go uses simple syntax `(type1, type2)`; C# uses ValueTuple `(type1, type2)`
- **Named Returns**: Both support naming return values, but Go's named returns can be used as variables in the function body (covered in next section)
- **Error Handling Pattern**: Go idiomatically returns `(result, error)`; C# typically uses exceptions instead
- **Legacy Alternatives**: C# has `out` parameters as an older alternative; Go only has multiple returns
- **Performance**: Both compile to efficient code; C#'s ValueTuple is a struct, so it's stack-allocated
- **Use Case Alignment**: Both solve the problem of returning multiple values from a function. Go uses this for error handling (`result, err`), while C# traditionally uses exceptions. C#'s tuples are now commonly used for returning related values without creating custom types. The syntax differs slightly, but the concept and use cases align well.
