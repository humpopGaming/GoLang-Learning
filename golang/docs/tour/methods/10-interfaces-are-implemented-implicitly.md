# Interfaces Are Implemented Implicitly

## Go Concept

A type implements an interface by implementing its methods. There is **no explicit declaration** of intent, no "implements" keyword. Implicit interfaces decouple the definition of an interface from its implementation, which allows interfaces to appear in any package without prearrangement.

This is a powerful feature that enables defining interfaces in consumer code rather than provider code, leading to more flexible and maintainable designs.

### Go Example

```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M() // Output: hello
}
```

## C# Equivalent

C# requires **explicit interface implementation** using the `:` syntax in the type declaration. You must explicitly state that a type implements an interface. This makes the relationship clear but requires the type to know about the interface at definition time.

C# also supports **explicit interface implementation** where methods are only accessible through the interface type, providing more control over the public API.

### C# Example

```csharp
using System;

public interface I
{
    void M();
}

// Explicit declaration that T implements I
public class T : I
{
    public string S { get; set; }
    
    // Must implement the interface method
    public void M()
    {
        Console.WriteLine(S);
    }
}

// Explicit interface implementation (method only accessible via interface)
public class TExplicit : I
{
    public string S { get; set; }
    
    // Only accessible when cast to I
    void I.M()
    {
        Console.WriteLine(S);
    }
}

class Program
{
    static void Main()
    {
        I i = new T { S = "hello" };
        i.M(); // Output: hello
        
        var t = new T { S = "world" };
        t.M(); // Can also call directly
        
        var te = new TExplicit { S = "explicit" };
        // te.M(); // Compile error: M is not accessible
        ((I)te).M(); // Must cast to interface
    }
}
```

## Key Differences

- **Declaration Required**: Go has no explicit declaration; C# requires `: IInterface` in type declaration
- **Decoupling**: Go interfaces can be defined after types exist; C# requires types to know interfaces upfront
- **Package Independence**: Go interfaces can be in any package; C# interfaces must be known to implementing types
- **Consumer-Driven Design**: Go enables interfaces defined by consumers; C# requires provider awareness
- **Refactoring**: Go allows adding interfaces without changing types; C# requires modifying type declarations
- **Explicit Implementation**: C# supports hiding methods behind interfaces; Go has no equivalent
- **Compile-Time Checking**: C# verifies implementation at type definition; Go checks at assignment
- **Visibility**: C# makes interface relationships explicit in code; Go's are discovered through usage
- **Flexibility vs Safety**: Go prioritizes flexibility; C# prioritizes explicit contracts
