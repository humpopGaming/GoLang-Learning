# Methods and Pointer Indirection (2)

## Go Concept

The automatic indirection works in the reverse direction as well. Methods with value receivers can be called on both values and pointers. When you call a value receiver method on a pointer, Go automatically dereferences the pointer.

For example, if `p` is a pointer and `Abs()` has a value receiver, calling `p.Abs()` is interpreted as `(*p).Abs()`. This convenience makes methods more ergonomic to use than functions.

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

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())     // OK: direct value call
	fmt.Println(AbsFunc(v))  // OK: function with value
	
	p := &Vertex{4, 3}
	fmt.Println(p.Abs())     // OK: Go interprets as (*p).Abs()
	// fmt.Println(AbsFunc(p)) // Compile error: cannot use p (type *Vertex) as type Vertex
	fmt.Println(AbsFunc(*p)) // OK: explicit dereference
}
```

## C# Equivalent

C# handles this differently based on type:
- **Classes** (reference types): Methods are always called through references, so there's no distinction
- **Structs** (value types): Instance methods can be called on struct variables, and the compiler handles any necessary copying

C# doesn't have the pointer/value method distinction that Go has, so there's no automatic dereferencing in the same sense.

### C# Example

```csharp
using System;

public struct Vertex
{
    public double X { get; set; }
    public double Y { get; set; }
    
    // Instance method with readonly modifier (doesn't modify struct)
    public readonly double Abs()
    {
        return Math.Sqrt(X * X + Y * Y);
    }
}

public static class VertexFunctions
{
    // Function taking value parameter
    public static double AbsFunc(Vertex v)
    {
        return Math.Sqrt(v.X * v.X + v.Y * v.Y);
    }
}

public class VertexClass
{
    public double X { get; set; }
    public double Y { get; set; }
    
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
        var v = new Vertex { X = 3, Y = 4 };
        Console.WriteLine(v.Abs());                    // OK: method call
        Console.WriteLine(VertexFunctions.AbsFunc(v)); // OK: function call
        
        // Class example (always works through references)
        var c = new VertexClass { X = 4, Y = 3 };
        Console.WriteLine(c.Abs()); // OK: reference type
    }
}
```

## Key Differences

- **Bidirectional Convenience**: Go auto-converts in both directions (`v.Method()` with pointer receiver, `p.Method()` with value receiver); C# has no equivalent
- **Method vs Function**: Go's convenience only applies to methods; C# functions require exact type matching
- **Pointer Syntax**: Go uses `*` for explicit dereference; C# rarely requires explicit dereferencing
- **Reference Types**: C# classes eliminate pointer/value distinction; Go makes it explicit
- **Struct Semantics**: C# struct methods always work on the struct value; Go distinguishes between value and pointer receivers
- **Syntactic Sugar**: Go provides maximum convenience for method calls; C# keeps reference/value semantics simpler but less flexible
- **Explicitness Trade-off**: Go hides pointer mechanics for ergonomics; C# makes struct/class distinction fundamental but simpler
