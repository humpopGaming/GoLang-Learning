# Appending to a Slice

## Go Concept

The **append** function adds elements to a slice. If the slice's capacity is too small, a new, larger array is allocated automatically.

`func append(s []T, vs ...T) []T`

The first parameter is a slice, followed by elements to append. The result is a new slice containing all elements (may reference new underlying array).

### Go Example

```go
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

Output:
```
len=0 cap=0 []
len=1 cap=1 [0]
len=2 cap=2 [0 1]
len=5 cap=6 [0 1 2 3 4]
```

## C# Equivalent

C# uses **List<T>.Add()** and **List<T>.AddRange()** for dynamic growth:

- `List<T>.Add(item)` — add single element
- `List<T>.AddRange(items)` — add multiple elements
- Arrays are fixed-size (no append)

List<T> automatically grows capacity when needed (typically doubles).

### C# Example

```csharp
using System;
using System.Collections.Generic;

class Program
{
    static void Main()
    {
        List<int> s = new List<int>();
        PrintList(s);

        // Add single element
        s.Add(0);
        PrintList(s);

        // Add another element
        s.Add(1);
        PrintList(s);

        // Add multiple elements at once
        s.AddRange(new[] { 2, 3, 4 });
        PrintList(s);

        // Alternative: Collection initializer for initial values
        List<int> s2 = new List<int> { 0, 1, 2, 3, 4 };
        PrintList(s2);
    }

    static void PrintList(List<int> s)
    {
        Console.WriteLine($"Count={s.Count} Capacity={s.Capacity} [{string.Join(", ", s)}]");
    }
}
```

Output:
```
Count=0 Capacity=0 []
Count=1 Capacity=4 [0]
Count=2 Capacity=4 [0, 1]
Count=5 Capacity=8 [0, 1, 2, 3, 4]
Count=5 Capacity=8 [0, 1, 2, 3, 4]
```

## Key Differences

- **Function vs Method**: Go uses `append()` function; C# uses `.Add()` method on List<T>
- **Return Value**: Go append returns new slice; C# Add mutates existing list (no return needed)
- **Multiple Elements**: Go `append(s, 1, 2, 3)`; C# `.AddRange(new[] {1, 2, 3})`
- **Nil Handling**: Go append works on nil slices; C# throws NullReferenceException on null lists
- **Capacity Growth**: Both automatically grow capacity (Go and C# use similar doubling strategies)
- **Assignment Required**: Go requires `s = append(s, x)` to capture new slice; C# mutates in place
- **Use Case Alignment**: Both provide dynamic growth for collections.

**Append patterns comparison**:
```go
// Go: functional style (returns new slice)
s := []int{}
s = append(s, 1)
s = append(s, 2, 3, 4)
```

```csharp
// C# imperative style (mutates list)
List<int> s = new List<int>();
s.Add(1);
s.AddRange(new[] { 2, 3, 4 });
```

**Different semantics**:
- **Go append**: 
  - Returns new slice (which may share underlying array)
  - Must reassign: `s = append(s, x)`
  - Works on nil slices
  - Can invalidate other slices sharing same array (if reallocation occurs)

- **C# Add**:
  - Mutates list in place
  - No reassignment: `list.Add(x)`
  - Throws on null
  - No sharing issues (List<T> owns its storage)

**Why Go requires reassignment**:
When capacity is exceeded, append allocates new array and returns new slice pointing to it. Old slice still points to old array.

```go
s1 := []int{1, 2}
s2 := s1
s1 = append(s1, 3)  // May reallocate
// s1 and s2 might now point to different arrays
```

C# List<T> doesn't have this issue — all references point to same list object.

**Capacity growth comparison**:
- **Go**: Capacity typically doubles (implementation-dependent)
- **C#**: Capacity doubles until 2048 elements, then grows by 25% (implementation-dependent)

**Key insight**: Go's append is functional (returns new value), while C# Add is imperative (mutates object). This reflects deeper design philosophies:

- **Go**: Value-oriented, explicit about potential reallocation
- **C#**: Object-oriented, mutations hidden behind method calls

**Common mistake for C# developers learning Go**:
```go
// WRONG: forgetting to reassign
s := []int{1, 2}
append(s, 3)  // Does nothing! Lost the returned slice

// CORRECT:
s = append(s, 3)  // Capture the returned slice
```

**For Go developers learning C#**:
```csharp
// WRONG: trying to reassign
List<int> list = new List<int> { 1, 2 };
list = list.Add(3);  // Compile error! Add returns void

// CORRECT:
list.Add(3);  // Mutates in place
```

The mental model shift is significant — Go's append is a function returning a slice, C#'s Add is a method mutating an object.
