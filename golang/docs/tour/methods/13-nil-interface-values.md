# Nil Interface Values

## Go Concept

A nil interface value holds neither value nor concrete type. Calling a method on a nil interface is a **runtime error** because there is no type inside the interface tuple to indicate which concrete method to call.

This is different from an interface holding a nil value: a nil interface has no type information at all, while an interface with a nil value still knows what type it should be.

### Go Example

```go
package main

import "fmt"

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	describe(i) // (<nil>, <nil>)
	
	// Uncommenting this would cause a runtime panic:
	// i.M()
	// panic: runtime error: invalid memory address or nil pointer dereference
	
	// This is different from an interface with a nil underlying value:
	var t *T
	i = t
	describe(i) // (<nil>, *main.T) - interface is NOT nil
	i.M()       // Works if M handles nil receiver
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil receiver>")
		return
	}
	fmt.Println(t.S)
}
```

## C# Equivalent

C# has a simpler model: interface variables can be null, and calling methods on null references throws a **NullReferenceException**. C# doesn't distinguish between "null interface" and "interface containing null value" in the same way Go does - if the interface variable is null, that's the end of the story.

C# 8.0+ introduced nullable reference types which provide compile-time null safety checking.

### C# Example

```csharp
using System;

#nullable enable

public interface I
{
    void M();
}

public class T : I
{
    public string S { get; set; }
    
    public void M()
    {
        Console.WriteLine(S ?? "<null>");
    }
}

class Program
{
    static void Describe(I? i)
    {
        if (i == null)
        {
            Console.WriteLine("(null, null)");
        }
        else
        {
            Console.WriteLine($"({i}, {i.GetType().Name})");
        }
    }
    
    static void Main()
    {
        I? i = null;
        Describe(i); // (null, null)
        
        // Uncommenting this would throw NullReferenceException
        // i.M();
        
        // C# approach: null-conditional operator
        i?.M(); // Safe - doesn't execute if i is null
        
        // With nullable reference types, compiler warns:
        // I nonNullable = null; // Warning: Converting null literal
        
        i = new T { S = "hello" };
        Describe(i); // (hello, T)
        i.M();       // hello
    }
}
```

## Key Differences

- **Nil Interface Semantics**: Go has nil interface (no type info); C# has null reference (no object)
- **Error Type**: Go causes runtime panic; C# throws NullReferenceException
- **Type Information**: Go nil interface has no type tuple; C# null reference is simply absent
- **Detection**: Go can check `i == nil`; C# checks for null reference the same way
- **Nullable Reference Types**: C# 8.0+ has compiler-enforced null safety; Go has runtime nil checks
- **Interface with Nil Value**: Go distinguishes this from nil interface; C# doesn't have this distinction
- **Safety**: Both are runtime errors, but C# has compiler warnings with nullable reference types
- **Common Practice**: Go checks for nil interface; C# uses null-conditional operators (`?.`, `??`)
- **Design Philosophy**: Go's model is more complex but more flexible; C#'s is simpler but less nuanced
