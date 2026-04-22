# Type Switches

## Go Concept

A **type switch** is a construct that permits several type assertions in series. It's like a regular switch statement, but the cases specify types (not values), and those values are compared against the type of the value held by the given interface value.

The syntax `switch v := i.(type)` declares a variable `v` that takes on the type and value of each case. The special keyword `type` is used instead of a concrete type. Type switches are a clean way to handle different types without multiple if statements.

### Go Example

```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case bool:
		if v {
			fmt.Println("true!")
		} else {
			fmt.Println("false!")
		}
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)       // Twice 21 is 42
	do("hello")  // "hello" is 5 bytes long
	do(true)     // true!
	do(3.14)     // I don't know about type float64!
}

// Multiple types in one case
func describe(i any) string {
	switch i.(type) {
	case int, int8, int16, int32, int64:
		return "integer type"
	case uint, uint8, uint16, uint32, uint64:
		return "unsigned integer type"
	case float32, float64:
		return "floating point type"
	case string:
		return "string type"
	case bool:
		return "boolean type"
	default:
		return "unknown type"
	}
}
```

## C# Equivalent

C# provides **switch expressions with type patterns** (C# 7.0+) and **pattern matching** (enhanced in C# 8.0+). Modern C# switch statements and expressions support type patterns, property patterns, positional patterns, and more complex pattern matching scenarios.

C# type switches are more powerful than Go's, supporting additional patterns like property matching, guards (when clauses), and recursive patterns. The C# switch expression (C# 8.0+) provides a more concise syntax.

### C# Example

```csharp
using System;

class Program
{
    // Traditional switch statement with type patterns (C# 7.0+)
    static void Do(object i)
    {
        switch (i)
        {
            case int v:
                Console.WriteLine($"Twice {v} is {v * 2}");
                break;
            case string v:
                Console.WriteLine($"\"{v}\" is {v.Length} bytes long");
                break;
            case bool v:
                Console.WriteLine(v ? "true!" : "false!");
                break;
            default:
                Console.WriteLine($"I don't know about type {i.GetType().Name}!");
                break;
        }
    }
    
    // Modern switch expression (C# 8.0+)
    static string Describe(object i) => i switch
    {
        int => "integer type",
        uint => "unsigned integer type",
        long => "long integer type",
        float => "float type",
        double => "double type",
        string => "string type",
        bool => "boolean type",
        _ => "unknown type"
    };
    
    // Advanced pattern matching with guards
    static string DescribeAdvanced(object i) => i switch
    {
        int n when n > 0 => "positive integer",
        int n when n < 0 => "negative integer",
        int => "zero",
        string s when s.Length > 10 => "long string",
        string s when s.Length > 0 => "short string",
        string => "empty string",
        _ => "other type"
    };
    
    // Property pattern matching (C# 8.0+)
    record Person(string Name, int Age);
    
    static string DescribePerson(object obj) => obj switch
    {
        Person { Age: > 18 } p => $"{p.Name} is an adult",
        Person { Age: <= 18 } p => $"{p.Name} is a minor",
        _ => "not a person"
    };
    
    static void Main()
    {
        Do(21);       // Twice 21 is 42
        Do("hello");  // "hello" is 5 bytes long
        Do(true);     // true!
        Do(3.14);     // I don't know about type Double!
        
        Console.WriteLine(Describe(42));        // integer type
        Console.WriteLine(DescribeAdvanced(5)); // positive integer
        
        var person = new Person("Alice", 25);
        Console.WriteLine(DescribePerson(person)); // Alice is an adult
    }
}
```

## Key Differences

- **Syntax**: Go uses `switch v := i.(type)`; C# uses `switch (i)` with type patterns or `i switch` expressions
- **Variable Scope**: Go's type switch variable changes type in each case; C# declares new variable per case
- **Multiple Types**: Go can list multiple types in one case; C# requires separate cases or or-patterns
- **Pattern Matching**: C# supports property patterns, guards, positional patterns; Go only supports type matching
- **Expression Form**: C# has switch expressions that return values; Go switch is always a statement
- **Guards**: C# supports `when` clauses for conditional matching; Go requires separate if inside case
- **Default Case**: Both support default case; C# uses `default` or `_` in expressions
- **Fall Through**: Go cases don't fall through by default; C# requires `break` in statement form (not expression)
- **Exhaustiveness**: C# warns on non-exhaustive switch expressions; Go doesn't require covering all types
- **Inheritance**: C# matches derived types; Go only matches concrete types
- **Performance**: Both compile to efficient jump tables when possible
- **Modern Idioms**: C# favors switch expressions and pattern matching; Go favors type switches with explicit cases
