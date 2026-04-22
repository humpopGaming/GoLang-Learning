# Function Closures

## Go Concept

Go functions may be **closures**. A closure is a function value that **references variables from outside its body**. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

- Closures capture variables from their surrounding scope
- The captured variables are shared between the closure and its enclosing function
- Each closure maintains its own captured state
- Closures enable data encapsulation and factory patterns

### Go Example

```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x  // Captures and modifies 'sum'
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()  // Two independent closures
	
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),    // Each closure has its own 'sum'
			neg(-2*i),
		)
	}
}
```

Output:
```
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90
```

**Key concepts**:
- `sum` is captured by the returned function
- Each call to `adder()` creates a new closure with its own `sum`
- The closure maintains state between calls

## C# Equivalent

C# **lambda expressions capture variables** from their surrounding scope, creating closures:

- Lambda expressions automatically capture variables from enclosing scope
- Captured variables are shared (same reference) with enclosing method
- Each closure instance maintains its own captured state
- Both value and reference types can be captured

### C# Example

```csharp
using System;

class Program
{
    static Func<int, int> Adder()
    {
        int sum = 0;
        return x =>
        {
            sum += x;  // Captures and modifies 'sum'
            return sum;
        };
    }

    static void Main()
    {
        var pos = Adder();  // Two independent closures
        var neg = Adder();
        
        for (int i = 0; i < 10; i++)
        {
            Console.WriteLine($"{pos(i)} {neg(-2 * i)}");
        }

        // Demonstrating closure behavior
        Console.WriteLine("\nClosure captures by reference:");
        int counter = 0;
        Action increment = () => counter++;
        Action decrement = () => counter--;
        
        increment();
        increment();
        Console.WriteLine(counter);  // 2
        decrement();
        Console.WriteLine(counter);  // 1

        // Multiple closures sharing the same variable
        Console.WriteLine("\nShared variable capture:");
        int shared = 0;
        var inc1 = new Func<int>(() => ++shared);
        var inc2 = new Func<int>(() => ++shared);
        
        Console.WriteLine(inc1());  // 1
        Console.WriteLine(inc2());  // 2
        Console.WriteLine(inc1());  // 3
    }
}
```

Output:
```
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90

Closure captures by reference:
2
1

Shared variable capture:
1
2
3
```

## Key Differences

- **Syntax**: Go uses `func` keyword; C# uses `=>` lambda operator
- **Capture Semantics**: Both capture by reference (can modify original variables)
- **Type Declaration**: Go infers return function type; C# uses `Func<>` or `Action<>`
- **Scope Rules**: Both capture variables from enclosing scope with same semantics
- **Multiple Closures**: Both allow multiple independent closures from same factory function
- **Loop Variable Capture**: Both have similar gotchas with loop variables (C# fixed in C# 5)
- **Memory Management**: Both keep captured variables alive (garbage collected)
- **Performance**: Both have similar overhead (closure objects allocated on heap)

**Different philosophies**:
- Go: Explicit `func` keyword, straightforward closure syntax
- C#: Concise lambda `=>` syntax, integrated with delegate types

**Common patterns**:

Go factory pattern:
```go
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}
```

C# factory pattern:
```csharp
Func<int, int> MakeMultiplier(int factor)
{
    return x => x * factor;
}
```

**Important gotcha (both languages)**:

Loop variable capture problem:
```go
// Go - WRONG
funcs := []func(){}
for i := 0; i < 3; i++ {
    funcs = append(funcs, func() {
        fmt.Println(i)  // Captures the loop variable!
    })
}
// All print 3!

// Go - CORRECT
for i := 0; i < 3; i++ {
    i := i  // Create new variable for each iteration
    funcs = append(funcs, func() {
        fmt.Println(i)
    })
}
```

```csharp
// C# - Fixed in C# 5+, but similar issue with foreach
List<Action> funcs = new List<Action>();
for (int i = 0; i < 3; i++)
{
    int localI = i;  // Capture local copy
    funcs.Add(() => Console.WriteLine(localI));
}
```

**Use cases**:
- **Event handlers**: Capture context for callbacks
- **Factory functions**: Create customized functions
- **Iterators**: Maintain state across calls
- **Encapsulation**: Hide implementation details
- **Partial application**: Bind some arguments, return function for remaining ones

Both languages handle closures similarly with nearly identical semantics. The main difference is syntax — Go's explicit `func` keyword vs C#'s concise `=>` operator. The underlying behavior, use cases, and gotchas are remarkably similar.
