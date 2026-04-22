# Exercise: Stringers

## Challenge

Make the `IPAddr` type implement `fmt.Stringer` to print the address as a dotted quad.

For instance, `IPAddr{1, 2, 3, 4}` should print as `"1.2.3.4"`.

## Starter Code

```go
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```

## Expected Output

```
loopback: 127.0.0.1
googleDNS: 8.8.8.8
```

## Hints

- The `fmt.Sprintf` function can help format strings with values
- Access array elements using index notation: `ip[0]`, `ip[1]`, etc.
- The format verb `%d` prints integers

## Solution

<details>
<summary>Click to reveal solution</summary>

```go
package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```

## C# Equivalent Exercise

```csharp
using System;
using System.Linq;

// Make IPAddr override ToString() to print as dotted quad
struct IPAddr
{
    private readonly byte[] octets;
    
    public IPAddr(byte a, byte b, byte c, byte d)
    {
        octets = new byte[] { a, b, c, d };
    }
    
    // TODO: Override ToString()
}

class Program
{
    static void Main()
    {
        var hosts = new Dictionary<string, IPAddr>
        {
            ["loopback"] = new IPAddr(127, 0, 0, 1),
            ["googleDNS"] = new IPAddr(8, 8, 8, 8)
        };
        
        foreach (var (name, ip) in hosts)
        {
            Console.WriteLine($"{name}: {ip}");
        }
    }
}
```

**C# Solution:**

```csharp
public override string ToString()
{
    return string.Join(".", octets);
    // OR: return $"{octets[0]}.{octets[1]}.{octets[2]}.{octets[3]}";
}
```

</details>

## Extension Challenges

1. **IPv6 Support**: Create an `IPAddr6` type that handles 16 bytes and formats as IPv6 (e.g., `2001:0db8:85a3::8a2e:0370:7334`)

2. **CIDR Notation**: Extend IPAddr to include a subnet mask and format as CIDR notation (e.g., `192.168.1.0/24`)

3. **Validation**: Add a constructor/factory function that validates the IP address bytes

4. **Formatting Options**: In C#, implement `IFormattable` to support different formats (dotted quad, hexadecimal, binary)

## Key Learning Points

- Implementing `fmt.Stringer` changes how your type is printed
- The `String()` method is called automatically by fmt package functions
- Array indexing works the same in the method implementation
- Custom string representations make debugging and logging easier
- In C#, override `ToString()` for the same effect
- Both approaches allow types to control their text representation
