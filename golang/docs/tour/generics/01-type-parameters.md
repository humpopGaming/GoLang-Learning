# Type Parameters

## Go Concept

Go supports **generic functions** using type parameters. Type parameters allow you to write functions that work with multiple types while maintaining type safety. A type parameter is declared in square brackets before the function's regular parameters.

Type parameters use **constraints** to specify what operations are allowed on the generic type. The most common constraint is `any` (equivalent to `interface{}`), but you can create custom constraints using interfaces.

### Go Example

```go
package main

import "fmt"

// Index returns the index of x in s, or -1 if not found
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15)) // Output: 2

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello")) // Output: -1
}
```

## C# Equivalent

C# has had **generic methods** since C# 2.0. Generic methods use angle brackets `<>` to declare type parameters and support constraints using the `where` clause. The concepts are very similar between Go and C#.

Key differences:
- Go uses square brackets `[]` for type parameters; C# uses angle brackets `<>`
- Go's `comparable` constraint is similar to C#'s `where T : IEquatable<T>` or requiring value equality
- C# has more built-in constraint keywords (`class`, `struct`, `new()`, `unmanaged`)
- Go infers type arguments more often; C# sometimes requires explicit specification

### C# Example

```csharp
using System;
using System.Collections.Generic;

public class Program
{
    // Generic method with constraint
    public static int Index<T>(List<T> list, T value) where T : IEquatable<T>
    {
        for (int i = 0; i < list.Count; i++)
        {
            if (list[i].Equals(value))
            {
                return i;
            }
        }
        return -1;
    }

    public static void Main()
    {
        // Works with integers
        var ints = new List<int> { 10, 20, 15, -10 };
        Console.WriteLine(Index(ints, 15)); // Output: 2

        // Works with strings
        var strings = new List<string> { "foo", "bar", "baz" };
        Console.WriteLine(Index(strings, "hello")); // Output: -1
    }
}
```

## Key Comparison

| Feature | Go | C# |
|---------|----|----|
| Syntax | `func Name[T constraint]()` | `void Name<T>() where T : constraint` |
| Type inference | Excellent, usually automatic | Good, but sometimes needs explicit types |
| Constraint syntax | Interface-based, inline | `where` clause, keyword-based |
| Multiple constraints | `interface{ A; B }` | `where T : A, B` |
| Numeric generics | Recently added with new constraints | Limited without custom interfaces |
