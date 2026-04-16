# Challenge 16: Generic Stack

## Objective

Implement a generic stack data structure that works with any type, practising generic type declarations with type parameters on structs.

## What You'll Learn

- Generic types: `type Stack[T any] struct { ... }`
- Methods on generic types
- How a type parameter on a struct lets you create one data structure for any type

## Tour Reference — Read These First

1. [Generic types](https://go.dev/tour/generics/2) — Types can be parameterized with type parameters

## What to Build

A `Stack[T]` data structure that supports Push, Pop, Peek, IsEmpty, and Size operations. Then demonstrate it with both `int` and `string` stacks.

## File to Create

```
challenge16/main.go
```

## Requirements

1. Define `type Stack[T any] struct` with a single field `items []T`
2. Implement these methods (all on `*Stack[T]`):
   - `Push(item T)` — adds an item to the top
   - `Pop() (T, bool)` — removes and returns the top item (returns zero value and `false` if empty)
   - `Peek() (T, bool)` — returns the top item without removing it (returns zero value and `false` if empty)
   - `IsEmpty() bool` — returns `true` if the stack has no items
   - `Size() int` — returns the number of items
3. In `main()`:
   - Create `intStack` as `Stack[int]{}` — push 10, 20, 30, then pop twice and print what you got
   - Create `strStack` as `Stack[string]{}` — push `"hello"`, `"world"`, peek, then pop all items

## Expected Output

```
--- Integer Stack ---
Pushed: 10, 20, 30
Size: 3
Popped: 30
Popped: 20
Size: 1
Is empty: false

--- String Stack ---
Pushed: hello, world
Peek: world
Popped: world
Popped: hello
Is empty: true
Pop from empty stack: "" (ok=false)
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge16/main.go
```

## Hints

- Generic struct:
  ```go
  type Stack[T any] struct {
      items []T
  }
  ```
- Push appends to the slice: `s.items = append(s.items, item)`
- Pop removes from the end (top of stack):
  ```go
  func (s *Stack[T]) Pop() (T, bool) {
      if len(s.items) == 0 {
          var zero T
          return zero, false
      }
      top := s.items[len(s.items)-1]
      s.items = s.items[:len(s.items)-1]
      return top, true
  }
  ```
- `var zero T` gives you the zero value of whatever type `T` is (0 for int, "" for string, etc.)
- Methods on generic types: `func (s *Stack[T]) Push(item T)` — the `[T]` appears in the receiver
