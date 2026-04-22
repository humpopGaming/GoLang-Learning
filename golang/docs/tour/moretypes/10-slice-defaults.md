# Slice Defaults

## Go Concept

When slicing, you may **omit the high or low bounds** to use their defaults.

- Default low bound: `0`
- Default high bound: length of the sliced array/slice

These expressions are equivalent:
```go
a[0:10]
a[:10]
a[0:]
a[:]
```

### Go Example

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)  // [3 5 7]

	s = s[:2]       // equivalent to s[0:2]
	fmt.Println(s)  // [3 5]

	s = s[1:]       // equivalent to s[1:len(s)]
	fmt.Println(s)  // [5]

	// Back to full slice
	s = []int{2, 3, 5, 7, 11, 13}
	s = s[:]        // equivalent to s[0:len(s)] - full slice
	fmt.Println(s)  // [2 3 5 7 11 13]
}
```

## C# Equivalent

C# has **range operators** (`..`) introduced in C# 8.0 that provide similar default behavior:

- `array[..5]` — from start to index 5
- `array[2..]` — from index 2 to end
- `array[..]` — full range

These work with arrays, spans, strings, and other types supporting ranges.

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        int[] s = { 2, 3, 5, 7, 11, 13 };

        // Using Range operator (..)
        Span<int> slice = s.AsSpan()[1..4];  // [3 5 7]
        Console.WriteLine(string.Join(" ", slice.ToArray()));

        slice = slice[..2];  // equivalent to [0..2]
        Console.WriteLine(string.Join(" ", slice.ToArray()));  // [3 5]

        slice = slice[1..];  // from index 1 to end
        Console.WriteLine(string.Join(" ", slice.ToArray()));  // [5]

        // Back to full array
        s = new int[] { 2, 3, 5, 7, 11, 13 };
        slice = s.AsSpan()[..];  // full range
        Console.WriteLine(string.Join(" ", slice.ToArray()));  // [2 3 5 7 11 13]

        // Alternative: array slicing returns new arrays (C# 8.0+)
        int[] arr = { 2, 3, 5, 7, 11, 13 };
        int[] sub1 = arr[1..4];
        int[] sub2 = arr[..3];
        int[] sub3 = arr[2..];
        Console.WriteLine(string.Join(" ", sub1));  // 3 5 7
        Console.WriteLine(string.Join(" ", sub2));  // 2 3 5
        Console.WriteLine(string.Join(" ", sub3));  // 5 7 11 13
    }
}
```

## Key Differences

- **Syntax Similarity**: Go uses `[:]`, C# uses `[..]` — very similar concepts
- **Default Bounds**: Both allow omitting start or end for defaults (0 and length)
- **Range Operator**: C# `..` is more recent (C# 8.0+); Go slicing is original feature
- **Copy vs View**: C# array slicing `arr[1..4]` creates new array; `span[1..4]` creates view (Go slices are views)
- **Type System**: Go ranges work on arrays/slices; C# ranges work on arrays/spans/strings/custom types with indexers
- **Use Case Alignment**: Both provide convenient defaults for common slicing patterns.

**Syntax comparison**:
| Operation | Go | C# (Span/Array) |
|-----------|-------|-----------------|
| First n elements | `s[:n]` | `s[..n]` |
| From index i to end | `s[i:]` | `s[i..]` |
| Full range | `s[:]` | `s[..]` |
| Middle section | `s[i:j]` | `s[i..j]` |

**Different implementations**:
- **Go**: Built-in slice syntax from day one, always creates views
- **C#**: Range operator added in C# 8.0, behavior depends on type:
  - Arrays: creates new array (copy)
  - Span<T>: creates view (no copy)

**Key insight**: The syntax is remarkably similar, showing convergent evolution. However:
- **Go slicing** is built-in, ubiquitous, and always efficient (views)
- **C# ranges** are newer, require specific types, and may copy (arrays) or view (spans)

For everyday C# code, you'd use ranges with arrays (which copies) or explicitly use `Span<T>` for view behavior. Go's approach is simpler — slicing always creates efficient views.
