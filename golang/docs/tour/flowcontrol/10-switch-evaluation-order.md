# Switch Evaluation Order

## Go Concept

Switch cases evaluate from top to bottom, **stopping when a case succeeds**. This is important for cases with side effects or expensive computations.

Cases are evaluated in order, and only the first matching case is executed. Subsequent cases are not even evaluated.

### Go Example

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```

## C# Equivalent

C# switch statements also evaluate cases in order and execute only the first matching case. However:

- C# switch traditionally only allowed constant case values
- C# 7.0+ pattern matching allows expressions, but they must be compile-time constants for traditional switches
- C# switch expressions (8.0+) can use runtime expressions

The evaluation order is the same, but C# has historically been more restrictive about what can be in a case.

### C# Example

```csharp
using System;

class Program
{
    static void Main()
    {
        Console.WriteLine("When's Saturday?");
        var today = DateTime.Now.DayOfWeek;
        var daysUntilSaturday = ((int)DayOfWeek.Saturday - (int)today + 7) % 7;

        // Traditional switch with constant cases
        switch (daysUntilSaturday)
        {
            case 0:
                Console.WriteLine("Today.");
                break;
            case 1:
                Console.WriteLine("Tomorrow.");
                break;
            case 2:
                Console.WriteLine("In two days.");
                break;
            default:
                Console.WriteLine("Too far away.");
                break;
        }

        // C# 8.0+ switch expression (cleaner)
        var message = daysUntilSaturday switch
        {
            0 => "Today.",
            1 => "Tomorrow.",
            2 => "In two days.",
            _ => "Too far away."
        };
        Console.WriteLine(message);
    }
}
```

## Key Differences

- **Expression Cases**: Go allows runtime expression cases directly; C# traditionally requires constants
- **Order Guarantee**: Both evaluate top-to-bottom and stop at first match
- **Side Effects**: Both only evaluate cases until a match is found
- **Pattern Matching**: Modern C# (7.0+) supports pattern matching cases; Go doesn't have pattern matching
- **Switch Expressions**: C# 8.0+ switch expressions make this pattern more concise; Go uses traditional switch
- **Use Case Alignment**: Both evaluate cases in order and stop at the first match. Go allows more flexible expressions in cases directly, while C# historically required computing the value first and switching on it. Modern C# switch expressions provide similar flexibility and conciseness.

The key learning point (evaluation order, stop at first match) is **identical in both languages**. The difference is that Go has always been more flexible with expression cases, while C# added this flexibility gradually through pattern matching and switch expressions.
