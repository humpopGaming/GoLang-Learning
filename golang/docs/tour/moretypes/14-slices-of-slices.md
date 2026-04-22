# Slices of Slices

## Go Concept

Slices can contain any type, including other slices. This creates **multi-dimensional slices** (often called "jagged arrays" because rows can have different lengths).

### Go Example

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// Players take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```

Output:
```
X _ X
O _ X
_ _ O
```

## C# Equivalent

C# supports **jagged arrays** (arrays of arrays) and **multi-dimensional arrays** (rectangular):

- **Jagged arrays**: `int[][]` — each row can have different length
- **Multi-dimensional arrays**: `int[,]` — all rows same length (rectangular)
- **Lists of Lists**: `List<List<T>>` — dynamic jagged structure

Go slices of slices are most like C# jagged arrays.

### C# Example

```csharp
using System;
using System.Collections.Generic;

class Program
{
    static void Main()
    {
        // Jagged array (most like Go slices of slices)
        string[][] board = new string[][]
        {
            new string[] { "_", "_", "_" },
            new string[] { "_", "_", "_" },
            new string[] { "_", "_", "_" }
        };

        // Players take turns
        board[0][0] = "X";
        board[2][2] = "O";
        board[1][2] = "X";
        board[1][0] = "O";
        board[0][2] = "X";

        for (int i = 0; i < board.Length; i++)
        {
            Console.WriteLine(string.Join(" ", board[i]));
        }

        Console.WriteLine("\nList of Lists (dynamic):");
        // List of Lists (dynamic jagged structure)
        List<List<string>> dynamicBoard = new List<List<string>>
        {
            new List<string> { "_", "_", "_" },
            new List<string> { "_", "_", "_" },
            new List<string> { "_", "_", "_" }
        };

        dynamicBoard[0][0] = "X";
        dynamicBoard[2][2] = "O";

        foreach (var row in dynamicBoard)
        {
            Console.WriteLine(string.Join(" ", row));
        }

        Console.WriteLine("\nMulti-dimensional array (rectangular):");
        // Multi-dimensional array - different syntax
        string[,] rectBoard = new string[,]
        {
            { "_", "_", "_" },
            { "_", "_", "_" },
            { "_", "_", "_" }
        };

        rectBoard[0, 0] = "X";
        rectBoard[2, 2] = "O";

        for (int i = 0; i < rectBoard.GetLength(0); i++)
        {
            for (int j = 0; j < rectBoard.GetLength(1); j++)
            {
                Console.Write(rectBoard[i, j] + " ");
            }
            Console.WriteLine();
        }
    }
}
```

Output:
```
X _ X
O _ X
_ _ O

List of Lists (dynamic):
X _ _
_ _ _
_ _ O

Multi-dimensional array (rectangular):
X _ _
_ _ _
_ _ O
```

## Key Differences

- **Type Syntax**: Go uses `[][]T`; C# uses `T[][]` (jagged) or `T[,]` (rectangular)
- **Jagged vs Rectangular**: Go only has jagged (flexible); C# has both jagged and rectangular arrays
- **Indexing**: Go uses `arr[i][j]`; C# jagged uses `arr[i][j]`; C# rectangular uses `arr[i,j]`
- **Variable Row Length**: Go slices of slices and C# jagged arrays allow it; C# rectangular arrays don't
- **Initialization**: Go slice literals more concise; C# requires `new` keywords
- **Use Case Alignment**: Both support nested structures, but Go only has one way (flexible), C# has multiple ways.

**Syntax comparison**:
```go
// Go: always jagged
board := [][]string{
    {"_", "_", "_"},
    {"_", "_", "_"},
}
```

```csharp
// C# jagged (most like Go)
string[][] board = new string[][] {
    new string[] { "_", "_", "_" },
    new string[] { "_", "_", "_" }
};

// C# rectangular (different structure)
string[,] board = new string[,] {
    { "_", "_", "_" },
    { "_", "_", "_" }
};
```

**Different philosophies**:
- **Go**: One flexible structure (slices of slices) — can have rows of different lengths
- **C#**: Two structures:
  - Jagged arrays `T[][]` — flexible, like Go
  - Rectangular arrays `T[,]` — fixed dimensions, more efficient for uniform data

**Memory layout differences**:
- Go `[][]int`: Each inner slice can be separate allocation
- C# `int[][]`: Each inner array is separate allocation (like Go)
- C# `int[,]`: Single contiguous memory block (more efficient, but less flexible)

**Key insight**: For dynamic, potentially irregular multi-dimensional data:
- **Go**: Use `[][]T` (only option)
- **C#**: Use `T[][]` jagged arrays (matches Go behavior)

For fixed, rectangular multi-dimensional data:
- **Go**: Still use `[][]T` (no special rectangular type)
- **C#**: Consider `T[,]` for better performance and memory efficiency

**Common use cases**:
- Game boards, matrices, tables with uniform size: C# `T[,]` is more efficient
- Variable-length rows (e.g., different-length sentences): Go `[][]T` and C# `T[][]` are appropriate

The Go approach is simpler (one type), but C# offers optimization opportunities for rectangular data at the cost of complexity.
