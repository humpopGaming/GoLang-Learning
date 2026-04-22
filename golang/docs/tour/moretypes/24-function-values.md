# Function Values

## Go Concept

**Functions are values too**. They can be passed around just like other values.

- Function values may be used as function arguments and return values
- Functions are **first-class citizens** in Go
- This enables higher-order functions, callbacks, and functional programming patterns
- Function types are specified by their parameter and return types

### Go Example

```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))  // 13

	fmt.Println(compute(hypot))  // 5
	fmt.Println(compute(math.Pow))  // 81

	// Inline anonymous function
	fmt.Println(compute(func(x, y float64) float64 {
		return x + y
	}))  // 7
}
```

**Key patterns**:
- Assign functions to variables: `hypot := func(...) {...}`
- Pass functions as arguments: `compute(hypot)`
- Define function types: `func(float64, float64) float64`
- Create inline anonymous functions

## C# Equivalent

C# has **delegates**, **Func<>**, and **Action<>** for function values:

- **Func<T, TResult>**: Generic delegate for functions that return a value
- **Action<T>**: Generic delegate for functions that return void
- **Lambda expressions**: Concise syntax for anonymous functions `(x, y) => x + y`
- **Delegate keyword**: Traditional delegate syntax (less common in modern C#)
- Functions are first-class values in C# too

### C# Example

```csharp
using System;

class Program
{
    // Method that takes a function as parameter
    static double Compute(Func<double, double, double> fn)
    {
        return fn(3, 4);
    }

    static void Main()
    {
        // Assign lambda to variable
        Func<double, double, double> hypot = (x, y) => 
            Math.Sqrt(x * x + y * y);
        
        Console.WriteLine(hypot(5, 12));  // 13

        Console.WriteLine(Compute(hypot));  // 5
        Console.WriteLine(Compute(Math.Pow));  // 81

        // Inline anonymous function (lambda)
        Console.WriteLine(Compute((x, y) => x + y));  // 7

        // Alternative: anonymous delegate (older style)
        Console.WriteLine(Compute(delegate(double x, double y) 
        {
            return x * y;
        }));  // 12

        // Action<> for void functions
        Action<string> print = message => Console.WriteLine(message);
        print("Hello from function value!");

        // Custom delegate (traditional approach)
        BinaryOperation add = (a, b) => a + b;
        Console.WriteLine(add(10, 20));  // 30
    }

    // Custom delegate type (traditional)
    delegate double BinaryOperation(double x, double y);
}
```

## Key Differences

- **Type Syntax**: Go uses `func(params) returns`; C# uses `Func<params, return>` or `Action<params>` or custom delegates
- **Generic Delegates**: Go has no generics for functions (pre-1.18, and limited after); C# has `Func<>` and `Action<>`
- **Lambda Syntax**: Go uses `func(x, y float64) float64 { ... }`; C# uses `(x, y) => ...` or `(x, y) => { ... }`
- **Type Inference**: Go requires full function signature in literals; C# can infer lambda parameter types
- **Void Functions**: Go uses zero return values; C# distinguishes with `Action<>` vs `Func<>`
- **Delegate Types**: C# has explicit delegate types; Go's function types are structural
- **Conciseness**: C# lambda syntax is more concise for simple functions; Go requires `func` keyword
- **Method References**: Both allow passing named functions directly

**Different philosophies**:
- Go: Functions are a simple, built-in type with straightforward syntax
- C#: Evolved from delegates to Func<>/Action<> to lambda expressions

**Syntax comparison**:
```go
// Go
fn := func(x, y int) int { return x + y }
```

```csharp
// C#
Func<int, int, int> fn = (x, y) => x + y;
```

**When to use**:
- **Go**: Use function values for callbacks, higher-order functions, strategy patterns
- **C#**: Same use cases, but choose `Func<>`, `Action<>`, or custom delegates based on void vs non-void

Both languages treat functions as first-class citizens, enabling functional programming patterns. C#'s lambda syntax is more concise for simple cases, while Go's syntax is more explicit and consistent. The underlying capability and use cases are essentially identical.
