# Exported Names

## Go Concept

In Go, a name is **exported** (visible outside its package) if it begins with a capital letter. Names starting with a lowercase letter are **unexported** (private to the package).

This capitalization rule applies to:
- Functions (`Println` is exported, `println` is not)
- Types (`Reader` is exported, `reader` is not)
- Struct fields (`Name` is exported, `name` is not)
- Constants and variables

When importing a package, you can only refer to its exported names. Any unexported names are not accessible from outside the package.

### Go Example

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)  // Pi is exported (uppercase)
	// fmt.Println(math.pi)  // ERROR: pi is unexported (lowercase)
}
```

## C# Equivalent

C# uses **access modifiers** (`public`, `private`, `protected`, `internal`) to control visibility. This provides more granular control than Go's simple capitalization rule:

- `public`: Accessible from anywhere (like Go's exported names)
- `private`: Accessible only within the class/struct (similar to Go's unexported for methods)
- `internal`: Accessible within the same assembly (no direct Go equivalent)
- `protected`: Accessible within the class and derived classes

C# requires explicit keywords, making visibility intentions more obvious but more verbose. Go's approach is simpler but less flexible.

### C# Example

```csharp
using System;

namespace MathLibrary
{
    public class Constants
    {
        public const double Pi = 3.14159;      // Accessible from anywhere
        private const double pi = 3.14159;     // Only within this class
        internal const double E = 2.71828;     // Within this assembly
    }
}

// Usage:
class Program
{
    static void Main()
    {
        Console.WriteLine(Constants.Pi);    // OK: Pi is public
        // Console.WriteLine(Constants.pi);  // ERROR: pi is private
        Console.WriteLine(Constants.E);      // OK: E is internal (same assembly)
    }
}
```

## Key Differences

- **Visibility Mechanism**: Go uses capitalization; C# uses explicit keywords
- **Granularity**: C# has more access levels (public, private, protected, internal, protected internal); Go has only two (exported/unexported)
- **Package vs Assembly**: Go's unexported is package-scoped; C# has both class-scoped (`private`) and assembly-scoped (`internal`)
- **Default Behavior**: Go defaults to unexported (lowercase); C# defaults to `private` for members, `internal` for top-level types
- **Clarity vs Brevity**: C# is more explicit and verbose; Go is more concise but relies on naming conventions
- **Use Case Alignment**: Both control API surface area, but C# provides finer-grained control for complex scenarios like inheritance and assembly boundaries. Go's simpler model works well for its composition-based design.
