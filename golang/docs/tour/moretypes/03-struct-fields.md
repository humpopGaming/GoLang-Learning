# Struct Fields

## Go Concept

Struct fields are accessed using a **dot** notation. This is straightforward and works the same whether you have a struct or a pointer to a struct (Go automatically dereferences pointers).

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
	v.X = 4
	fmt.Println(v.X)
}
```

## C# Equivalent

C# also uses **dot notation** for accessing struct and class fields/properties. The syntax is identical.

C# distinguishes between:
- **Fields**: Direct data members
- **Properties**: Methods that look like fields (getters/setters)

Go only has fields; C# has both concepts.

### C# Example

```csharp
using System;

struct Vertex
{
    public int X;  // Field
    public int Y;  // Field

    // Property (C# concept, no Go equivalent)
    public int Sum => X + Y;

    public Vertex(int x, int y)
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
        v.X = 4;  // Dot notation, same as Go
        Console.WriteLine(v.X);
        Console.WriteLine(v.Sum);  // Property access (looks like field)
    }
}
```

## Key Differences

- **Dot Notation**: Identical in both languages
- **Properties vs Fields**: C# has properties (computed values); Go only has fields
- **Automatic Dereferencing**: Go automatically dereferences pointers when accessing fields; C# does the same for references
- **Encapsulation**: C# uses properties for controlled access; Go uses methods (getters/setters) explicitly
- **Use Case Alignment**: Both access struct members with dots. The concepts align perfectly. **C# extends this with properties** which provide computed values that look like fields. Go doesn't have properties — you write methods explicitly.

Example of equivalent patterns:
```go
// Go: Explicit methods
type Rectangle struct {
    Width, Height int
}
func (r Rectangle) Area() int {
    return r.Width * r.Height
}
```

```csharp
// C#: Property (looks like a field)
struct Rectangle
{
    public int Width;
    public int Height;
    public int Area => Width * Height;  // Property
}
```

C#'s properties provide convenient syntax for computed values. Go achieves the same with methods but requires calling them like functions `area()` instead of accessing like fields `Area`.
