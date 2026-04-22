# Slice Length and Capacity

## Go Concept

A slice has both a **length** and a **capacity**.

- **Length**: number of elements in the slice — `len(s)`
- **Capacity**: number of elements in the underlying array, counting from the first element in the slice — `cap(s)`

You can extend a slice's length by re-slicing it, provided it has sufficient capacity.

### Go Example

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	// Extend its length
	s = s[:4]
	printSlice(s)

	// Drop its first two values
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

Output:
```
len=6 cap=6 [2 3 5 7 11 13]
len=0 cap=6 []
len=4 cap=6 [2 3 5 7]
len=2 cap=4 [5 7]
```

## C# Equivalent

C# arrays and lists have different length/capacity concepts:

- **Arrays**: Only `Length` property (fixed size, no capacity concept)
- **List<T>**: Has both `Count` (length) and `Capacity` (allocated space)
- **Span<T>**: Only `Length` property (view into existing memory)

**List<T>** is most similar to Go slices for dynamic behavior, but it's a full collection, not a view.

### C# Example

```csharp
using System;
using System.Collections.Generic;

class Program
{
    static void Main()
    {
        // List<T> has Count and Capacity
        List<int> list = new List<int> { 2, 3, 5, 7, 11, 13 };
        PrintList(list);

        // Remove all items (Count=0, Capacity unchanged)
        list.Clear();
        PrintList(list);

        // Can't "extend" like Go - must add elements
        list.AddRange(new[] { 2, 3, 5, 7 });
        PrintList(list);

        // Remove first two
        list.RemoveRange(0, 2);
        PrintList(list);

        Console.WriteLine("\nSpan example:");
        // Span only has Length
        int[] arr = { 2, 3, 5, 7, 11, 13 };
        Span<int> span = arr;
        PrintSpan(span);

        span = span.Slice(0, 4);  // Reduce length
        PrintSpan(span);

        span = span.Slice(2);  // Drop first two
        PrintSpan(span);
    }

    static void PrintList(List<int> list)
    {
        Console.WriteLine($"Count={list.Count} Capacity={list.Capacity} [{string.Join(", ", list)}]");
    }

    static void PrintSpan(Span<int> span)
    {
        Console.WriteLine($"Length={span.Length} [{string.Join(", ", span.ToArray())}]");
    }
}
```

Output:
```
Count=6 Capacity=6 [2, 3, 5, 7, 11, 13]
Count=0 Capacity=6 []
Count=4 Capacity=6 [2, 3, 5, 7]
Count=2 Capacity=6 [5, 7]

Span example:
Length=6 [2, 3, 5, 7, 11, 13]
Length=4 [2, 3, 5, 7]
Length=2 [5, 7]
```

## Key Differences

- **Length Function**: Go uses `len(s)` function; C# uses `.Length` or `.Count` property
- **Capacity Function**: Go uses `cap(s)` function; C# uses `.Capacity` property (List<T> only)
- **Capacity Concept**: Go capacity is about underlying array; C# capacity is about allocated space (different meaning)
- **Re-slicing**: Go can extend slice length via re-slicing if capacity allows; C# Span can only shrink, List<T> requires adding elements
- **Arrays vs Views**: C# arrays have no capacity concept; Span<T> (view) has no capacity; List<T> (collection) has capacity
- **Use Case Alignment**: 

**Go slice length/capacity model**:
- Length: current visible elements
- Capacity: available space in underlying array
- Can reveal hidden elements by re-slicing: `s = s[:cap(s)]`

**C# equivalent concepts**:
- **Span<T>**: Only length (it's a view, can't grow)
- **List<T>**: Count (used elements) and Capacity (allocated space, can grow automatically)
- **Arrays**: Only Length (fixed forever)

**Key insight**: Go slices let you "hide" array elements with slicing and reveal them later (if capacity allows). C# doesn't have this pattern:
- `Span<T>` can only shrink via `.Slice()`
- `List<T>` grows by adding elements, not by revealing hidden ones

**Different mental models**:
- **Go**: Slice is a window into an array; you can slide the window and change its size (up to capacity)
- **C#**: 
  - Span is a window (can only shrink)
  - List is a growing collection (adds/removes elements)
  - Arrays are fixed

The Go model is unique and powerful — you can efficiently manipulate window position and size. C# requires choosing between fixed views (Span) or dynamic collections (List), neither of which exactly matches Go's slice behavior.
