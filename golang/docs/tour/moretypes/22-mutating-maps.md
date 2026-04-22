# Mutating Maps

## Go Concept

Map operations in Go:

- **Insert or update**: `m[key] = elem`
- **Retrieve**: `elem = m[key]`
- **Delete**: `delete(m, key)`
- **Test if key exists**: `elem, ok = m[key]`
  - If `key` is in `m`, `ok` is `true`
  - If not, `ok` is `false` and `elem` is the zero value for the map's element type
  - If `elem` or `ok` have not been declared, use short declaration: `elem, ok := m[key]`

### Go Example

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)

	// Insert
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// Update
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// Delete
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])  // 0 (zero value)

	// Test for presence
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// Common pattern: check before using
	if val, exists := m["Answer"]; exists {
		fmt.Println("Found:", val)
	} else {
		fmt.Println("Not found")
	}
}
```

## C# Equivalent

C# **Dictionary** provides similar operations:

- **Insert or update**: `dict[key] = value` or `dict.Add(key, value)` (throws if exists)
- **Retrieve**: `value = dict[key]` (throws if not found)
- **Delete**: `dict.Remove(key)` (returns bool)
- **Test if key exists**: `dict.TryGetValue(key, out value)` or `dict.ContainsKey(key)`
- **Safe operations**: Use `TryGetValue` to avoid exceptions

### C# Example

```csharp
using System;
using System.Collections.Generic;

class Program
{
    static void Main()
    {
        var m = new Dictionary<string, int>();

        // Insert (two ways)
        m["Answer"] = 42;
        // m.Add("Answer", 42);  // Alternative: throws if key exists
        Console.WriteLine($"The value: {m["Answer"]}");

        // Update
        m["Answer"] = 48;
        Console.WriteLine($"The value: {m["Answer"]}");

        // Delete
        m.Remove("Answer");
        // Accessing non-existent key throws KeyNotFoundException
        // Console.WriteLine($"The value: {m["Answer"]}");  // Would throw!

        // Test for presence - TryGetValue (recommended)
        if (m.TryGetValue("Answer", out int value))
        {
            Console.WriteLine($"The value: {value} Present? True");
        }
        else
        {
            Console.WriteLine($"The value: 0 Present? False");
        }

        // Alternative: ContainsKey (less efficient for retrieving)
        if (m.ContainsKey("Answer"))
        {
            Console.WriteLine($"Found: {m["Answer"]}");
        }
        else
        {
            Console.WriteLine("Not found");
        }

        // C# 7+ pattern: TryGetValue with pattern matching
        if (m.TryGetValue("Answer", out var v))
        {
            Console.WriteLine($"Found: {v}");
        }
        else
        {
            Console.WriteLine("Not found");
        }
    }
}
```

## Key Differences

- **Access Safety**: Go returns zero value for missing keys; C# throws `KeyNotFoundException`
- **Existence Check**: Go uses comma-ok idiom `v, ok := m[key]`; C# uses `TryGetValue(key, out value)` or `ContainsKey(key)`
- **Delete Operation**: Go uses built-in `delete(m, key)` function; C# uses `m.Remove(key)` method
- **Delete Return**: Go's `delete` returns nothing; C#'s `Remove` returns `bool` indicating if key existed
- **Insert vs Update**: Go has one syntax for both; C# has `Add` (insert only, throws if exists) and indexer (insert or update)
- **Idiomatic Patterns**: Go's comma-ok is built into the language; C# uses out parameters and method calls
- **Performance**: Go's comma-ok does one lookup; C# `ContainsKey` + indexer does two (use `TryGetValue` for one lookup)
- **Nil Map**: Go allows reading from nil maps (returns zero value); C# null dictionary throws `NullReferenceException`

**Different philosophies**:
- Go: Safe by default (missing keys return zero values), comma-ok idiom for explicit checks
- C#: Exceptions for missing keys, explicit methods for safe operations

**Best practices alignment**:
- **Go**: Always use `v, ok := m[key]` when key might not exist
- **C#**: Always use `TryGetValue` when key might not exist (avoids exceptions)

Both languages encourage checking existence before access, but use different mechanisms. Go's comma-ok idiom is more concise and integrated, while C#'s approach is more explicit with method calls. The functional outcomes are equivalent.
