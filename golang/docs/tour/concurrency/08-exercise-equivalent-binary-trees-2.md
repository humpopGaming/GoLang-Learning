# Exercise: Equivalent Binary Trees (Part 2)

## Continuation

This is the continuation/solution discussion for the [Equivalent Binary Trees exercise](./07-exercise-equivalent-binary-trees.md). If you haven't attempted the exercise yet, go back and try it first!

## Solution Walkthrough

### Step 1: Implement Walk

The key insight is to use **in-order traversal** (left, root, right) to visit nodes in sorted order.

```go
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkHelper(t, ch)
	close(ch) // Close channel when done
}

// walkHelper does the recursive traversal
func walkHelper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkHelper(t.Left, ch)   // Visit left subtree
	ch <- t.Value             // Send value
	walkHelper(t.Right, ch)  // Visit right subtree
}
```

### Step 2: Implement Same

Compare values from two trees by walking both concurrently:

```go
// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	// Walk both trees concurrently
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	// Compare values from both channels
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		
		// If one is closed but not the other, different lengths
		if ok1 != ok2 {
			return false
		}
		
		// Both closed, trees are same
		if !ok1 {
			return true
		}
		
		// Values differ
		if v1 != v2 {
			return false
		}
	}
}
```

### Alternative Same Implementation (using range)

```go
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 {
			return false
		}
	}
	
	// Check if ch2 has leftover values
	_, ok := <-ch2
	return !ok // Should be closed
}
```

## Complete Solution

```go
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkHelper(t, ch)
	close(ch)
}

func walkHelper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkHelper(t.Left, ch)
	ch <- t.Value
	walkHelper(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		
		if ok1 != ok2 {
			return false
		}
		
		if !ok1 {
			return true
		}
		
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	// Test Walk
	fmt.Println("Walking tree.New(1):")
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}

	// Test Same
	fmt.Println("\nTesting Same:")
	fmt.Println("tree.New(1) == tree.New(1):", Same(tree.New(1), tree.New(1)))
	fmt.Println("tree.New(1) == tree.New(2):", Same(tree.New(1), tree.New(2)))
}
```

## C# Complete Solution

### Using Channel<T>

```csharp
using System;
using System.Threading.Channels;
using System.Threading.Tasks;

class Program
{
    static async Task Walk(Tree t, ChannelWriter<int> writer)
    {
        await WalkHelper(t, writer);
        writer.Complete();
    }
    
    static async Task WalkHelper(Tree t, ChannelWriter<int> writer)
    {
        if (t == null) return;
        
        await WalkHelper(t.Left, writer);
        await writer.WriteAsync(t.Value);
        await WalkHelper(t.Right, writer);
    }
    
    static async Task<bool> Same(Tree t1, Tree t2)
    {
        var ch1 = Channel.CreateUnbounded<int>();
        var ch2 = Channel.CreateUnbounded<int>();
        
        var walk1 = Walk(t1, ch1.Writer);
        var walk2 = Walk(t2, ch2.Writer);
        
        while (true)
        {
            var read1Task = ch1.Reader.WaitToReadAsync();
            var read2Task = ch2.Reader.WaitToReadAsync();
            
            var ok1 = await read1Task;
            var ok2 = await read2Task;
            
            if (ok1 != ok2)
                return false;
            
            if (!ok1)
                return true;
            
            var v1 = await ch1.Reader.ReadAsync();
            var v2 = await ch2.Reader.ReadAsync();
            
            if (v1 != v2)
                return false;
        }
    }
    
    static async Task Main()
    {
        // Test Walk
        Console.WriteLine("Walking Tree.New(1):");
        var channel = Channel.CreateUnbounded<int>();
        var walkTask = Walk(Tree.New(1), channel.Writer);
        
        await foreach (var value in channel.Reader.ReadAllAsync())
        {
            Console.WriteLine(value);
        }
        await walkTask;
        
        // Test Same
        Console.WriteLine("\nTesting Same:");
        Console.WriteLine($"Tree.New(1) == Tree.New(1): {await Same(Tree.New(1), Tree.New(1))}");
        Console.WriteLine($"Tree.New(1) == Tree.New(2): {await Same(Tree.New(1), Tree.New(2))}");
    }
}
```

### Using IAsyncEnumerable (More Idiomatic)

```csharp
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

class Program
{
    static async IAsyncEnumerable<int> Walk(Tree t)
    {
        if (t == null) yield break;
        
        await foreach (var value in Walk(t.Left))
            yield return value;
        
        yield return t.Value;
        
        await foreach (var value in Walk(t.Right))
            yield return value;
    }
    
    static async Task<bool> Same(Tree t1, Tree t2)
    {
        await using var enum1 = Walk(t1).GetAsyncEnumerator();
        await using var enum2 = Walk(t2).GetAsyncEnumerator();
        
        while (true)
        {
            var has1 = await enum1.MoveNextAsync();
            var has2 = await enum2.MoveNextAsync();
            
            if (has1 != has2)
                return false;
            
            if (!has1)
                return true;
            
            if (enum1.Current != enum2.Current)
                return false;
        }
    }
    
    static async Task Main()
    {
        // Test Walk
        Console.WriteLine("Walking Tree.New(1):");
        await foreach (var value in Walk(Tree.New(1)))
        {
            Console.WriteLine(value);
        }
        
        // Test Same
        Console.WriteLine("\nTesting Same:");
        Console.WriteLine($"Tree.New(1) == Tree.New(1): {await Same(Tree.New(1), Tree.New(1))}");
        Console.WriteLine($"Tree.New(1) == Tree.New(2): {await Same(Tree.New(1), Tree.New(2))}");
    }
}
```

## Key Learnings

### Go Concepts

1. **Goroutines for traversal** - Concurrent tree walking
2. **Channels for results** - Collecting values from tree
3. **Closing channels** - Signaling completion
4. **Range over channels** - Consuming until closed
5. **Multiple channel coordination** - Comparing two streams

### C# Concepts

1. **Async methods** - All traversal becomes async
2. **Channel<T>** - Direct port of Go channels
3. **IAsyncEnumerable** - More idiomatic for sequences
4. **await foreach** - Async iteration
5. **Multiple async operations** - Task coordination

## Performance Notes

**Go:**
- Goroutines are very lightweight
- Creating goroutines for tree traversal is efficient
- Channel operations are fast

**C#:**
- Tasks have more overhead than goroutines
- IAsyncEnumerable is more efficient for sequences
- Consider synchronous version for CPU-bound work

## When to Use Each Approach

**Go's approach** is best when:
- You need true concurrent tree processing
- Trees are large and traversal is expensive
- You're building concurrent systems

**C# Channel approach** is best when:
- Porting Go code to C#
- You need explicit producer-consumer pattern
- Building concurrent pipelines

**C# IAsyncEnumerable** is best when:
- Tree traversal is I/O-bound (remote trees)
- You want idiomatic C# code
- Simplicity is preferred over explicit concurrency

## Extension Exercises

1. **Count Nodes**: Modify to count total nodes instead of comparing
2. **Find Value**: Use channels to search for a value across multiple trees concurrently
3. **Parallel Traversal**: Process tree nodes in parallel (visit multiple branches concurrently)
4. **Timeout**: Add a timeout to the tree comparison
