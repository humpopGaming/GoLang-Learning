# Nil Slices

## Go Concept

The **zero value of a slice is nil**. A nil slice has:
- length = 0
- capacity = 0
- no underlying array

### Go Example

```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	
	if s == nil {
		fmt.Println("nil!")
	}

	// Nil slice can be used with append
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
```

Output:
```
[] 0 0
nil!
[1] 1 1
```

## C# Equivalent

C# has different concepts for "nothing":

- **null**: Reference is not pointing to any object
- **Empty collections**: Collection exists but has zero elements
- **Default values**: What you get with `default(T)` or uninitialized fields

For collections:
- **Array**: `null` vs `new int[0]` (empty but not null)
- **List<T>**: `null` vs `new List<int>()` (empty but not null)
- **Span<T>**: `default(Span<T>)` is empty (but Span is a value type, not null)

### C# Example

```csharp
using System;
using System.Collections.Generic;

class Program
{
    static void Main()
    {
        // Array: null vs empty
        int[] arr = null;
        Console.WriteLine($"null array: {arr == null}");
        // Console.WriteLine(arr.Length);  // NullReferenceException!

        arr = new int[0];  // Empty array (not null)
        Console.WriteLine($"empty array: Length={arr.Length}, IsNull={arr == null}");

        // List: null vs empty
        List<int> list = null;
        Console.WriteLine($"\nnull list: {list == null}");

        list = new List<int>();  // Empty list (not null)
        Console.WriteLine($"empty list: Count={list.Count}, IsNull={list == null}");

        // Can add to empty list (not null)
        list.Add(1);
        Console.WriteLine($"after Add: Count={list.Count}, [{string.Join(", ", list)}]");

        // Span: default is empty (value type, can't be null)
        Span<int> span = default;
        Console.WriteLine($"\ndefault span: Length={span.Length}, IsEmpty={span.IsEmpty}");
    }
}
```

Output:
```
null array: True
empty array: Length=0, IsNull=False

null list: True
empty list: Count=0, IsNull=False
after Add: Count=1, [1]

default span: Length=0, IsEmpty=True
```

## Key Differences

- **Nil vs Null**: Go's `nil` is the zero value for slices; C#'s `null` means no object exists (different concept)
- **Zero Value**: Go nil slice is ready to use; C# null collection throws NullReferenceException if accessed
- **Empty vs Nil**: Go distinguishes `nil` slice from empty slice `[]int{}`; C# distinguishes `null` from `new int[0]`
- **Append to Nil**: Go allows `append(nil, value)` — works fine; C# can't call methods on null references
- **Value vs Reference**: Go slices are descriptors (can be nil); C# Span<T> is value type (can't be null, only empty)
- **Use Case Alignment**:

**Go nil slice behavior**:
```go
var s []int        // nil slice
s = append(s, 1)   // Works! Creates underlying array
```

**C# equivalent attempts**:
```csharp
List<int> list = null;
list.Add(1);           // Throws NullReferenceException!

// Must initialize:
list = new List<int>();
list.Add(1);           // Now works
```

**Different guarantees**:
- **Go**: nil slices are usable — you can pass them to functions, append to them, range over them
- **C#**: null references are not usable — accessing throws exception, must check or initialize

**Best practices comparison**:
- **Go**: `var s []int` creates usable nil slice — common pattern
- **C#**: Always initialize collections to avoid null checks:
  ```csharp
  List<int> list = new List<int>();  // Never null
  // Or use null-coalescing:
  list = list ?? new List<int>();
  ```

**Key insight**: Go's nil slices are **safe and functional** — they work with all slice operations. C#'s null collections are **dangerous** — they cause runtime errors. This leads to different patterns:

- **Go**: Nil slices are expected and handled naturally
- **C#**: Avoid null collections; initialize to empty instead

The Go approach is simpler — one zero value that works everywhere. C# requires distinguishing null (error-prone) from empty (safe), leading to more defensive programming and initialization boilerplate.
