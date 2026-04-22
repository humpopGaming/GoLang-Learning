# Imports

## Go Concept

Go uses the `import` keyword to bring in other packages. You can write multiple import statements or use a "factored" import statement with parentheses, which is the preferred style.

Import paths are string literals that specify the location of packages. For standard library packages, these are simple names like "fmt" or "math". For external packages, they're typically URLs like "github.com/user/package".

### Go Example

```go
package main

// Multiple import statements (works but not idiomatic)
import "fmt"
import "math"

// Factored import (preferred Go style)
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```

## C# Equivalent

C# uses **using directives** to bring namespaces into scope. The syntax is similar, but the underlying mechanism differs:

- C# `using` statements bring namespace contents into scope
- Go `import` statements make package identifiers available with their package name prefix
- C# can alias namespaces and types with `using` aliases
- Go always requires the package name prefix (e.g., `fmt.Println`) unless you use dot imports (discouraged)

Both languages support grouping imports for cleaner code.

### C# Example

```csharp
using System;
using System.Math;  // Note: Math is a class, not a namespace

// Or grouped (C# 10+):
// using System;
// using static System.Math;  // For static members

namespace MyProgram
{
    class Program
    {
        static void Main()
        {
            // Without 'using static', you need: Math.Sqrt(7)
            Console.WriteLine($"Now you have {Math.Sqrt(7)} problems.");
        }
    }
}
```

## Key Differences

- **Prefix Requirements**: Go always requires package prefix (e.g., `fmt.Println`); C# brings names directly into scope
- **Static Imports**: C# has `using static` for importing static members; Go doesn't have an equivalent
- **Import Path Format**: Go uses paths/URLs; C# uses namespace names
- **Factored Style**: Both support grouping, but Go strongly prefers the factored import style
- **Unused Imports**: Go is strict about unused imports (compile error); C# issues warnings
- **Use Case Alignment**: Both reduce typing and organize dependencies, but Go's approach maintains clearer package boundaries
