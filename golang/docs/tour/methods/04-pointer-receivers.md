# Pointer Receivers

## Go Concept

Methods can have **pointer receivers**. This is indicated by the `*` before the type name in the receiver. Pointer receivers allow the method to modify the value that the receiver points to. Methods with value receivers operate on a copy of the original value.

Pointer receivers are essential when you need to modify the receiver or when the receiver is a large struct that would be expensive to copy.

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

// Pointer receiver - can modify the original value
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Value receiver - operates on a copy
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)           // Modifies v
	fmt.Println(v.Abs())  // Output: 50
}
```

## C# Equivalent

C# has different semantics based on whether a type is a **class** (reference type) or **struct** (value type). Instance methods on classes naturally modify the original object. For structs, methods modify a copy unless the struct is passed by reference or the method is called on a reference to the struct.

Extension methods in C# **cannot** use `ref this` to modify the receiver, which is a significant limitation compared to Go.

### C# Example

```csharp
using System;

// Struct (value type)
public struct VertexStruct
{
    public double X { get; set; }
    public double Y { get; set; }
    
    // Instance method on struct - 'this' is readonly by default
    // Use 'readonly' modifier to indicate no modification
    public readonly double Abs()
    {
        return Math.Sqrt(X * X + Y * Y);
    }
    
    // Mutable instance method - can modify 'this'
    public void Scale(double f)
    {
        X *= f;
        Y *= f;
    }
}

// Class (reference type)
public class VertexClass
{
    public double X { get; set; }
    public double Y { get; set; }
    
    // Instance method on class - naturally modifies the object
    public void Scale(double f)
    {
        X *= f;
        Y *= f;
    }
    
    public double Abs()
    {
        return Math.Sqrt(X * X + Y * Y);
    }
}

class Program
{
    static void Main()
    {
        // Struct example
        var vs = new VertexStruct { X = 3, Y = 4 };
        vs.Scale(10);
        Console.WriteLine(vs.Abs()); // Output: 50
        
        // Class example
        var vc = new VertexClass { X = 3, Y = 4 };
        vc.Scale(10);
        Console.WriteLine(vc.Abs()); // Output: 50
    }
}
```

## Key Differences

- **Explicit vs Implicit**: Go requires explicit `*` for pointer receivers; C# behavior depends on class vs struct
- **Extension Methods**: Go pointer receivers work everywhere; C# extension methods cannot use `ref this`
- **Struct Mutation**: Go pointer receivers enable struct mutation; C# struct methods mutate `this` directly but need careful handling
- **Default Behavior**: Go copies by default (value receiver); C# classes reference by default, structs copy
- **Clarity**: Go makes pointer/value distinction explicit in method signature; C# hides it in type declaration (class vs struct)
- **Performance**: Both use pointers to avoid copies of large data structures; Go makes the choice explicit per method
- **Safety**: Go's explicit syntax makes the intention clearer; C# requires understanding class vs struct semantics
