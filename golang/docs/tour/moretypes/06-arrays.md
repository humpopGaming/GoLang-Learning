# Arrays

## Go Concept

An array in Go has a **fixed size** that's part of its type. The type `[n]T` is an array of `n` values of type `T`.

Arrays in Go are value types — they're copied when assigned. The size cannot change after declaration.

### Go Example

```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

## C# Equivalent

C# also has **arrays** with fixed size:

- Declaration: `int[] primes` or `int[6] primes`
- Arrays are reference types (unlike Go where they're value types)
- Size is fixed after creation
- Zero-indexed like Go

The key difference is that C# arrays are reference types (stored on heap), while Go arrays are value types (copied on assignment).

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        string[] a = new string[2];
        a[0] = "Hello";
        a[1] = "World";
        Console.WriteLine($"{a[0]} {a[1]}");
        Console.WriteLine($"[{string.Join(" ", a)}]");

        int[] primes = new int[] { 2, 3, 5, 7, 11, 13 };
        // or: int[] primes = { 2, 3, 5, 7, 11, 13 };
        Console.WriteLine($"[{string.Join(" ", primes)}]");
    }
}
```

## Key Differences

- **Value vs Reference**: Go arrays are value types (copied); C# arrays are reference types (referenced)
- **Type System**: Go array type includes size `[6]int`; C# array type doesn't `int[]`
- **Size in Type**: In Go, `[5]int` and `[6]int` are different types; in C# both are `int[]`
- **Heap vs Stack**: Go arrays can be stack or heap allocated; C# arrays are always heap-allocated
- **Common Usage**: Go developers prefer slices over arrays; C# developers use arrays or List<T>
- **Use Case Alignment**: Both provide fixed-size sequential collections. However, **Go arrays are rarely used directly** — Go developers use slices (covered next). C# arrays are more common, but `List<T>` is often preferred for dynamic collections.

**Important note**: In Go, you'll almost always use **slices** instead of arrays. Arrays are the underlying data structure, but slices provide a more flexible interface. This is different from C# where arrays are commonly used directly.

```go
// Go: Arrays rarely used directly
var arr [5]int        // Array (value type, size in type)
slice := arr[:]       // Slice (view of array, commonly used)
```

```csharp
// C#: Arrays commonly used
int[] arr = new int[5];  // Array (reference type)
List<int> list = new List<int>();  // List for dynamic size
```
