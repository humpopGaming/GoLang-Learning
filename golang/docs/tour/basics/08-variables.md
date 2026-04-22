# Variables

## Go Concept

The `var` statement declares one or more variables. The type is specified last (after the variable name). Variables can be declared at package or function level.

A `var` declaration can include initializers, one per variable. If an initializer is present, the type can be omitted and the variable will take the type of the initializer.

### Go Example

```go
package main

import "fmt"

// Package-level variables
var c, python, java bool

func main() {
	// Function-level variables
	var i int
	fmt.Println(i, c, python, java)
}
```

## C# Equivalent

C# also uses variable declarations with type-first syntax. Variables can be:

- **Class/struct fields**: Similar to Go's package-level variables
- **Local variables**: Inside methods, similar to Go's function-level variables
- **Properties**: A C# concept with no direct Go equivalent

The key syntax difference is that C# puts the type first, while Go puts it after the variable name.

### C# Example

```csharp
using System;

class Program
{
    // Class-level fields (similar to Go package-level vars)
    static bool c, python, java;

    static void Main()
    {
        // Local variable
        int i;
        Console.WriteLine($"{i} {c} {python} {java}");
    }
}

// Or with properties (C# concept):
class Config
{
    public bool C { get; set; }
    public bool Python { get; set; }
    public bool Java { get; set; }
}
```

## Key Differences

- **Syntax Order**: Go uses `var name type`; C# uses `type name`
- **Default Values**: Both initialize to zero values (0, false, null), but Go is more explicit about this
- **Declaration Style**: Go's `var` can declare multiple variables of the same type; C# can too but syntax differs slightly
- **Properties vs Fields**: C# has properties (with getters/setters); Go only has exported/unexported fields
- **Scope**: Go has package-level variables; C# has class-level fields and static members
- **Use Case Alignment**: Both provide variable declarations with types. The main difference is syntax (type position) and Go's simpler model without properties. Both have similar scoping rules (package/class level vs function/method level) and both initialize variables to sensible zero values by default.
