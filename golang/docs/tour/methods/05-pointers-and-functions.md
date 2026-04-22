# Pointers and Functions

## Go Concept

When comparing methods with pointer receivers to functions with pointer parameters, the behavior is identical. A function with a pointer parameter must receive a pointer argument, while methods with pointer receivers can be called on either a pointer or a value (Go handles the conversion automatically).

This demonstrates the relationship between Go's method syntax and its underlying function mechanics.

### Go Example

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Function with pointer parameter
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Function with value parameter
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Method with pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Method with value receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	ScaleFunc(&v, 10)      // Must pass pointer to function
	fmt.Println(AbsFunc(v)) // Output: 50
	
	v = Vertex{3, 4}
	v.Scale(10)            // Method works with value (Go converts automatically)
	fmt.Println(v.Abs())   // Output: 50
}
```

## C# Equivalent

C# uses **ref** and **out** parameters to pass value types by reference, and **in** for readonly references. For reference types (classes), you always pass a reference to the object. C# does not automatically convert between value and reference passing - you must be explicit.

### C# Example

```csharp
using System;

public struct Vertex
{
    public double X { get; set; }
    public double Y { get; set; }
    
    // Instance method that modifies the struct
    public void Scale(double f)
    {
        X *= f;
        Y *= f;
    }
    
    public readonly double Abs()
    {
        return Math.Sqrt(X * X + Y * Y);
    }
}

public static class VertexFunctions
{
    // Function with ref parameter (pass by reference)
    public static void ScaleFunc(ref Vertex v, double f)
    {
        v.X *= f;
        v.Y *= f;
    }
    
    // Function with value parameter (pass by value)
    public static double AbsFunc(Vertex v)
    {
        return Math.Sqrt(v.X * v.X + v.Y * v.Y);
    }
}

class Program
{
    static void Main()
    {
        var v = new Vertex { X = 3, Y = 4 };
        VertexFunctions.ScaleFunc(ref v, 10); // Must use 'ref' keyword
        Console.WriteLine(VertexFunctions.AbsFunc(v)); // Output: 50
        
        v = new Vertex { X = 3, Y = 4 };
        v.Scale(10); // Method mutates the struct in place
        Console.WriteLine(v.Abs()); // Output: 50
    }
}
```

## Key Differences

- **Automatic Conversion**: Go methods auto-convert `v.Scale()` to `(&v).Scale()`; C# requires explicit `ref` keyword
- **Function Call Syntax**: Go functions require `&v`; C# requires `ref v` at call site
- **Method Convenience**: Go methods handle pointer/value automatically; C# struct methods work on the struct directly
- **Reference Types**: C# classes always pass references; Go must explicitly use pointers
- **Syntactic Sugar**: Go's method syntax provides convenience; C# makes the distinction explicit everywhere
- **Safety**: C# `ref` at call site makes mutation visible; Go's auto-conversion is more convenient but less explicit
- **Out Parameters**: C# has `out` for output-only parameters; Go returns multiple values instead
