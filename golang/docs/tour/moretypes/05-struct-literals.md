# Struct Literals

## Go Concept

A struct literal creates a new struct value by listing field values. You can:

- List values in order: `Vertex{1, 2}`
- Name fields explicitly: `Vertex{X: 1, Y: 2}` or `Vertex{X: 1}` (omitted fields get zero values)
- Create a pointer with `&`: `&Vertex{1, 2}`

The `&` prefix returns a pointer to the newly allocated struct.

### Go Example

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```

## C# Equivalent

C# has multiple ways to initialize structs:

- **Constructor**: `new Vertex(1, 2)`
- **Object initializer**: `new Vertex { X = 1, Y = 2 }`
- **Default**: `new Vertex()` or `default(Vertex)`
- **Target-typed new** (C# 9.0+): `Vertex v = new(1, 2)`

C# doesn't have the`&` operator for structs since structs are value types. To get a reference, you'd use a class or work with ref.

### C# Example

```csharp
using System;

struct Vertex
{
    public int X;
    public int Y;

    public Vertex(int x, int y)
    {
        X = x;
        Y = y;
    }

    public override string ToString() => $"{{X:{X} Y:{Y}}}";
}

class Program
{
    // Different initialization styles
    static Vertex v1 = new Vertex(1, 2);           // Constructor
    static Vertex v2 = new Vertex { X = 1 };       // Object initializer
    static Vertex v3 = new Vertex();               // Default constructor
    static Vertex v4 = default;                     // Default value

    static void Main()
    {
        Console.WriteLine($"{v1} {v2} {v3} {v4}");

        // C# 9.0+ target-typed new
        Vertex v5 = new(1, 2);
        
        // For reference semantics, use a class
        VertexClass vc = new VertexClass { X = 1, Y = 2 };
    }
}

class VertexClass
{
    public int X { get; set; }
    public int Y { get; set; }
}
```

## Key Differences

- **Positional Initialization**: Go supports `Vertex{1, 2}`; C# requires constructor or named members
- **Named Fields**: Both support named initialization; syntax differs slightly
- **Omitted Fields**: Both initialize omitted fields to zero/default values
- **Pointer Creation**: Go's `&Vertex{}` creates pointer; C# structs are value types (use classes for references)
- **Default Values**: Go uses `Vertex{}`; C# uses `new Vertex()`, `default(Vertex)`, or `default`
- **Use Case Alignment**: Both create struct instances with initial values. Go's syntax is more concise for simple cases. C#'s object initializers are more verbose but work with constructors and properties.

**Creation patterns**:
```go
// Go: Concise
v := Vertex{1, 2}
p := &Vertex{1, 2}
```

```csharp
// C#: More explicit
var v = new Vertex(1, 2);
// or
var v = new Vertex { X = 1, Y = 2 };
// For reference semantics, use class
var vc = new VertexClass { X = 1, Y = 2 };
```

Both achieve the same goal (initializing structs), but Go's syntax is terser. C#'s approach requires more ceremony but provides more flexibility with constructors and properties.
