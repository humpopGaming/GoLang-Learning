# Generic Types

## Go Concept

In addition to generic functions, Go supports **generic types** (structs, interfaces, and other type definitions). A generic type has type parameters that are specified when you create an instance of that type. This allows you to create data structures that work with any type while maintaining type safety.

Generic types are particularly useful for container types like lists, trees, and custom collections.

### Go Example

```go
package main

import "fmt"

// List is a generic singly-linked list that can hold values of any type
type List[T any] struct {
	next *List[T]
	val  T
}

// Add appends a new element to the end of the list
func (l *List[T]) Add(val T) {
	if l.next == nil {
		l.next = &List[T]{val: val}
	} else {
		l.next.Add(val)
	}
}

// PrintAll prints all elements in the list
func (l *List[T]) PrintAll() {
	current := l
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}

func main() {
	// Create a list of integers
	var intList List[int]
	intList.val = 1
	intList.Add(2)
	intList.Add(3)
	
	fmt.Println("Integer list:")
	intList.PrintAll()
	
	// Create a list of strings
	var strList List[string]
	strList.val = "hello"
	strList.Add("world")
	
	fmt.Println("\nString list:")
	strList.PrintAll()
}
```

## C# Equivalent

C# has supported **generic classes and structs** since C# 2.0. The syntax and capabilities are very similar to Go's generic types. C# generic classes are the foundation of the .NET collections framework (List<T>, Dictionary<K,V>, etc.).

Key differences:
- Go uses square brackets `[]`; C# uses angle brackets `<>`
- C# has covariance/contravariance for generic interfaces (Go does not)
- C# can have static members in generic types (Go cannot have type-level state)
- Both support constraints on type parameters

### C# Example

```csharp
using System;

// Generic linked list class
public class List<T>
{
    public List<T> Next { get; set; }
    public T Value { get; set; }

    public List(T value)
    {
        Value = value;
    }

    // Add appends a new element to the end of the list
    public void Add(T value)
    {
        if (Next == null)
        {
            Next = new List<T>(value);
        }
        else
        {
            Next.Add(value);
        }
    }

    // PrintAll prints all elements in the list
    public void PrintAll()
    {
        var current = this;
        while (current != null)
        {
            Console.WriteLine(current.Value);
            current = current.Next;
        }
    }
}

public class Program
{
    public static void Main()
    {
        // Create a list of integers
        var intList = new List<int>(1);
        intList.Add(2);
        intList.Add(3);
        
        Console.WriteLine("Integer list:");
        intList.PrintAll();
        
        // Create a list of strings
        var strList = new List<string>("hello");
        strList.Add("world");
        
        Console.WriteLine("\nString list:");
        strList.PrintAll();
    }
}
```

## Key Comparison

| Feature | Go | C# |
|---------|----|----|
| Syntax | `type Name[T constraint] struct` | `class Name<T> where T : constraint` |
| Instantiation | `var x Name[int]` | `var x = new Name<int>()` |
| Methods | Methods can use type parameters | Methods can use type parameters |
| Nested generics | `List[List[int]]` | `List<List<int>>` |
| Variance | Not supported | Covariance (`out`), Contravariance (`in`) |
| Constraints | Interface-based | Keyword and interface-based |
