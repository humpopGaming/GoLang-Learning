# Methods

## Go Concept

Go doesn't have classes, but you can define methods on types. A method is a function with a special **receiver** argument that appears between the `func` keyword and the method name. The receiver gives the method access to the type's data.

Methods allow you to attach behavior to your custom types. The receiver can be any type you define in your package, not just structs.

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

// Abs method has a receiver of type Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // Output: 5
}
```

## C# Equivalent

C# has **instance methods** that are defined inside classes and structs. These methods have implicit access to the instance through the `this` keyword. C# also has **extension methods** which can add methods to existing types without modifying them.

### C# Example

```csharp
using System;

public struct Vertex
{
    public double X { get; set; }
    public double Y { get; set; }
    
    // Instance method
    public double Abs()
    {
        return Math.Sqrt(X * X + Y * Y);
    }
}

// Extension method (defined in a static class)
public static class VertexExtensions
{
    public static double Abs(this Vertex v)
    {
        return Math.Sqrt(v.X * v.X + v.Y * v.Y);
    }
}

class Program
{
    static void Main()
    {
        var v = new Vertex { X = 3, Y = 4 };
        Console.WriteLine(v.Abs()); // Output: 5
    }
}
```

## Key Differences

- **Definition Location**: Go methods are defined outside the type; C# instance methods are defined inside classes/structs
- **Receiver Syntax**: Go uses explicit receiver syntax `(v Vertex)`; C# uses implicit `this` access
- **Type Flexibility**: Go methods can be defined on any type (including built-in types via type aliases); C# instance methods only on classes/structs you control
- **Extension Methods**: Go's method syntax works like C# extension methods but is more natural and doesn't require special static class syntax
- **No Classes Required**: Go achieves object-oriented behavior without classes; C# requires classes or structs
- **Method Set**: Go methods are not part of the type declaration; C# methods are part of the type definition
