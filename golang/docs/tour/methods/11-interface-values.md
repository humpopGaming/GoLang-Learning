# Interface Values

## Go Concept

Under the hood, interface values can be thought of as a tuple of a value and a concrete type: `(value, type)`. An interface value holds a value of a specific concrete type. Calling a method on an interface value executes the method of the same name on its underlying concrete type.

Understanding interface values is key to understanding how Go's polymorphism works and how type information is preserved at runtime.

### Go Example

```go
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i) // (&{Hello}, *main.T)
	i.M()       // Hello

	i = F(math.Pi)
	describe(i) // (3.141592653589793, main.F)
	i.M()       // 3.141592653589793
}
```

## C# Equivalent

C# handles interface values through **boxing** and **unboxing** for value types, and through regular reference semantics for reference types. When you assign a struct to an interface variable, the struct is boxed into a heap-allocated object. Reference types don't require boxing.

C# interface variables store a reference to an object that implements the interface. The type information is preserved through the object's virtual method table (vtable).

### C# Example

```csharp
using System;

public interface I
{
    void M();
}

public class T : I
{
    public string S { get; set; }
    
    public void M()
    {
        Console.WriteLine(S);
    }
}

public struct F : I
{
    private readonly double value;
    
    public F(double value)
    {
        this.value = value;
    }
    
    public void M()
    {
        Console.WriteLine(value);
    }
}

class Program
{
    static void Describe(I i)
    {
        Console.WriteLine($"({i}, {i.GetType()})");
    }
    
    static void Main()
    {
        I i;
        
        // Reference type - no boxing needed
        i = new T { S = "Hello" };
        Describe(i); // (T, T)
        i.M();       // Hello
        
        // Value type - boxing occurs
        i = new F(Math.PI);
        Describe(i); // (F, F)
        i.M();       // 3.14159265358979
    }
}
```

## Key Differences

- **Internal Representation**: Go uses `(value, type)` tuple; C# uses object references with type metadata
- **Boxing**: C# boxes value types when assigned to interfaces; Go's interface values don't box in the same sense
- **Value Semantics**: Go can store actual values in interfaces; C# structs are boxed (copied to heap)
- **Pointer Support**: Go interfaces can hold pointers or values; C# interfaces hold references (boxed for value types)
- **Type Information**: Both preserve runtime type information; C# uses vtables, Go uses interface tables (itabs)
- **Performance**: Go's approach can be more efficient for small values; C# boxing adds allocation overhead
- **Nil Handling**: Go distinguishes nil interface from interface with nil value; C# has simpler null semantics
- **Memory**: Go interface values are typically two words (pointer + type); C# references are one word (plus boxed object for structs)
