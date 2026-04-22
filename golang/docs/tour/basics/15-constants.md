# Constants

## Go Concept

Constants are declared with the `const` keyword. Constants can be character, string, boolean, or numeric values. Constants cannot be declared using the `:=` syntax.

Constants in Go are compile-time values and can be used wherever a variable of the same type would be valid. They provide a way to give meaningful names to fixed values.

### Go Example

```go
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```

## C# Equivalent

C# has **constants** using the `const` keyword, which work very similarly to Go constants:

- Must be assigned at declaration
- Value must be compile-time constant
- Can be any primitive type or string
- Cannot use `var` with `const`

C# also has **readonly** fields, which are different from constants — they're assigned at runtime (in constructor) and can hold complex types. Go doesn't have an exact equivalent to `readonly`.

### C# Example

```csharp
using System;

class Program
{
    const double Pi = 3.14;
    
    static void Main()
    {
        const string World = "世界";
        Console.WriteLine($"Hello {World}");
        Console.WriteLine($"Happy {Pi} Day");

        const bool Truth = true;
        Console.WriteLine($"Go rules? {Truth}");
    }
}

// C# also has readonly (different concept):
class Config
{
    readonly DateTime CreatedAt;
    
    public Config()
    {
        CreatedAt = DateTime.Now;  // Set at runtime, not compile-time
    }
}
```

## Key Differences

- **Syntax**: Both use `const` keyword with similar syntax
- **Readonly Fields**: C# has `readonly` for runtime-assigned constants; Go doesn't have this concept
- **Type Inference**: Go constants can be untyped (covered in next section); C# constants must have a specific type
- **Scope**: Both support package/class level and function/method level constants
- **Allowed Types**: Go allows character, string, boolean, numeric; C# allows any primitive type, enums, or string
- **Computed Constants**: Neither language allows runtime-computed constants; both require compile-time values
- **Use Case Alignment**: Both provide compile-time constants for fixed values. The use cases align perfectly for primitive values and strings. **C# extends the concept with `readonly`** for values that should be immutable after initialization but aren't known at compile time (like timestamps, configuration loaded at startup). Go doesn't have this feature — you'd use regular variables and rely on convention not to modify them.

The main functional difference is C#'s `readonly` concept, which Go lacks. For pure compile-time constants, they work identically.
