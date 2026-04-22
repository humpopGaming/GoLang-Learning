# Stringers

## Go Concept

One of the most ubiquitous interfaces in Go is **Stringer**, defined by the `fmt` package:

```go
type Stringer interface {
    String() string
}
```

A Stringer is a type that can describe itself as a string. The `fmt` package (and many others) look for this interface to print values. Implementing `String()` is similar to implementing `toString()` in other languages.

When you implement the Stringer interface, functions like `fmt.Println` will automatically use your custom string representation instead of the default one.

### Go Example

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Person implements fmt.Stringer
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

// IPAddr implements fmt.Stringer
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

type Point struct {
	X, Y int
}

// Point with custom formatting
func (p Point) String() string {
	return fmt.Sprintf("Point{x=%d, y=%d}", p.X, p.Y)
}

func main() {
	p := Person{"Arthur Dent", 42}
	fmt.Println(p) // Arthur Dent (42 years)
	
	ip := IPAddr{127, 0, 0, 1}
	fmt.Println(ip) // 127.0.0.1
	
	pt := Point{10, 20}
	fmt.Println(pt) // Point{x=10, y=20}
	
	// Stringer is used implicitly
	fmt.Printf("Person: %v\n", p)  // Person: Arthur Dent (42 years)
	fmt.Printf("IP: %s\n", ip)     // IP: 127.0.0.1
}
```

## C# Equivalent

C# provides two main mechanisms for custom string representation:
1. **ToString() override**: Overriding `Object.ToString()` method
2. **IFormattable interface**: For format-specific string representations

Every object in C# inherits `ToString()` from `Object`, so you override it rather than implementing an interface. For more sophisticated formatting, you can implement `IFormattable` which allows different format strings.

### C# Example

```csharp
using System;

// Basic ToString override
class Person
{
    public string Name { get; set; }
    public int Age { get; set; }
    
    public override string ToString()
    {
        return $"{Name} ({Age} years)";
    }
}

// ToString with struct
struct IPAddr
{
    private readonly byte[] octets;
    
    public IPAddr(byte a, byte b, byte c, byte d)
    {
        octets = new byte[] { a, b, c, d };
    }
    
    public override string ToString()
    {
        return $"{octets[0]}.{octets[1]}.{octets[2]}.{octets[3]}";
    }
}

// IFormattable for advanced formatting
struct Point : IFormattable
{
    public int X { get; set; }
    public int Y { get; set; }
    
    // Standard ToString
    public override string ToString()
    {
        return ToString("G", null);
    }
    
    // IFormattable.ToString for custom formats
    public string ToString(string format, IFormatProvider provider)
    {
        if (string.IsNullOrEmpty(format)) format = "G";
        
        return format.ToUpperInvariant() switch
        {
            "G" => $"Point{{x={X}, y={Y}}}",
            "C" => $"({X}, {Y})",           // Compact
            "V" => $"[{X}; {Y}]",           // Vector notation
            _ => throw new FormatException($"Format '{format}' not supported")
        };
    }
}

// Record with custom ToString (C# 9.0+)
record Product(string Name, decimal Price)
{
    public override string ToString() => $"{Name}: ${Price:F2}";
}

class Program
{
    static void Main()
    {
        var p = new Person { Name = "Arthur Dent", Age = 42 };
        Console.WriteLine(p); // Arthur Dent (42 years)
        
        var ip = new IPAddr(127, 0, 0, 1);
        Console.WriteLine(ip); // 127.0.0.1
        
        var pt = new Point { X = 10, Y = 20 };
        Console.WriteLine(pt);              // Point{x=10, y=20}
        Console.WriteLine(pt.ToString("C", null)); // (10, 20)
        Console.WriteLine(pt.ToString("V", null)); // [10; 20]
        
        // ToString is used implicitly
        Console.WriteLine($"Person: {p}");  // Person: Arthur Dent (42 years)
        Console.WriteLine($"IP: {ip}");     // IP: 127.0.0.1
        
        var prod = new Product("Widget", 29.99m);
        Console.WriteLine(prod); // Widget: $29.99
    }
}
```

## Key Differences

- **Interface vs Override**: Go implements `Stringer` interface; C# overrides `ToString()` method from Object
- **Type System**: Go's approach is more decoupled (any type can implement); C# inherits from Object base class
- **Formatting Options**: Go has single `String()` method; C# can implement `IFormattable` for format strings
- **Implicit Use**: Both are called implicitly by formatting functions (fmt in Go, string interpolation in C#)
- **Null Safety**: Go doesn't call methods on nil interfaces; C# can call ToString() on null (returns empty string or throws)
- **Default Behavior**: Go prints struct field values by default; C# prints type name by default
- **Format Verbs**: Go uses format verbs like `%v`, `%s`; C# uses format strings and IFormattable
- **Performance**: Go's interface dispatch is simple; C# virtual method call has similar performance
- **Multiple Formats**: Go uses different format verbs; C# uses IFormattable with format strings
- **Discovery**: Go Stringer is interface-based, optional; C# ToString() is always available
- **Testing**: Both approaches are easy to test; Go can test interface conformance explicitly
- **Common Practice**: Both are idiomatic ways to provide string representations of custom types
