# Range

## Go Concept

The **range** form of the `for` loop iterates over a slice or map, returning two values:
- **index** (position)
- **value** (copy of element)

```go
for index, value := range slice {
    // use index and value
}
```

### Go Example

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

Output:
```
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128
```

## C# Equivalent

C# provides several iteration mechanisms:

- **foreach**: Iterates over collection (no index by default)
- **for with index**: Traditional indexed loop
- **Enumerable.Select**: LINQ with index
- **foreach with Index/Value tuples** (C# 7.0+)

**foreach** is most similar to Go's range, but doesn't provide index by default.

### C# Example

```csharp
using System;
using System.Linq;

class Program
{
    static void Main()
    {
        int[] pow = { 1, 2, 4, 8, 16, 32, 64, 128 };

        // Option 1: foreach (value only, like Go range with only value)
        Console.WriteLine("foreach (value only):");
        foreach (int v in pow)
        {
            Console.WriteLine($"Value: {v}");
        }

        // Option 2: traditional for loop (index and value)
        Console.WriteLine("\nfor loop (index + value):");
        for (int i = 0; i < pow.Length; i++)
        {
            Console.WriteLine($"2**{i} = {pow[i]}");
        }

        // Option 3: LINQ Select with index (functional style)
        Console.WriteLine("\nLINQ Select:");
        pow.Select((value, index) => new { index, value })
           .ToList()
           .ForEach(x => Console.WriteLine($"2**{x.index} = {x.value}"));

        // Option 4: foreach with tuples (C# 7.0+, cleaner)
        Console.WriteLine("\nforeach with Index:");
        foreach (var (value, index) in pow.Select((v, i) => (v, i)))
        {
            Console.WriteLine($"2**{index} = {value}");
        }

        // Option 5: Index/Item pattern (C# 9.0+)
        Console.WriteLine("\nPattern with range operator:");
        for (int i = 0; i < pow.Length; i++)
        {
            Console.WriteLine($"2**{i} = {pow[i]}");
        }
    }
}
```

Output:
```
foreach (value only):
Value: 1
Value: 2
...

for loop (index + value):
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128

LINQ Select:
2**0 = 1
2**1 = 2
...
```

## Key Differences

- **Index Access**: Go range provides index by default; C# foreach doesn't (must use for loop or LINQ)
- **Syntax Simplicity**: Go `for i, v := range slice` is concise; C# requires more verbose patterns
- **Value Copies**: Go range creates copy of value; C# foreach iterates by reference (for reference types) or by value (for value types)
- **Built-in Feature**: Go range is language keyword; C# foreach is simpler (value-only by default)
- **Common Usage**: Go range is ubiquitous; C# developers choose between foreach (simple) and for (with index)
- **Use Case Alignment**:

**Iteration patterns comparison**:

| Need | Go | C# |
|------|----|----|
| Value only | `for _, v := range s` | `foreach (var v in s)` |
| Index + value | `for i, v := range s` | `for (int i = 0; i < s.Length; i++)` |
| Index only | `for i := range s` | `for (int i = 0; i < s.Length; i++)` |

**Different defaults**:
- **Go**: range gives you both index and value (discard what you don't need)
- **C#**: foreach gives you value only (use for loop if you need index)

**Key insight**: Go's range is designed for the common case of needing both index and value, with easy syntax to discard either. C#'s foreach is designed for the common case of only needing the value.

**When to use what in C#**:
```csharp
// Need only values? Use foreach
foreach (var item in collection)
    Process(item);

// Need index? Use for loop
for (int i = 0; i < array.Length; i++)
    Console.WriteLine($"{i}: {array[i]}");

// Need index in LINQ chain? Use Select overload
items.Select((item, index) => ...)
```

**Go range benefits**:
- Consistent syntax for arrays, slices, maps, channels
- Always safe (no index out of bounds)
- Clear intent (iterating over collection, not manual index management)

**C# foreach benefits**:
- Simpler for common case (value only)
- Works with any IEnumerable<T>
- Clearer when index isn't needed

**Performance note**: Both Go range and C# foreach are efficient. Go range may copy values (problematic for large structs). C# foreach behavior depends on whether type is value type or reference type.

**Common pattern in Go**:
```go
for i, item := range items {
    // Both index and value available
}
```

**Equivalent C# (most readable)**:
```csharp
for (int i = 0; i < items.Length; i++)
{
    var item = items[i];
    // Both index and value available
}
```

The Go approach is more uniform (one construct for all iteration), while C# offers optimized constructs for specific use cases (foreach for simple iteration, for for indexed iteration).
