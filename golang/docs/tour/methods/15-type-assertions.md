# Type Assertions

## Go Concept

A **type assertion** provides access to an interface value's underlying concrete value. The syntax `t := i.(T)` asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.

If `i` does not hold a `T`, the statement will trigger a panic. To test whether an interface value holds a specific type, use the two-value form: `t, ok := i.(T)`. If `i` holds a `T`, then `t` will be the underlying value and `ok` will be true. Otherwise, `ok` will be false and `t` will be the zero value of type `T`, and no panic occurs.

### Go Example

```go
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// Type assertion with panic risk
	s := i.(string)
	fmt.Println(s) // hello

	// Safe type assertion
	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	// Failed type assertion (safe form)
	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	// This would panic: f = i.(float64)
	
	// Practical usage
	var val any = 42
	if num, ok := val.(int); ok {
		fmt.Printf("It's an int: %d\n", num)
	}
}
```

## C# Equivalent

C# provides multiple ways to check and convert types:
1. **Cast operator** `(T)`: Traditional casting that throws `InvalidCastException` on failure
2. **as operator**: Safe casting that returns null for reference types on failure
3. **is operator**: Pattern matching for type checking
4. **is with declaration**: Combined type check and cast (C# 7.0+)

C# distinguishes between value types and reference types, affecting how type checking works. The `as` operator only works with reference types and nullable value types, while `is` works with all types.

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        object i = "hello";
        
        // Traditional cast (throws on failure)
        string s = (string)i;
        Console.WriteLine(s); // hello
        
        // Safe cast with 'as' (reference types only)
        string s2 = i as string;
        Console.WriteLine(s2 != null); // True
        
        // Failed safe cast
        double? d = i as double?; // null for reference type
        Console.WriteLine(d == null); // True
        
        // Pattern matching with 'is' (C# 7.0+)
        if (i is string str)
        {
            Console.WriteLine($"It's a string: {str}");
        }
        
        // Type check only
        if (i is string)
        {
            Console.WriteLine("It's a string");
        }
        
        // Value type example
        object val = 42;
        
        // Must use cast or is pattern for value types
        if (val is int num)
        {
            Console.WriteLine($"It's an int: {num}");
        }
        
        // This throws InvalidCastException
        // double d2 = (double)val;
        
        // Unboxing requires exact type match
        int intVal = (int)val;
        Console.WriteLine(intVal); // 42
    }
}
```

## Key Differences

- **Syntax**: Go uses `v.(T)` syntax; C# uses `(T)v`, `v as T`, or `v is T pattern`
- **Safety Options**: Go provides optional safety with two-value return; C# has separate safe (`as`, `is`) and unsafe (cast) operators
- **Failure Behavior**: Go panics or returns false; C# throws exception or returns null
- **Pattern Matching**: C# `is` with declaration combines check and cast in one statement; Go requires separate steps
- **Value Types**: C# `as` doesn't work with value types (requires nullable); Go type assertions work uniformly
- **Type Hierarchy**: C# considers inheritance (can cast to base types); Go only checks concrete type
- **Common Pattern**: Go uses `if val, ok := i.(Type); ok {}` idiom; C# uses `if (i is Type val) {}`
- **Performance**: Go type assertions are very fast; C# casting can involve vtable lookups
- **Unboxing**: C# requires exact type match when unboxing value types; Go doesn't have boxing
- **Readability**: C# pattern matching is more concise; Go is more explicit about the check
