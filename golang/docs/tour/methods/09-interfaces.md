# Interfaces

## Go Concept

An **interface type** is defined as a set of method signatures. A value of interface type can hold any value that implements those methods. Interfaces provide a way to specify the behavior of an object: if something can do *this*, then it can be used *here*.

Interfaces in Go are one of the key tools for abstraction and polymorphism.

### Go Example

```go
package main

import (
	"fmt"
	"math"
)

// Abser interface declares a single method
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // MyFloat implements Abser
	a = &v // *Vertex implements Abser
	// a = v // Error: Vertex does not implement Abser (Abs has pointer receiver)

	fmt.Println(a.Abs()) // Output: 5
}
```

## C# Equivalent

C# **interfaces** work similarly but require explicit implementation declaration using the `:` syntax. C# interfaces can declare methods, properties, events, and indexers. Starting with C# 8.0, interfaces can also have default implementations.

C# interfaces must be explicitly listed in the type declaration, making the relationship between types and interfaces more visible but less flexible.

### C# Example

```csharp
using System;

// Interface declaration
public interface IAbser
{
    double Abs();
}

// Struct implementing interface explicitly
public struct MyFloat : IAbser
{
    private readonly double value;
    
    public MyFloat(double value)
    {
        this.value = value;
    }
    
    public double Abs()
    {
        return value < 0 ? -value : value;
    }
}

// Class implementing interface
public class Vertex : IAbser
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
        IAbser a;
        
        a = new MyFloat(-Math.Sqrt(2));
        Console.WriteLine(a.Abs()); // Output: 1.41421356237309
        
        a = new Vertex { X = 3, Y = 4 };
        Console.WriteLine(a.Abs()); // Output: 5
    }
}
```

## Key Differences

- **Implicit vs Explicit**: Go interfaces are implemented implicitly (duck typing); C# requires explicit `: IInterface` declaration
- **Flexibility**: Go types can satisfy interfaces without knowing about them; C# types must declare interface implementation
- **Discoverability**: C# interfaces are visible in type declarations; Go interfaces are satisfied implicitly
- **Versioning**: Go can define new interfaces for existing types; C# requires types to know about interfaces upfront
- **Coupling**: Go interfaces reduce coupling; C# creates explicit dependencies
- **Refactoring**: Go interfaces can be added later; C# interfaces must be part of initial design
- **Type Checking**: C# checks interface implementation at type definition; Go checks at assignment/usage
- **Pointer Receivers**: Go distinguishes `T` vs `*T` implementing interfaces; C# has no such distinction
