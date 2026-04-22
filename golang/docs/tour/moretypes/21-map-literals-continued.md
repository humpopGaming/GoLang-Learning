# Map Literals Continued

## Go Concept

If the top-level type is just a type name, you can **omit it from the elements of the literal**.

- When the value type is obvious, you can use a shorthand notation
- Simply provide the values in braces without repeating the type name
- Makes code more concise and readable
- The compiler infers the type from the map's declared type

### Go Example

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},  // Omitted "Vertex"
	"Google":    {37.42202, -122.08408}, // Omitted "Vertex"
}

func main() {
	fmt.Println(m)
}
```

Compare to the previous version where `Vertex` was explicitly stated:
```go
"Bell Labs": Vertex{40.68433, -74.39967},
```

## C# Equivalent

C# has **target-typed new expressions** (C# 9+) and **implicit object creation**:

- **Target-typed new** (C# 9+): Use `new()` when type can be inferred
- **Object initializers**: Can sometimes omit type with proper context
- Type inference from declaration simplifies code
- Earlier versions required full type specification

### C# Example

```csharp
using System;
using System.Collections.Generic;

class Program
{
    struct Vertex
    {
        public double Lat { get; set; }
        public double Long { get; set; }
        
        public Vertex(double lat, double lon)
        {
            Lat = lat;
            Long = lon;
        }
    }

    static void Main()
    {
        // C# 9+ target-typed new expression
        Dictionary<string, Vertex> m = new()
        {
            ["Bell Labs"] = new(40.68433, -74.39967),  // Omitted "Vertex"
            ["Google"] = new(37.42202, -122.08408)     // Omitted "Vertex"
        };

        // Or with var
        var m2 = new Dictionary<string, Vertex>
        {
            ["Bell Labs"] = new Vertex(40.68433, -74.39967),  // Must specify Vertex here
            ["Google"] = new Vertex(37.42202, -122.08408)
        };

        // Older C# versions required full specification
        Dictionary<string, Vertex> m3 = new Dictionary<string, Vertex>
        {
            ["Bell Labs"] = new Vertex(40.68433, -74.39967),
            ["Google"] = new Vertex(37.42202, -122.08408)
        };

        foreach (var kvp in m)
        {
            Console.WriteLine($"{kvp.Key}: Lat={kvp.Value.Lat}, Long={kvp.Value.Long}");
        }
    }
}
```

## Key Differences

- **Type Omission**: Go omits type name in composite literals when obvious; C# uses `new()` in C# 9+
- **Context Requirements**: Go infers from map's value type; C# infers from variable's declared type
- **Syntax**: Go uses `{values}` without type; C# uses `new(values)` or requires full type
- **Version Support**: Go had this from early versions; C# added target-typed new in C# 9 (2020)
- **Readability**: Both achieve similar conciseness in modern versions
- **Backward Compatibility**: Go syntax unchanged; C# evolved through versions
- **var Usage**: Go always allows omission with type context; C# with `var` requires explicit `Vertex` constructor

**Different philosophies**:
- Go: Type inference integrated early, consistent syntax
- C#: Type inference evolved gradually through language versions

Both languages recognize that repeating type names is verbose when the compiler can infer them. Go's approach is slightly more implicit (just braces), while C# requires the `new()` keyword to signal construction. The result is similarly concise code in modern versions of both languages.

**When to use**: Use type omission when the type is obvious from context. It reduces noise and improves readability. The explicit version remains valid and might be clearer in complex scenarios.
