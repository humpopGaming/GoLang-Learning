# Pointers

## Go Concept

Go has **pointers**. A pointer holds the memory address of a value.

- The type `*T` is a pointer to a `T` value
- The zero value of a pointer is `nil`
- The `&` operator generates a pointer to its operand
- The `*` operator denotes the pointer's underlying value (dereferencing)

Unlike C, Go has no pointer arithmetic — you cannot increment pointers or do address calculations.

### Go Example

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

## C# Equivalent

C# has **pointers** in unsafe contexts and **references** in safe contexts:

- **References**: Most common, used with classes (reference types)
- **ref/out parameters**: Pass value types by reference
- **Unsafe pointers**: C-style pointers requiring `unsafe` keyword

For most use cases, C# developers use references (classes) or ref parameters. Unsafe pointers are rarely needed and require special compiler flags.

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        int i = 42, j = 2701;

        // C# approach: ref parameters
        ModifyValue(ref i);
        Console.WriteLine(i);  // 21

        DivideValue(ref j);
        Console.WriteLine(j);  // 73

        // Unsafe pointer (rarely used)
        unsafe
        {
            int value = 42;
            int* ptr = &value;
            Console.WriteLine(*ptr);  // 42
            *ptr = 21;
            Console.WriteLine(value);  // 21
        }
    }

    static void ModifyValue(ref int value)
    {
        value = 21;
    }

    static void DivideValue(ref int value)
    {
        value = value / 37;
    }
}
```

## Key Differences

- **Pointer Syntax**: Go uses `*T`, `&`, `*` operators; C# has similar syntax in unsafe code but typically uses `ref`
- **Safety**: Go pointers are memory-safe (no arithmetic); C# unsafe pointers allow arithmetic but require unsafe blocks
- **Common Usage**: Go uses pointers extensively; C# uses references (classes) and `ref` parameters more commonly
- **Value vs Reference Types**: C# distinguishes structs (value) and classes (reference); Go has pointers for both scenarios
- **Nil vs Null**: Go pointers can be `nil`; C# references can be `null`, C# pointers in unsafe context can be null
- **No Pointer Arithmetic**: Go doesn't allow pointer arithmetic; C# unsafe pointers do
- **Use Case Alignment**: Both allow passing by reference and modifying values through indirection. **Go uses pointers as the primary mechanism**; **C# uses references (for classes) and `ref`/`out` keywords (for value types)**. 

**Different philosophies**:
- Go: Pointers are common and safe, used for efficiency and mutation
- C#: References are built into classes; `ref`/`out` used for value types; unsafe pointers rarely needed

For passing large structs or allowing mutation, Go uses pointers; C# uses either ref parameters (value types) or naturally uses references (classes). The use cases align, but the mechanisms differ significantly.
