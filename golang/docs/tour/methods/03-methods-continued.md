# Methods Continued

## Go Concept

You can declare methods on **any type** you define in your package, not just structs. However, you can only define methods on types defined in the same package. You cannot define methods directly on built-in types like `int`, but you can create a type alias and define methods on that.

This allows you to add behavior to simple types, making them more expressive and capable.

### Go Example

```go
package main

import (
	"fmt"
	"math"
)

// Define a custom type based on float64
type MyFloat float64

// Method on the custom numeric type
func (f MyFloat) Abs() MyFloat {
	if f < 0 {
		return -f
	}
	return f
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs()) // Output: 1.4142135623730951
	
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // Output: 5
}
```

## C# Equivalent

C# **extension methods** allow you to add methods to existing types without modifying them or creating derived types. However, they must be defined as static methods in static classes with the `this` modifier on the first parameter.

Unlike Go, you cannot create true type aliases with new methods in C#. You can create wrapper types (structs or classes), but they require explicit conversion.

### C# Example

```csharp
using System;

// Extension method on built-in type
public static class DoubleExtensions
{
    public static double Abs(this double f)
    {
        return f < 0 ? -f : f;
    }
}

public struct Vertex
{
    public double X { get; set; }
    public double Y { get; set; }
}

// Extension method on custom type
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
        double f = -Math.Sqrt(2);
        Console.WriteLine(f.Abs()); // Output: 1.41421356237309
        
        var v = new Vertex { X = 3, Y = 4 };
        Console.WriteLine(v.Abs()); // Output: 5
    }
}
```

## Key Differences

- **Type Aliases**: Go allows true type aliases with methods; C# extension methods work on the original type
- **Syntax Overhead**: Go methods look like regular methods; C# requires static class and `this` parameter syntax
- **Type Safety**: Go type aliases are distinct types; C# extension methods don't create new types
- **Package Restriction**: Go methods must be in the same package as the type; C# extension methods can be anywhere
- **Name Conflicts**: Go type aliases avoid conflicts; C# extension methods can conflict with instance methods
- **Discoverability**: Go methods show in type documentation; C# extension methods require namespace imports
- **Built-in Type Limitation**: Go requires type aliases for built-in types; C# extension methods work directly on built-in types
