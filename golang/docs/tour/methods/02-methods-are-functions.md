# Methods Are Functions

## Go Concept

A method is just a function with a receiver argument. You can rewrite any method as a regular function by making the receiver the first parameter. This demonstrates that methods are syntactic sugar for functions that take the receiver as an explicit argument.

This reveals the underlying mechanics: when you call `v.Abs()`, Go essentially calls `Abs(v)` behind the scenes.

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

// Method with receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Same operation as a regular function
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())     // Method call: 5
	fmt.Println(AbsFunc(v))  // Function call: 5
}
```

## C# Equivalent

In C#, instance methods have an implicit `this` parameter, but you can't directly convert them to static functions with the same ease. However, extension methods in C# are explicitly defined as static methods with a `this` parameter modifier, making them similar to Go's approach.

### C# Example

```csharp
using System;

public struct Vertex
{
    public double X { get; set; }
    public double Y { get; set; }
    
    // Instance method (implicit 'this' parameter)
    public double Abs()
    {
        return Math.Sqrt(this.X * this.X + this.Y * this.Y);
    }
}

public static class VertexHelpers
{
    // Extension method (explicit 'this' parameter)
    public static double AbsExtension(this Vertex v)
    {
        return Math.Sqrt(v.X * v.X + v.Y * v.Y);
    }
    
    // Regular static function
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
        Console.WriteLine(v.Abs());                     // Instance method: 5
        Console.WriteLine(v.AbsExtension());            // Extension method: 5
        Console.WriteLine(VertexHelpers.AbsFunc(v));   // Static function: 5
    }
}
```

## Key Differences

- **Transparency**: Go methods and functions are interchangeable; C# instance methods hide the `this` parameter
- **Extension Method Syntax**: C# extension methods use `this` modifier explicitly; Go methods use receiver syntax naturally
- **Conversion**: Go methods can be trivially rewritten as functions; C# instance methods cannot (only extension methods)
- **Conceptual Model**: Go treats methods as syntax for receiver functions; C# treats methods as fundamentally different from functions
- **Flexibility**: Go's approach makes it easy to choose between method and function style; C# locks you into the choice at definition time
- **Uniformity**: All Go methods work the same way; C# has three different patterns (instance methods, extension methods, static methods)
