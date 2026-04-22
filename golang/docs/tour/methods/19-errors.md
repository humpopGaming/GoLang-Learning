# Errors

## Go Concept

Go programs express error state with **error values**. The `error` type is a built-in interface defined as:

```go
type error interface {
    Error() string
}
```

Functions often return an `error` value, and calling code should handle errors by testing whether the error equals `nil`. A nil error denotes success; a non-nil error denotes failure.

**Error handling in Go is explicit and immediate** - you check and handle errors right where they occur, not through exception unwinding. This makes error paths visible in the code and forces programmers to think about error conditions.

### Go Example

```go
package main

import (
	"fmt"
	"time"
)

// Custom error type
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// Function that might fail
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %v by zero", a)
	}
	return a / b, nil
}

// Wrapping errors (Go 1.13+)
func processFile(filename string) error {
	// Simulate error
	err := fmt.Errorf("file not found: %s", filename)
	return fmt.Errorf("processFile failed: %w", err)
}

func main() {
	// Basic error handling
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	
	// Custom error
	if err := run(); err != nil {
		fmt.Println(err) // at 2026-04-21..., it didn't work
	}
	
	// Error unwrapping
	err = processFile("data.txt")
	fmt.Println(err) // processFile failed: file not found: data.txt
}
```

## C# Equivalent

C# uses **exceptions** as the primary error handling mechanism. Exceptions are thrown using `throw` and caught using `try-catch` blocks. This is fundamentally different from Go's explicit error returns.

**The paradigm difference is profound:**
- **Go**: Errors are values. They flow through normal return paths. You check `if err != nil`.
- **C#**: Errors are exceptional. They flow through separate exception paths. You wrap code in `try-catch`.

Modern C# also supports result types (discriminated unions) for functional-style error handling, but exceptions remain the dominant pattern.

### C# Example

```csharp
using System;

// Custom exception type
class MyException : Exception
{
    public DateTime When { get; }
    public string What { get; }
    
    public MyException(DateTime when, string what) 
        : base($"at {when}, {what}")
    {
        When = when;
        What = what;
    }
}

class Program
{
    // Method that throws exception
    static void Run()
    {
        throw new MyException(DateTime.Now, "it didn't work");
    }
    
    // Method that might fail
    static double Divide(double a, double b)
    {
        if (b == 0)
            throw new DivideByZeroException($"Cannot divide {a} by zero");
        return a / b;
    }
    
    // Wrapping exceptions
    static void ProcessFile(string filename)
    {
        try
        {
            // Simulate error
            throw new FileNotFoundException($"file not found: {filename}");
        }
        catch (FileNotFoundException ex)
        {
            throw new InvalidOperationException("processFile failed", ex);
        }
    }
    
    static void Main()
    {
        // Basic exception handling
        try
        {
            double result = Divide(10, 0);
            Console.WriteLine($"Result: {result}");
        }
        catch (DivideByZeroException ex)
        {
            Console.WriteLine($"Error: {ex.Message}");
        }
        
        // Custom exception
        try
        {
            Run();
        }
        catch (MyException ex)
        {
            Console.WriteLine(ex.Message);
        }
        
        // Exception wrapping
        try
        {
            ProcessFile("data.txt");
        }
        catch (InvalidOperationException ex)
        {
            Console.WriteLine(ex.Message);
            Console.WriteLine($"Inner: {ex.InnerException?.Message}");
        }
    }
}

// C# Result type pattern (functional approach)
class Result<T>
{
    public T Value { get; }
    public Exception Error { get; }
    public bool IsSuccess => Error == null;
    
    private Result(T value, Exception error)
    {
        Value = value;
        Error = error;
    }
    
    public static Result<T> Success(T value) => new Result<T>(value, null);
    public static Result<T> Failure(Exception error) => new Result<T>(default, error);
}

// Using Result type (Go-like pattern in C#)
static Result<double> DivideResult(double a, double b)
{
    if (b == 0)
        return Result<double>.Failure(new DivideByZeroException());
    return Result<double>.Success(a / b);
}
```

## Key Differences

- **Paradigm**: Go errors are return values; C# exceptions are thrown and caught
- **Control Flow**: Go errors are explicit in normal flow; C# exceptions jump to catch blocks
- **Syntax**: Go uses `if err != nil`; C# uses `try-catch-finally` blocks
- **Performance**: Go error returns are fast; C# exception throwing is expensive (but catching is cheap)
- **Visibility**: Go errors visible in function signatures; C# exceptions can be invisible (not in signature)
- **Handling Location**: Go handles errors immediately; C# can handle errors up the call stack
- **Multiple Errors**: Go can return multiple errors or wrap errors; C# has InnerException for wrapping
- **Nil/Null**: Go nil error means success; C# no exception means success
- **Idiomatic Use**: Go checks every error immediately; C# uses exceptions for exceptional cases
- **Design Philosophy**: Go: "Errors are expected, handle them explicitly"; C#: "Exceptions for exceptional situations"
- **Recovery**: Go uses explicit checks; C# uses catch blocks and finally for cleanup
- **Stack Traces**: C# exceptions capture stack traces automatically; Go errors don't (unless wrapped with stack info)
- **Type Safety**: Go errors are just values; C# exceptions have type hierarchy for different catch blocks
- **Code Verbosity**: Go error handling is more verbose; C# try-catch can be more concise for happy path
- **Mixed Approaches**: C# can use Result types for Go-like error handling, but exceptions are idiomatic

## When To Use Each Approach

**Go's explicit errors are better when:**
- Error conditions are part of normal operation
- You want errors to be visible in function signatures
- You need predictable performance
- You want to force callers to handle errors

**C#'s exceptions are better when:**
- Errors truly are exceptional (rare, unexpected)
- You want clean happy-path code
- You need to propagate errors up multiple call levels
- You want automatic stack traces and rich error information

Both approaches are valid; they represent different philosophies about error handling in system design.
