# Exercise: Errors

## Challenge

Copy your `Sqrt` function from the [earlier exercise](../../03-numeric-types.md) and modify it to return an `error` value.

`Sqrt` should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type:

```go
type ErrNegativeSqrt float64
```

and make it an `error` by giving it a:

```go
func (e ErrNegativeSqrt) Error() string
```

method such that `ErrNegativeSqrt(-2).Error()` returns `"cannot Sqrt negative number: -2"`.

**Note:** A call to `fmt.Sprint(e)` inside the `Error` method will send the program into an infinite loop. You can avoid this by converting `e` first: `fmt.Sprint(float64(e))`. Why?

## Starter Code

```go
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// TODO: Implement this method
	return ""
}

func Sqrt(x float64) (float64, error) {
	// TODO: Return an error for negative numbers
	// Otherwise, use Newton's method to calculate sqrt
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
```

## Expected Output

```
1.4142135623730951 <nil>
0 cannot Sqrt negative number: -2
```

## Hints

- Check if `x < 0` at the start of `Sqrt`
- Return `0, ErrNegativeSqrt(x)` for negative numbers
- In the `Error()` method, convert `e` to `float64` before formatting to avoid infinite recursion
- Use Newton's method: `z -= (z*z - x) / (2*z)` for positive numbers

## Solution

<details>
<summary>Click to reveal solution</summary>

```go
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))  // 1.4142135623730951 <nil>
	fmt.Println(Sqrt(-2)) // 0 cannot Sqrt negative number: -2
	
	// Demonstrating error handling
	if result, err := Sqrt(-4); err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
```

**Why the infinite loop?**

If you write `return fmt.Sprint(e)` in the Error() method:
1. fmt.Sprint calls e.Error() to get the string
2. e.Error() calls fmt.Sprint(e)
3. This creates infinite recursion!

Converting to float64 breaks the cycle because fmt.Sprint treats float64 differently.

</details>

## C# Equivalent Exercise

```csharp
using System;

// Create a custom exception for negative square root
class NegativeSqrtException : Exception
{
    public double Value { get; }
    
    public NegativeSqrtException(double value)
        : base($"cannot Sqrt negative number: {value}")
    {
        Value = value;
    }
}

class MathHelper
{
    // TODO: Implement Sqrt that throws exception for negative numbers
    public static double Sqrt(double x)
    {
        // Throw NegativeSqrtException if x < 0
        // Otherwise use Newton's method
        return 0;
    }
}

class Program
{
    static void Main()
    {
        try
        {
            Console.WriteLine(Sqrt(2));
            Console.WriteLine(Sqrt(-2));
        }
        catch (NegativeSqrtException ex)
        {
            Console.WriteLine($"Error: {ex.Message}");
        }
    }
}
```

**C# Solution:**

```csharp
public static double Sqrt(double x)
{
    if (x < 0)
        throw new NegativeSqrtException(x);
    
    double z = 1.0;
    for (int i = 0; i < 10; i++)
    {
        z -= (z * z - x) / (2 * z);
    }
    return z;
}
```

**C# Result Type Pattern (Go-like approach):**

```csharp
class Result<T>
{
    public T Value { get; }
    public string Error { get; }
    public bool IsSuccess => Error == null;
    
    public static Result<T> Success(T value) => new Result<T> { Value = value };
    public static Result<T> Failure(string error) => new Result<T> { Error = error };
}

public static Result<double> SqrtResult(double x)
{
    if (x < 0)
        return Result<double>.Failure($"cannot Sqrt negative number: {x}");
    
    double z = 1.0;
    for (int i = 0; i < 10; i++)
        z -= (z * z - x) / (2 * z);
    
    return Result<double>.Success(z);
}

// Usage
var result = SqrtResult(-2);
if (result.IsSuccess)
    Console.WriteLine(result.Value);
else
    Console.WriteLine($"Error: {result.Error}");
```

## Extension Challenges

1. **Error Wrapping**: Make Sqrt call another function and wrap its error with additional context

2. **Multiple Error Types**: Create different error types for different conditions (negative, NaN, infinity)

3. **Error Validation**: Add an `Is` method to check error types without full type assertion

4. **Sentinel Errors**: Create package-level error variables like `var ErrNegative = errors.New("negative value")`

5. **C# Comparison**: Implement both exception-based and Result-based versions in C# and compare

## Key Learning Points

- Custom error types implement the `error` interface by defining `Error() string`
- Errors are returned as regular values in Go
- Type conversion prevents infinite recursion in Error() methods
- Go encourages explicit error checking at each step
- C# uses exceptions for error conditions (different paradigm)
- Both approaches have tradeoffs in verbosity, performance, and clarity
