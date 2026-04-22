# Slices Are Like References to Arrays

## Go Concept

A **slice does not store any data** — it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array. Other slices that share the same underlying array will see those changes.

### Go Example

```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```

Output:
```
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]
```

## C# Equivalent

C# has similar behavior with **Span<T>** and **ArraySegment<T>** — they are views into arrays, not copies.

Modifying elements through a span affects the underlying array. Multiple spans over the same array will see each other's changes.

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        string[] names = { "John", "Paul", "George", "Ringo" };
        Console.WriteLine(string.Join(" ", names));

        // Span<T> creates views into the array
        Span<string> a = names.AsSpan(0, 2);
        Span<string> b = names.AsSpan(1, 3);
        
        Console.WriteLine($"[{string.Join(" ", a.ToArray())}] [{string.Join(" ", b.ToArray())}]");

        // Modifying b affects the underlying array
        b[0] = "XXX";
        
        Console.WriteLine($"[{string.Join(" ", a.ToArray())}] [{string.Join(" ", b.ToArray())}]");
        Console.WriteLine(string.Join(" ", names));
    }
}
```

Output:
```
John Paul George Ringo
[John Paul] [Paul George]
[John XXX] [XXX George]
John XXX George Ringo
```

## Key Differences

- **Reference Semantics**: Both Go slices and C# spans are references/views — they don't copy data
- **Syntax**: Go has built-in slice syntax; C# requires `AsSpan()` method calls
- **Shared Modifications**: Both see changes to the underlying array when multiple views overlap
- **Default Behavior**: Go slices are always references; C# arrays can be copied with `.ToArray()` or used as views with `Span<T>`
- **Common Pattern**: Go developers use slices constantly; C# developers use `Span<T>` for performance scenarios but often copy arrays
- **Memory Efficiency**: Both avoid copying data when slicing/spanning
- **Use Case Alignment**: Go slices and C# `Span<T>` both provide efficient views that reference underlying data. **Key difference**: Go slices can grow with `append()` (creating new array if needed); **C# spans are fixed-size views**.

**Philosophy difference**:
- **Go**: Slices as references is the default, expected behavior
- **C#**: Spans are opt-in for performance; default pattern is copying arrays or using `List<T>`

The reference behavior is identical when used, but **Go makes this the primary pattern** while **C# treats it as an optimization technique**. Most C# code would use `List<T>` or copy arrays rather than manage span references.
