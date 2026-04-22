# Exercise: Maps

## Go Concept

**Exercise**: Implement `WordCount`. It should return a map of the counts of each "word" in the string `s`. The `wc.Test` function runs a test suite against the provided function and prints success or failure.

This exercise tests your understanding of:
- Creating and manipulating maps
- Iterating over strings
- Splitting strings into words
- Updating map values

You might find `strings.Fields` helpful.

### Go Example

```go
package main

import (
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)
	
	for _, word := range words {
		wordMap[word]++
	}
	
	return wordMap
}

func main() {
	test := "I am learning Go!"
	result := WordCount(test)
	fmt.Println(result)
	// Output: map[I:1 am:1 learning:1 Go!:1]
	
	test2 := "the quick brown fox jumped over the lazy dog"
	result2 := WordCount(test2)
	fmt.Println(result2)
	// Output: map[the:2 quick:1 brown:1 fox:1 jumped:1 over:1 lazy:1 dog:1]
}
```

**Key techniques demonstrated**:
1. Create a map with `make(map[string]int)`
2. Use `strings.Fields(s)` to split by whitespace
3. Iterate with `range` over the slice of words
4. Increment counts using `wordMap[word]++` (zero value is 0 for int)
5. Return the populated map

## C# Equivalent

C# implementation using **Dictionary** and LINQ:

This exercise demonstrates:
- Dictionary creation and manipulation
- String splitting
- LINQ grouping and counting
- Both imperative and functional approaches

### C# Example

```csharp
using System;
using System.Collections.Generic;
using System.Linq;

class Program
{
    // Imperative approach (similar to Go)
    static Dictionary<string, int> WordCount(string s)
    {
        var wordMap = new Dictionary<string, int>();
        var words = s.Split(new[] { ' ', '\t', '\n' }, 
                           StringSplitOptions.RemoveEmptyEntries);
        
        foreach (var word in words)
        {
            if (wordMap.ContainsKey(word))
                wordMap[word]++;
            else
                wordMap[word] = 1;
        }
        
        return wordMap;
    }

    // More idiomatic C#: using TryGetValue
    static Dictionary<string, int> WordCountIdiomatic(string s)
    {
        var wordMap = new Dictionary<string, int>();
        var words = s.Split(new[] { ' ', '\t', '\n' }, 
                           StringSplitOptions.RemoveEmptyEntries);
        
        foreach (var word in words)
        {
            if (wordMap.TryGetValue(word, out int count))
                wordMap[word] = count + 1;
            else
                wordMap[word] = 1;
        }
        
        return wordMap;
    }

    // LINQ functional approach
    static Dictionary<string, int> WordCountLinq(string s)
    {
        return s.Split(new[] { ' ', '\t', '\n' }, 
                      StringSplitOptions.RemoveEmptyEntries)
                .GroupBy(word => word)
                .ToDictionary(group => group.Key, group => group.Count());
    }

    static void Main()
    {
        string test = "I am learning Go!";
        var result = WordCount(test);
        
        foreach (var kvp in result)
        {
            Console.WriteLine($"{kvp.Key}: {kvp.Value}");
        }
        // Output: I: 1, am: 1, learning: 1, Go!: 1

        Console.WriteLine();

        string test2 = "the quick brown fox jumped over the lazy dog";
        var result2 = WordCountLinq(test2);
        
        foreach (var kvp in result2)
        {
            Console.WriteLine($"{kvp.Key}: {kvp.Value}");
        }
        // Output includes: the: 2, quick: 1, etc.
    }
}
```

## Key Differences

- **Zero Value**: Go's int zero value (0) allows direct increment `m[key]++`; C# requires checking existence first
- **Increment Pattern**: Go: `m[key]++` always works; C# needs `ContainsKey` check or `TryGetValue`
- **String Splitting**: Go's `strings.Fields` auto-handles whitespace; C# `Split` requires explicit separators
- **Functional Style**: Go is imperative; C# offers LINQ functional approach with `GroupBy` and `ToDictionary`
- **Idiomatic Check**: Go's comma-ok simple; C# has `TryGetValue` pattern for efficiency
- **Iteration**: Go's `range` is concise; C# uses `foreach` over `KeyValuePair`
- **Code Length**: Go's zero value makes it more concise; C# needs explicit existence checks (imperative) or LINQ (functional)

**Different philosophies**:
- Go: Simple, direct, zero values eliminate null checks
- C#: Multiple approaches (imperative vs LINQ), explicit safety checks

**Why C# is more verbose**:
- No zero value concept for reference types in dictionaries
- Must check existence before incrementing
- More defensive programming required

**C# advantages**:
- LINQ provides elegant functional solution
- `TryGetValue` is very efficient (single lookup)
- More explicit about what's happening

Both solve the problem effectively, but Go's zero value semantics make the imperative approach cleaner. C#'s LINQ provides an alternative that's arguably more readable once you understand the functional approach.
