# Go Tour Reference with C# Cross-Reference

This directory contains comprehensive Go Tour documentation with C# equivalents. Each page covers a Go concept with examples and explains how it maps to C#, including detailed discussions of differences where no direct equivalent exists.

## 🌐 View in Browser

**Open `index.html` in your web browser** for the best reading experience! The HTML version includes:
- Interactive navigation sidebar
- Syntax-highlighted code examples
- Previous/Next page navigation buttons
- Responsive mobile-friendly design
- Visual distinction between Go and C# sections

Simply open `golang/docs/tour/index.html` in any modern web browser.

## How to Use This Tour

Each page includes:
- Go concept explanation with code examples
- C# equivalent with code examples
- Detailed alignment discussion
- Key differences highlighted

Start with basics and progress through the sections, or jump to specific concepts as needed.

## Tour Contents

### Basics

1. [Packages](basics/01-packages.md) — How Go organizes code into packages
2. [Imports](basics/02-imports.md) — Importing packages into your code
3. [Exported Names](basics/03-exported-names.md) — Public visibility in Go
4. [Functions](basics/04-functions.md) — Declaring and calling functions
5. [Functions Continued](basics/05-functions-continued.md) — Type shorthand for parameters
6. [Multiple Results](basics/06-multiple-results.md) — Functions returning multiple values
7. [Named Return Values](basics/07-named-return-values.md) — Naming return values for clarity
8. [Variables](basics/08-variables.md) — Declaring variables with var
9. [Variables with Initializers](basics/09-variables-with-initializers.md) — Declaring and initializing variables
10. [Short Variable Declarations](basics/10-short-variable-declarations.md) — The := syntax for local variables
11. [Basic Types](basics/11-basic-types.md) — Go's built-in types
12. [Zero Values](basics/12-zero-values.md) — Default values for uninitialized variables
13. [Type Conversions](basics/13-type-conversions.md) — Converting between types
14. [Type Inference](basics/14-type-inference.md) — Automatic type detection
15. [Constants](basics/15-constants.md) — Declaring constant values
16. [Numeric Constants](basics/16-numeric-constants.md) — High-precision numeric constants

### Flow Control

1. [For](flowcontrol/01-for.md) — The standard for loop
2. [For Continued](flowcontrol/02-for-continued.md) — For loop variations
3. [For is Go's While](flowcontrol/03-for-is-while.md) — While-style loops in Go
4. [Forever](flowcontrol/04-forever.md) — Infinite loops
5. [If](flowcontrol/05-if.md) — Basic conditional statements
6. [If with Short Statement](flowcontrol/06-if-with-short-statement.md) — If with initialization
7. [If and Else](flowcontrol/07-if-and-else.md) — If-else chains
8. [Exercise: Loops and Functions](flowcontrol/08-exercise-loops-and-functions.md) — Practice with loops
9. [Switch](flowcontrol/09-switch.md) — Switch statements in Go
10. [Switch Evaluation Order](flowcontrol/10-switch-evaluation-order.md) — How switch cases are evaluated
11. [Switch with No Condition](flowcontrol/11-switch-with-no-condition.md) — Switch as cleaner if-else
12. [Defer](flowcontrol/12-defer.md) — Deferred function execution
13. [Stacking Defers](flowcontrol/13-stacking-defers.md) — Multiple defer statements

### More Types

1. [Pointers](moretypes/01-pointers.md) — Pointer basics and usage
2. [Structs](moretypes/02-structs.md) — Defining structured data types
3. [Struct Fields](moretypes/03-struct-fields.md) — Accessing struct members
4. [Pointers to Structs](moretypes/04-pointers-to-structs.md) — Working with struct pointers
5. [Struct Literals](moretypes/05-struct-literals.md) — Creating struct instances
6. [Arrays](moretypes/06-arrays.md) — Fixed-size arrays
7. [Slices](moretypes/07-slices.md) — Dynamic arrays in Go
8. [Slices are Like References](moretypes/08-slices-are-like-references.md) — Understanding slice behavior
9. [Slice Literals](moretypes/09-slice-literals.md) — Creating slices directly
10. [Slice Defaults](moretypes/10-slice-defaults.md) — Default bounds for slicing
11. [Slice Length and Capacity](moretypes/11-slice-length-and-capacity.md) — Understanding len and cap
12. [Nil Slices](moretypes/12-nil-slices.md) — Zero value for slices
13. [Creating Slices with Make](moretypes/13-creating-slices-with-make.md) — Using make to create slices
14. [Slices of Slices](moretypes/14-slices-of-slices.md) — Multi-dimensional slices
15. [Appending to Slices](moretypes/15-appending-to-slices.md) — Growing slices dynamically
16. [Range](moretypes/16-range.md) — Iterating over slices and arrays
17. [Range Continued](moretypes/17-range-continued.md) — Range variations
18. [Exercise: Slices](moretypes/18-exercise-slices.md) — Practice with slices
19. [Maps](moretypes/19-maps.md) — Key-value data structures
20. [Map Literals](moretypes/20-map-literals.md) — Creating maps directly
21. [Map Literals Continued](moretypes/21-map-literals-continued.md) — More map literal forms
22. [Mutating Maps](moretypes/22-mutating-maps.md) — Adding, updating, deleting map entries
23. [Exercise: Maps](moretypes/23-exercise-maps.md) — Practice with maps
24. [Function Values](moretypes/24-function-values.md) — Functions as first-class values
25. [Function Closures](moretypes/25-function-closures.md) — Functions that capture scope

### Methods and Interfaces

1. [Methods](methods/01-methods.md) — Methods vs functions
2. [Methods are Functions](methods/02-methods-are-functions.md) — Understanding method receivers
3. [Methods Continued](methods/03-methods-continued.md) — Methods on any type
4. [Pointer Receivers](methods/04-pointer-receivers.md) — Modifying receivers with pointers
5. [Pointers and Functions](methods/05-pointers-and-functions.md) — Comparing methods and functions
6. [Methods and Pointer Indirection](methods/06-methods-and-pointer-indirection.md) — Automatic dereferencing
7. [Methods and Pointer Indirection 2](methods/07-methods-and-pointer-indirection-2.md) — More on indirection
8. [Choosing a Value or Pointer Receiver](methods/08-choosing-value-or-pointer-receiver.md) — When to use each
9. [Interfaces](methods/09-interfaces.md) — Defining behavior contracts
10. [Interfaces are Implemented Implicitly](methods/10-interfaces-are-implemented-implicitly.md) — No explicit implements
11. [Interface Values](methods/11-interface-values.md) — How interfaces store values
12. [Interface Values with Nil Underlying Values](methods/12-interface-values-with-nil-underlying-values.md) — Nil handling
13. [Nil Interface Values](methods/13-nil-interface-values.md) — Uninitialized interfaces
14. [The Empty Interface](methods/14-the-empty-interface.md) — Interface{} and any
15. [Type Assertions](methods/15-type-assertions.md) — Extracting concrete types
16. [Type Switches](methods/16-type-switches.md) — Switching on type
17. [Stringers](methods/17-stringers.md) — The fmt.Stringer interface
18. [Exercise: Stringers](methods/18-exercise-stringers.md) — Practice with Stringers
19. [Errors](methods/19-errors.md) — The error interface
20. [Exercise: Errors](methods/20-exercise-errors.md) — Custom error types
21. [Readers](methods/21-readers.md) — The io.Reader interface
22. [Exercise: Readers](methods/22-exercise-readers.md) — Implementing readers
23. [Exercise: rot13Reader](methods/23-exercise-rot13reader.md) — Reader decorator pattern

### Generics

1. [Type Parameters](generics/01-type-parameters.md) — Generic functions with type parameters
2. [Generic Types](generics/02-generic-types.md) — Generic structs and types

### Concurrency

1. [Goroutines](concurrency/01-goroutines.md) — Lightweight concurrent execution
2. [Channels](concurrency/02-channels.md) — Communication between goroutines
3. [Buffered Channels](concurrency/03-buffered-channels.md) — Channels with capacity
4. [Range and Close](concurrency/04-range-and-close.md) — Iterating over channel values
5. [Select](concurrency/05-select.md) — Multiplexing channel operations
6. [Default Selection](concurrency/06-default-selection.md) — Non-blocking channel operations
7. [Exercise: Equivalent Binary Trees](concurrency/07-exercise-equivalent-binary-trees.md) — Trees with channels
8. [Exercise: Equivalent Binary Trees 2](concurrency/08-exercise-equivalent-binary-trees-2.md) — Continued
9. [sync.Mutex](concurrency/09-sync-mutex.md) — Mutual exclusion for shared data

## Navigation Tips

- **Sequential Learning**: Follow the sections in order (Basics → Flow Control → More Types → Methods → Generics → Concurrency)
- **Concept Reference**: Use this index to jump to specific concepts
- **Cross-Reference**: Each page includes C# equivalents to help translate between languages
- **Challenge Integration**: Pages are referenced by challenge files in the parent docs directory

## Additional Resources

- [Official Go Tour](https://go.dev/tour/) — Interactive version with exercises
- [Go Documentation](https://go.dev/doc/) — Official Go documentation
- [C# Documentation](https://learn.microsoft.com/en-us/dotnet/csharp/) — Official C# documentation
