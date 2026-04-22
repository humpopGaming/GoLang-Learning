# Methods and Pointer Indirection

## Go Concept

Methods with pointer receivers can be called on both pointers and values. Go automatically takes the address of the value when needed. This is a convenience feature: when you call `v.Scale(5)` where `v` is a value and `Scale` has a pointer receiver, Go interprets it as `(&v).Scale(5)`.

This automatic indirection only works for methods, not functions. Functions with pointer parameters must be called with explicit pointer arguments.

### Go Example

```go
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)           // OK: Go interprets as (&v).Scale(2)
	// ScaleFunc(v, 10)  // Compile error: cannot use v (type Vertex) as type *Vertex
	
	p := &Vertex{4, 3}
	p.Scale(3)           // OK: pointer receiver with pointer value
	ScaleFunc(p, 8)      // OK: function with pointer parameter
	
	fmt.Printf("v = %v, p = %v\n", v, p) // v = {6 8}, p = &{96 72}
}
```

## C# Equivalent

C# has different behavior for **reference types** (classes) versus **value types** (structs). Classes always work through references, so there's no confusion. For structs, methods are called directly on the value, and any mutations affect the variable directly (not a copy, as methods receive a reference to the struct variable).

C# does not have automatic address-taking. You must explicitly use `ref` when calling functions that need reference parameters.

### C# Example

```csharp
using System;

public struct Vertex
{
    public double X { get; set; }
    public double Y { get; set; }
    
    // Instance method on struct - mutates in place
    public void Scale(double f)
    {
        X *= f;
        Y *= f;
    }
    
    public override string ToString() => $"({X}, {Y})";
}

public static class VertexFunctions
{
    // Function requiring ref parameter
    public static void ScaleFunc(ref Vertex v, double f)
    {
        v.X *= f;
        v.Y *= f;
    }
}

class Program
{
    static void Main()
    {
        var v = new Vertex { X = 3, Y = 4 };
        v.Scale(2);              // OK: struct method mutates v directly
        VertexFunctions.ScaleFunc(ref v, 5); // OK: must use 'ref' keyword
        
        Console.WriteLine($"v = {v}"); // v = (30, 40)
    }
}
```

## Key Differences

- **Automatic Conversion**: Go auto-converts `v.Scale()` to `(&v).Scale()`; C# has no such conversion
- **Method vs Function**: Go's automatic indirection only applies to methods; C# has no automatic indirection at all
- **Explicitness**: Go hides pointer details for convenience; C# requires explicit `ref` keyword
- **Reference Types**: C# classes always use references implicitly; Go requires explicit pointer types
- **Struct Behavior**: C# struct methods work on the struct variable directly; Go value receivers work on copies
- **Call Site Clarity**: C# `ref` makes mutation visible at call site; Go's automatic conversion is invisible
- **Consistency**: Go treats methods and functions differently; C# treats them more uniformly (both require explicit `ref`)
