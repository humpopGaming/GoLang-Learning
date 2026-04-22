# Exercise: Equivalent Binary Trees

## Go Exercise

**Objective:** Implement a function to determine whether two binary trees store the same sequence of values using goroutines and channels.

There can be many different binary trees with the same sequence of values stored in their leaves. For example, here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13:

```
    3             8
   / \           / \
  1   8         3  13
     / \       / \
    5  13     1   5
              \
               2
```

## The Problem

Write a function `Walk` that walks the tree and sends all values to a channel. Then write a function `Same` that uses `Walk` to determine whether two trees contain the same values.

### Tree Structure

```go
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
```

### What You'll Learn

- Tree traversal algorithms (in-order traversal)
- Using goroutines to perform concurrent work
- Using channels to collect results
- Closing channels to signal completion
- Using `range` to receive all values from a channel

### Reference

This exercise is from the [Tour of Go - Exercise: Equivalent Binary Trees](https://go.dev/tour/concurrency/7)

## Starter Code

```go
package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// TODO: Implement tree traversal
	// Don't forget to close the channel when done!
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// TODO: Use Walk to compare trees
	return false
}

func main() {
	// Test Walk
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}

	// Test Same
	fmt.Println(Same(tree.New(1), tree.New(1))) // should be true
	fmt.Println(Same(tree.New(1), tree.New(2))) // should be false
}
```

## Hints

<details>
<summary>Hint 1: Tree Traversal</summary>

You need to visit nodes in a specific order to get values in ascending order. For a binary search tree, in-order traversal (left, root, right) gives sorted values.
</details>

<details>
<summary>Hint 2: Recursive Helper</summary>

Create a helper function that does the actual traversal recursively. The main `Walk` function can start the goroutine and close the channel.
</details>

<details>
<summary>Hint 3: Closing the Channel</summary>

Remember to close the channel after sending all values. The `range` loop in the caller needs this to know when to stop.
</details>

<details>
<summary>Hint 4: Comparing Trees</summary>

For `Same`, create two channels and walk both trees concurrently. Compare values received from each channel. If they ever differ, return false.
</details>

## Expected Output

```
1 2 3 4 5 6 7 8 9 10
true
false
```

## C# Equivalent Exercise

In C#, you would use async/await with channels or async enumerables:

```csharp
using System;
using System.Collections.Generic;
using System.Threading.Channels;
using System.Threading.Tasks;

class Tree
{
    public Tree Left { get; set; }
    public int Value { get; set; }
    public Tree Right { get; set; }
    
    // Helper to create a test tree
    public static Tree New(int k)
    {
        // Implementation omitted for brevity
        // Returns a tree with values 1..10 in some structure
        return null;
    }
}

class Program
{
    // Walk the tree and send values to channel
    static async Task Walk(Tree t, ChannelWriter<int> writer)
    {
        // TODO: Implement tree traversal
        // Don't forget to call writer.Complete() when done!
    }
    
    // Helper for recursive traversal
    static async Task WalkHelper(Tree t, ChannelWriter<int> writer)
    {
        if (t == null) return;
        
        await WalkHelper(t.Left, writer);
        await writer.WriteAsync(t.Value);
        await WalkHelper(t.Right, writer);
    }
    
    // Determine if trees contain same values
    static async Task<bool> Same(Tree t1, Tree t2)
    {
        // TODO: Use Walk to compare trees
        return false;
    }
    
    static async Task Main()
    {
        // Test Walk
        var channel = Channel.CreateUnbounded<int>();
        var walkTask = Walk(Tree.New(1), channel.Writer);
        
        await foreach (var value in channel.Reader.ReadAllAsync())
        {
            Console.WriteLine(value);
        }
        
        await walkTask;
        
        // Test Same
        Console.WriteLine(await Same(Tree.New(1), Tree.New(1))); // true
        Console.WriteLine(await Same(Tree.New(1), Tree.New(2))); // false
    }
}
```

### C# Alternative: IAsyncEnumerable

A more C#-idiomatic approach using async enumerables:

```csharp
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
    var enum1 = Walk(t1).GetAsyncEnumerator();
    var enum2 = Walk(t2).GetAsyncEnumerator();
    
    while (true)
    {
        var has1 = await enum1.MoveNextAsync();
        var has2 = await enum2.MoveNextAsync();
        
        if (has1 != has2) return false;
        if (!has1) return true; // Both ended
        
        if (enum1.Current != enum2.Current)
            return false;
    }
}
```

## Key Comparison

| Concept | Go | C# (Channel) | C# (IAsyncEnumerable) |
|---------|----|--------------|-----------------------|
| **Traversal** | Send to channel in goroutine | WriteAsync in task | yield return |
| **Completion** | close(ch) | writer.Complete() | Method ends |
| **Consumption** | for v := range ch | await foreach | await foreach |
| **Recursion** | Natural with goroutines | Natural with async | Natural with yield |
| **Paradigm** | Concurrent message passing | Async sequence | Async sequence |

## Tips for C# Version

1. **Use IAsyncEnumerable** for simpler code - it's more idiomatic for this problem
2. **Remember async all the way** - tree traversal becomes async
3. **Don't forget Complete()** if using Channel - equivalent to closing in Go
4. **Consider synchronous version** - sometimes simpler if you don't need true concurrency
