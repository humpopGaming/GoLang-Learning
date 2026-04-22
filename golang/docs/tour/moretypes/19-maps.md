# Maps

## Go Concept

A **map** maps keys to values. Maps are Go's built-in associative data type (sometimes called hashes or dicts in other languages).

- The zero value of a map is `nil`
- A `nil` map has no keys, nor can keys be added
- The `make` function returns a map of the given type, initialized and ready to use
- Maps are reference types — when you assign a map to a new variable, both refer to the same underlying data

### Go Example

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```

## C# Equivalent

C# has **Dictionary<TKey, TValue>** as the primary map/dictionary type:

- **Dictionary<TKey, TValue>**: Generic hash table, most common choice
- Must be instantiated before use (no concept of nil dictionary)
- Also has **Hashtable** (non-generic, legacy) and **ConcurrentDictionary** (thread-safe)
- Dictionaries are reference types like maps

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
        // Must instantiate before use
        Dictionary<string, Vertex> m = new Dictionary<string, Vertex>();
        
        m["Bell Labs"] = new Vertex(40.68433, -74.39967);
        
        Console.WriteLine($"Lat: {m["Bell Labs"].Lat}, Long: {m["Bell Labs"].Long}");
        
        // Modern C# (C# 9+)
        Dictionary<string, Vertex> m2 = new();
        m2["Bell Labs"] = new Vertex(40.68433, -74.39967);
    }
}
```

## Key Differences

- **Declaration Syntax**: Go uses `map[KeyType]ValueType`; C# uses `Dictionary<TKey, TValue>`
- **Instantiation**: Go uses `make(map[KeyType]ValueType)` or map literals; C# uses `new Dictionary<TKey, TValue>()`
- **Zero Value**: Go has `nil` maps that can't be used; C# doesn't have nil dictionaries (must instantiate)
- **Reference Types**: Both are reference types — assigning to a new variable creates an alias, not a copy
- **Type Parameters**: Go uses square brackets for key/value types; C# uses angle brackets (generics)
- **Thread Safety**: Go maps are not thread-safe (need sync.Mutex); C# has `ConcurrentDictionary<TKey, TValue>` for thread-safe scenarios
- **Nil Map Behavior**: Go allows declaring `var m map[string]int` (nil map, reads return zero value, writes panic); C# requires instantiation
- **Access Pattern**: Both use bracket notation `m["key"]` for access

**Different philosophies**:
- Go: Explicit creation with `make`, allows nil maps (read-only)
- C#: Must instantiate, no nil dictionary concept

The core functionality is nearly identical — both provide O(1) average-case lookups and are reference types. The main difference is Go's nil map concept vs C#'s requirement to always instantiate.
