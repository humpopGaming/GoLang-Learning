# Range Continued

## Go Concept

You can **skip the index or value** by assigning to `_` (blank identifier).

```go
for i, _ := range pow    // skip value
for _, value := range pow // skip index
```

If you only want the index, you can omit the second variable entirely:
```go
for i := range pow
```

### Go Example

```go
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	
	// Only want index - omit value
	for i := range pow {
		pow[i] = 1 << uint(i)  // == 2**i
	}
	
	// Only want value - discard index with _
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```

Output:
```
1
2
4
8
16
32
64
128
256
512
```

## C# Equivalent

C# doesn't have a direct blank identifier like Go's `_`. Instead, you:

- **Discard value**: Use discard pattern `_` (C# 7.0+) or just don't declare variable
- **Skip index**: Use `foreach` (no index available by default)
- **Need index only**: Use `for` loop and don't use array value

C# 7.0+ introduced **discards** (`_`) similar to Go's blank identifier.

### C# Example

```csharp
using System;
using System.Linq;

class Program
{
    static void Main()
    {
        int[] pow = new int[10];

        // Only want index - traditional for loop
        for (int i = 0; i < pow.Length; i++)
        {
            pow[i] = 1 << i;  // == 2**i
        }

        // Only want value - foreach (no index available)
        foreach (int value in pow)
        {
            Console.WriteLine(value);
        }

        Console.WriteLine("\nWith discard patterns:");

        // Using LINQ Select with discard for index (C# 7.0+)
        pow.Select((value, _) => value)
           .ToList()
           .ForEach(v => Console.WriteLine(v));

        // Discard in deconstruction (C# 7.0+)
        var tuples = pow.Select((value, index) => (value, index));
        foreach (var (value, _) in tuples)
        {
            Console.WriteLine(value);
        }

        // If you really want index only (less common)
        for (int i = 0; i < pow.Length; i++)
        {
            // Use i, ignore pow[i]
            Console.WriteLine($"Index: {i}");
        }
    }
}
```

Output:
```
1
2
4
8
16
...
```

## Key Differences

- **Blank Identifier**: Go uses `_` as true blank identifier; C# `_` is a discard (similar but newer feature)
- **Index-only Syntax**: Go `for i := range s` is concise; C# requires `for (int i = 0; i < s.Length; i++)`
- **Value-only Syntax**: Go `for _, v := range s`; C# `foreach (var v in s)` (simpler, no explicit discard)
- **Language Age**: Go had `_` from start; C# added discards in C# 7.0 (2017)
- **Compiler Behavior**: Go `_` doesn't allocate/compute; C# `_` may still compute (depends on context)
- **Use Case Alignment**:

**Skip index (value only)**:
```go
// Go
for _, value := range items {
    fmt.Println(value)
}
```

```csharp
// C# - simpler, no discard needed
foreach (var value in items)
{
    Console.WriteLine(value);
}
```

**Skip value (index only)**:
```go
// Go - very concise
for i := range items {
    fmt.Println(i)
}
```

```csharp
// C# - more verbose
for (int i = 0; i < items.Length; i++)
{
    Console.WriteLine(i);
}
```

**Skip both (just iterate, count, etc.)**:
```go
// Go
count := 0
for range items {
    count++
}
```

```csharp
// C# - use Length or Count
int count = items.Length;  // Or items.Count() for IEnumerable
```

**Different philosophies**:
- **Go**: Uniform `range` syntax with explicit discards for unneeded values
- **C#**: Different constructs for different needs (foreach for values, for for indices)

**Go blank identifier (`_`) features**:
- Ignores assignment (compiler doesn't complain about unused variables)
- Can be used in multiple contexts (imports, function returns, range)
- Explicitly communicates intent to ignore

**C# discard (`_`) features** (C# 7.0+):
- Similar to Go's `_` in many contexts
- Can be used in pattern matching, deconstruction, out parameters
- Newer addition to language

**Before C# 7.0 discards**:
```csharp
// Had to declare and ignore variable
foreach (var item in items)
{
    int unused = item.Id;  // Compiler warning about unused
}

// Or use underscore as variable name (convention, not language feature)
foreach (var _ in items)
{
    // _variable name convention
}
```

**Key insight**: Go's range with blank identifier provides uniform syntax for all iteration patterns. C# optimizes for the common case (foreach for values) but becomes more verbose when you need indices.

**Common Go patterns**:
```go
// Initialize slice based on index only
for i := range data {
    data[i] = computeValue(i)
}

// Count elements (ignore values)
count := 0
for range items {
    count++
}

// Check if slice is empty (range on empty slice doesn't execute)
for range items {
    fmt.Println("Has items")
    break
}
```

**C# equivalents** (various patterns):
```csharp
// Initialize array based on index
for (int i = 0; i < data.Length; i++)
{
    data[i] = ComputeValue(i);
}

// Count elements
int count = items.Length;  // Direct property access

// Check if has items
if (items.Length > 0)
{
    Console.WriteLine("Has items");
}
```

The Go approach emphasizes consistency (one iteration construct with selective ignoring), while C# emphasizes directness (use appropriate property/loop for each need).
