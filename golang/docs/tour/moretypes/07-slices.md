# Slices

## Go Concept

A **slice** is a dynamically-sized, flexible view into the elements of an array. Slices are much more common than arrays in Go code.

The type `[]T` is a slice with elements of type `T`. A slice is formed by specifying two indices, low and high bound: `a[low : high]`. This selects a half-open range including the first element but excluding the last.

### Go Example

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)  // [3 5 7]
}
```

## C# Equivalent

C# doesn't have a built-in "slice" type like Go. The closest equivalents are:

1. **ArraySegment<T>**: A view into an array (closest to Go slices)
2. **Span<T>** and **Memory<T>** (modern C#): Efficient views over contiguous memory
3. **LINQ `Skip().Take()`**: Creates new collections (not a view)
4. **Arrays or List<T>**: Different concepts (full collections, not views)

**Span<T>** (C# 7.2+) is the closest equivalent — it's a view over contiguous memory like Go slices, but with some differences.

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        int[] primes = { 2, 3, 5, 7, 11, 13 };

        // Option 1: ArraySegment (older, less convenient)
        ArraySegment<int> segment = new ArraySegment<int>(primes, 1, 3);
        Console.WriteLine(string.Join(" ", segment));  // 3 5 7

        // Option 2: Span<T> (modern, most like Go slices)
        Span<int> s = primes.AsSpan(1, 3);
        Console.WriteLine(string.Join(" ", s.ToArray()));  // 3 5 7

        // Option 3: LINQ (creates new array, not a view)
        var s2 = primes.Skip(1).Take(3);
        Console.WriteLine(string.Join(" ", s2));  // 3 5 7

        // Modifying span affects original array
        s[0] = 100;
        Console.WriteLine(primes[1]);  // 100
    }
}
```

## Key Differences

- **Built-in vs Library**: Go slices are built-in language feature; C# requires `Span<T>` or `ArraySegment<T>`
- **Syntax**: Go has clean slice syntax `arr[1:4]`; C# requires method calls `arr.AsSpan(1, 3)`
- **Common Usage**: Go slices are ubiquitous; C# developers often use full arrays or `List<T>` instead
- **Dynamic Growth**: Go slices can grow with `append()`; C# `Span<T>` is fixed-size (use `List<T>` for growth)
- **Heap vs Stack**: Go slices can reference heap or stack; C# `Span<T>` can span stack/heap but has restrictions
- **Use Case Alignment**: Go slices provide efficient views into arrays and are the primary collection type. **C# doesn't have a direct equivalent** for everyday use. 

**Different approaches**:
- **Go**: Slices are the standard way to work with sequences (dynamic, efficient, views)
- **C#**: Multiple options depending on need:
  - `Span<T>`/`Memory<T>` for high-performance views (modern)
  - `ArraySegment<T>` for views (older)
  - `List<T>` for dynamic collections (most common)
  - Arrays for fixed-size collections

**Go slices are unique** — they combine efficient views with dynamic growth (via `append`). C# splits this into multiple types: `Span<T>` for views, `List<T>` for dynamic collections. This is a **fundamental difference** in how the languages handle sequential collections.

For most practical C# code, you'd use `List<T>` where Go uses slices, accepting that `List<T>` is a full collection rather than a view.
