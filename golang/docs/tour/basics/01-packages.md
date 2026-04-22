# Packages

## Go Concept

Every Go program is made up of packages. Programs start running in package `main`. The package declaration at the top of a file determines which package the code belongs to.

Packages provide modularity and code organization. By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with `package rand`.

### Go Example

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

## C# Equivalent

C# uses **namespaces** to organize code, which serve a similar purpose to Go packages. However, there are important differences in how they work:

- In C#, namespaces are purely for logical organization and name scoping
- In Go, packages are the unit of compilation and encapsulation
- C# has separate concepts for assemblies (compilation units) and namespaces (organization)
- Go's package system is simpler: one package = one compilation unit

The entry point in C# is typically a `Main` method in any class, though modern C# (9.0+) supports top-level statements similar to Go's approach.

### C# Example

```csharp
using System;

namespace MyProgram
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("My favorite number is " + Random.Shared.Next(10));
        }
    }
}

// Or with C# 9.0+ top-level statements:
// using System;
// Console.WriteLine("My favorite number is " + Random.Shared.Next(10));
```

## Key Differences

- **Package vs Namespace**: Go packages are compilation units; C# namespaces are just naming constructs
- **Main Entry Point**: Go requires `package main` and a `main()` function; C# requires a `Main` method (or top-level statements)
- **File Organization**: All files in a Go directory must be in the same package; C# files can contain multiple namespaces
- **Visibility**: Go uses capitalization for exports; C# uses explicit `public`/`private` keywords
- **Use Case Alignment**: Both organize code and prevent naming conflicts, but Go's packages are more tightly coupled to the build system
