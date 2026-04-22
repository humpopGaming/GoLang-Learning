# Map Literals

## Go Concept

**Map literals** are like struct literals, but the keys are required.

- Map literals provide a convenient way to initialize maps with data
- Each element must have a key specified
- Useful for creating and populating maps in one expression
- The type is inferred from the declaration when using var or := with a literal

### Go Example

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
```

## C# Equivalent

C# has **collection initializers** for dictionaries, introduced in C# 3.0:

- **Collection initializer syntax**: Uses `{ }` with key-value pairs
- **Index initializer syntax** (C# 6+): Uses `["key"] = value` syntax
- Must specify the dictionary type explicitly or use `var` with `new Dictionary<,>`
- Provides a concise way to initialize and populate dictionaries

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
        // Collection initializer (C# 3.0+)
        var m = new Dictionary<string, Vertex>
        {
            { "Bell Labs", new Vertex(40.68433, -74.39967) },
            { "Google", new Vertex(37.42202, -122.08408) }
        };

        // Index initializer (C# 6+) - more concise
        var m2 = new Dictionary<string, Vertex>
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

- **Syntax**: Go uses `map[string]Type{...}`; C# uses `new Dictionary<string, Type> {...}` or index initializers
- **Key Specification**: Go uses `"key": value,`; C# uses `{ "key", value }` or `["key"] = value`
- **Type Declaration**: Go can omit `Vertex` in some contexts; C# requires `new Vertex(...)` or constructor
- **Trailing Commas**: Go requires trailing comma after last element; C# makes it optional
- **Conciseness**: Go's syntax is slightly more concise; C# requires `new` keyword
- **Type Inference**: Go infers the entire map type from literal; C# requires `var` with `new Dictionary<,>` or explicit type
- **Multiple Syntaxes**: Go has one syntax; C# evolved through multiple versions (collection initializer, index initializer)

**Different philosophies**:
- Go: Single, consistent literal syntax built into the language
- C#: Multiple initialization styles evolved over language versions

Both provide readable ways to create and initialize maps/dictionaries with data. C#'s index initializer syntax (`["key"] = value`) is closest to Go's style. The key difference is that C# always requires the `new Dictionary<,>` instantiation, while Go's literals are more integrated into the type system.
