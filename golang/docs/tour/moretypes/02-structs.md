# Structs

## Go Concept

A `struct` is a collection of fields. Structs group together data to create custom types.

Structs in Go are value types — they're copied when assigned or passed to functions (unless you use pointers).

### Go Example

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

## C# Equivalent

C# has both **structs** and **classes**:

- **struct**: Value type, stack-allocated, copied on assignment
- **class**: Reference type, heap-allocated, referenced on assignment

C# structs are similar to Go structs (both value types), but C# also has classes which are reference types. Go doesn't have this distinction — you use pointers when you want reference semantics.

### C# Example

```csharp
using System;

// C# struct (value type, like Go struct)
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

// C# class (reference type, different from Go)
class VertexClass
{
    public int X;
    public int Y;

    public VertexClass(int x, int y)
    {
        X = x;
        Y = y;
    }
}

class Program
{
    static void Main()
    {
        var v = new Vertex(1, 2);
        Console.WriteLine(v);

        // Value type behavior (copied)
        var v2 = v;
        v2.X = 10;
        Console.WriteLine($"v.X: {v.X}, v2.X: {v2.X}");  // v.X still 1

        // Reference type behavior (referenced)
        var vc = new VertexClass(1, 2);
        var vc2 = vc;
        vc2.X = 10;
        Console.WriteLine($"vc.X: {vc.X}");  // vc.X is now 10
    }
}
```

## Key Differences

- **Struct vs Class**: Go only has structs; C# has both structs (value) and classes (reference)
- **Default Semantics**: Go structs are value types; you use pointers for reference semantics. C# chooses with struct/class keywords
- **Constructors**: C# structs can have constructors; Go doesn't have constructors (use factory functions)
- **Inheritance**: C# structs can implement interfaces; Go structs can too (through methods)
- **Field Visibility**: Go uses capitalization; C# uses `public`/`private` keywords
- **Methods**: Both can have methods attached to structs
- **Use Case Alignment**: Both define custom data types with multiple fields. **Key difference**: C# makes you choose between value (struct) and reference (class) semantics upfront. Go uses structs for everything and adds pointers when you need reference semantics.

**When to use each**:
- Go: Always use struct, add pointer when needed (`*Vertex`)
- C#: Use struct for small value types (coordinates, colors); use class for most objects

Go's approach is simpler (one concept), C#'s approach is more explicit (two concepts). Both work well for their respective ecosystems.
