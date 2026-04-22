# Pointers to Structs

## Go Concept

Struct fields can be accessed through a struct pointer. Go provides convenient syntax — you write `p.X` instead of `(*p).X`. This automatic dereferencing makes working with pointers more convenient.

### Go Example

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9  // Convenient: Go automatically dereferences
	fmt.Println(v)
}
```

## C# Equivalent

C# has **automatic dereferencing for references** (classes), but for unsafe pointers, you must use `->` operator or explicit dereferencing.

Most C# code uses classes (reference types) or ref parameters, not pointers:

- **Classes** (reference types): Automatic dereferencing with `.`
- **Ref parameters**: Access as if by value, but modifications affect original
- **Unsafe pointers**: Use `->` or explicit dereference `(*p).`

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
    static void Main()
    {
        // ref parameter approach (most common in C#)
        var v = new Vertex(1, 2);
        ModifyVertex(ref v);
        Console.WriteLine(v);

        // Unsafe pointer approach (rare)
        unsafe
        {
            Vertex v2 = new Vertex(1, 2);
            Vertex* p = &v2;
            p->X = 1000000000;  // Arrow operator for pointers
            // or: (*p).X = 1000000000;
            Console.WriteLine(v2);
        }

        // Class (reference type) - most idiomatic C#
        var vc = new VertexClass { X = 1, Y = 2 };
        vc.X = 1000000000;  // Automatic dereferencing
        Console.WriteLine($"{{X:{vc.X} Y:{vc.Y}}}");
    }

    static void ModifyVertex(ref Vertex v)
    {
        v.X = 1000000000;
    }
}

class VertexClass
{
    public int X { get; set; }
    public int Y { get; set; }
}
```

## Key Differences

- **Automatic Dereferencing**: Go uses `.` for both structs and pointers; C# uses `.` for references, `->` for unsafe pointers
- **Common Pattern**: Go commonly uses pointers to structs; C# commonly uses classes (references) or ref parameters
- **Pointer Syntax**: Go's `p.X` is convenient; C# unsafe pointers require `p->X` or `(*p).X`
- **Safety**: Go pointers are safe; C# pointers require unsafe context
- **Idioms**: Go passes `*Struct`; C# passes `ref Struct` or uses classes
- **Use Case Alignment**: Both provide ways to modify structs through indirection. **Go uses pointers as the standard approach**. **C# has three approaches**: classes (reference types, most common), ref parameters (for value types), or unsafe pointers (rare).

**Typical patterns**:
```go
// Go: Pointer to struct (common)
func modifyVertex(v *Vertex) {
    v.X = 100
}
```

```csharp
// C# Option 1: ref parameter (common for structs)
void ModifyVertex(ref Vertex v)
{
    v.X = 100;
}

// C# Option 2: class instead of struct (most common)
class Vertex { public int X; }
void ModifyVertex(Vertex v)  // Reference type, no ref needed
{
    v.X = 100;
}
```

The use case (modifying struct data efficiently) is the same, but C# developers typically either use classes or ref parameters rather than pointers. Go's approach with pointers is more uniform and simpler.
