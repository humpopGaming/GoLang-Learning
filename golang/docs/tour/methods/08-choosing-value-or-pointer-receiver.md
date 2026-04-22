# Choosing a Value or Pointer Receiver

## Go Concept

There are two reasons to use a pointer receiver:
1. **To modify the receiver**: The method needs to mutate the value that the receiver points to
2. **To avoid copying**: When the receiver is a large struct, using a pointer avoids copying on every method call

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. This consistency makes the type easier to reason about.

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

// Pointer receiver - modifies the value
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Pointer receiver - for consistency, even though it doesn't modify
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type LargeStruct struct {
	data [1000]float64
}

// Pointer receiver - avoids copying large struct
func (ls *LargeStruct) Sum() float64 {
	sum := 0.0
	for _, v := range ls.data {
		sum += v
	}
	return sum
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs()) // &{15 20} 25
	
	ls := &LargeStruct{}
	ls.data[0] = 10
	fmt.Println(ls.Sum()) // 10
}
```

## C# Equivalent

C# uses a different approach: choose between **class** (reference type) and **struct** (value type) when defining the type. This fundamental decision determines how instances are passed and stored:

- **Classes**: Always passed by reference, methods naturally modify the object, suitable for larger data
- **Structs**: Passed by value (copied), suitable for small immutable data, use `readonly` modifier for methods that don't modify

Modern C# encourages `readonly struct` for immutability and performance.

### C# Example

```csharp
using System;

// Class (reference type) - methods naturally modify the object
public class VertexClass
{
    public double X { get; set; }
    public double Y { get; set; }
    
    // Methods on classes don't need special syntax
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

// Struct (value type) - use readonly for non-mutating methods
public struct VertexStruct
{
    public double X { get; set; }
    public double Y { get; set; }
    
    // Mutating method
    public void Scale(double f)
    {
        X *= f;
        Y *= f;
    }
    
    // Readonly method - doesn't modify state
    public readonly double Abs()
    {
        return Math.Sqrt(X * X + Y * Y);
    }
}

// Large struct - prefer readonly to avoid defensive copies
public readonly struct LargeStruct
{
    private readonly double[] data;
    
    public LargeStruct(int size)
    {
        data = new double[size];
    }
    
    public double Sum()
    {
        double sum = 0;
        foreach (var v in data)
            sum += v;
        return sum;
    }
}

class Program
{
    static void Main()
    {
        var vc = new VertexClass { X = 3, Y = 4 };
        vc.Scale(5);
        Console.WriteLine($"{vc.X}, {vc.Y}, {vc.Abs()}"); // 15, 20, 25
    }
}
```

## Key Differences

- **Per-Method vs Per-Type**: Go chooses pointer/value per method; C# chooses class/struct per type
- **Mutability Control**: Go uses pointer receivers for mutation; C# uses mutable vs readonly structs
- **Mixing Receivers**: Go discourages mixing; C# doesn't have the concept (all methods on a type work the same way)
- **Copy Avoidance**: Go uses pointer receivers; C# uses classes (reference types)
- **Consistency**: Go recommends all pointer or all value receivers; C# enforces consistency through type choice
- **Performance**: Go pointer receivers avoid copies explicitly; C# classes avoid copies implicitly
- **Small Types**: Go can use value receivers; C# uses structs with readonly methods
- **Large Types**: Go uses pointer receivers; C# uses classes (reference types)
