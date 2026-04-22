# Slice Literals

## Go Concept

A **slice literal** is like an array literal without the length. It creates an array and then builds a slice that references it.

Slice literal: `[]T{values}` — creates array and returns slice
Array literal: `[N]T{values}` — creates fixed-size array

### Go Example

```go
package main

import "fmt"

func main() {
	// Array literal (fixed size)
	arr := [3]bool{true, true, false}
	fmt.Printf("Array: %v, Type: %T\n", arr, arr)

	// Slice literal (no size specified)
	s := []bool{true, true, false}
	fmt.Printf("Slice: %v, Type: %T\n", s, s)

	// Slice literal with struct
	points := []struct {
		x, y int
	}{
		{0, 0},
		{1, 2},
		{2, 4},
	}
	fmt.Println(points)
}
```

Output:
```
Array: [true true false], Type: [3]bool
Slice: [true true false], Type: []bool
[{0 0} {1 2} {2 4}]
```

## C# Equivalent

C# has **collection initializers** for arrays and lists, but they work differently:

- **Array initialization**: `new int[] { 1, 2, 3 }` creates fixed-size array
- **List initialization**: `new List<int> { 1, 2, 3 }` creates dynamic list
- **No separate slice literal**: C# doesn't distinguish slice literals from array literals

### C# Example

```csharp
using System;
using System.Collections.Generic;

class Program
{
    static void Main()
    {
        // Array initialization (similar to Go array literal)
        int[] arr = new int[] { 1, 2, 3 };
        // Or: int[] arr = { 1, 2, 3 };
        Console.WriteLine($"Array: [{string.Join(", ", arr)}], Type: {arr.GetType()}");

        // List initialization (most similar to Go slice usage)
        List<int> list = new List<int> { 1, 2, 3 };
        Console.WriteLine($"List: [{string.Join(", ", list)}], Type: {list.GetType()}");

        // Collection with anonymous types
        var points = new[]
        {
            new { x = 0, y = 0 },
            new { x = 1, y = 2 },
            new { x = 2, y = 4 }
        };
        
        foreach (var p in points)
            Console.WriteLine($"{{{p.x} {p.y}}}");
    }
}
```

Output:
```
Array: [1, 2, 3], Type: System.Int32[]
List: [1, 2, 3], Type: System.Collections.Generic.List`1[System.Int32]
{0 0}
{1 2}
{2 4}
```

## Key Differences

- **Literal Syntax**: Go `[]int{1,2,3}` creates slice; C# `new int[] {1,2,3}` or `{1,2,3}` creates array
- **Type Distinction**: Go clearly distinguishes `[]T` (slice) from `[N]T` (array); C# treats both as arrays
- **Dynamic Behavior**: Go slice literals create dynamic slices; C# array literals create fixed arrays
- **List Alternative**: C# `List<T>` initialization `new List<int> {1,2,3}` is closer to Go slice behavior (dynamic)
- **Brevity**: Go: `[]int{1,2,3}`; C# array: `new int[] {1,2,3}` or `{1,2,3}` (context-dependent)
- **Use Case Alignment**: Go slice literals are the standard way to create collections inline. **C# equivalent depends on need**:
  - Use `new int[] {...}` for fixed-size (but this is an array, not a slice)
  - Use `new List<int> {...}` for dynamic collections (closest to Go slices)
  - Use `.AsSpan()` if you need a view (different use case)

**Different defaults**:
- **Go**: Slice literals are concise and create the primary collection type (dynamic, growable)
- **C#**: Array literals create fixed arrays; must use `List<T>` for dynamic behavior

**Key insight**: When you write `[]int{1,2,3}` in Go, you get a dynamic, growable slice. When you write `{1,2,3}` in C#, you get a fixed-size array. For equivalent functionality, C# developers use `new List<int> {1,2,3}`, which is more verbose.

The **Go approach is simpler** — one literal syntax for the common case. C# separates fixed (arrays) from dynamic (lists), requiring explicit choice.
