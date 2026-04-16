# Challenge 21: Todo CLI (Capstone)

## Objective

Combine everything you've learned into a single project: a command-line todo list application. This challenge ties together packages, functions, types, control flow, structs, slices, maps, methods, interfaces, error handling, and concurrency.

## What You'll Learn

This capstone reinforces all previous concepts by using them in one cohesive program:
- Packages and imports
- Functions with multiple return values
- Variables, constants, type conversions
- `for` loops, `if/else`, `switch`
- Structs with methods (pointer receivers)
- Slices, maps, and `range`
- Interfaces (`fmt.Stringer`, `error`)
- Closures (for ID generation)
- Goroutines and channels (for autosave)
- `sync.Mutex` (for thread-safe access)

## Tour Reference

All previous challenge tour pages apply. This is your chance to use everything together without looking things up.

## What to Build

A command-line todo list that:
- Stores todo items in a struct with ID, title, and done status
- Supports add, complete, remove, and list operations
- Generates unique IDs using a closure
- Has a Stringer interface for pretty printing
- Has error handling for invalid operations
- Runs an auto-save simulation in a background goroutine

## File to Create

```
challenge21/main.go
```

## Requirements

### Data Types
1. Define a `Todo` struct with fields: `ID int`, `Title string`, `Done bool`
2. Implement `fmt.Stringer` on `Todo` — format: `"[X] #1: Buy groceries"` (done) or `"[ ] #1: Buy groceries"` (not done)
3. Define a `TodoList` struct with:
   - `mu sync.Mutex`
   - `items []Todo`
   - `nextID func() int` — a closure that generates incrementing IDs

### Error Handling
4. Define `type NotFoundError struct { ID int }` implementing the `error` interface with message `"todo not found: #<ID>"`

### Methods on TodoList
5. `NewTodoList() *TodoList` — constructor that initializes the list and creates the ID-generator closure (starts at 1)
6. `Add(title string) Todo` — adds a new todo and returns it (use mutex)
7. `Complete(id int) error` — marks a todo as done; returns `NotFoundError` if ID doesn't exist (use mutex)
8. `Remove(id int) error` — removes a todo by ID; returns `NotFoundError` if ID doesn't exist (use mutex)
9. `List() []Todo` — returns a copy of all todos (use mutex)
10. `Summary() (total int, done int, pending int)` — returns counts using named return values (use mutex)

### Concurrency
11. Write `func autoSave(ch <-chan string, done chan<- bool)` — a goroutine that listens on `ch` for "save" signals and prints `"[AutoSave] Saving..."` each time. When the channel is closed, it prints `"[AutoSave] Shutting down."` and sends `true` to `done`.
12. In `main()`, create a save channel, start `autoSave` as a goroutine, and send a save signal after each add/complete/remove operation.

### Main Flow
13. In `main()`:
    - Create a `TodoList`
    - Start `autoSave`
    - Add: "Learn Go basics", "Build a CLI app", "Practice concurrency"
    - List all todos
    - Complete todo #1
    - Complete todo #999 (should print error)
    - Remove todo #2
    - List all todos again
    - Print summary
    - Close the save channel and wait for autosave to shut down

## Expected Output

```
Added: [ ] #1: Learn Go basics
[AutoSave] Saving...
Added: [ ] #2: Build a CLI app
[AutoSave] Saving...
Added: [ ] #3: Practice concurrency
[AutoSave] Saving...

All todos:
  [ ] #1: Learn Go basics
  [ ] #2: Build a CLI app
  [ ] #3: Practice concurrency

Completed: [X] #1: Learn Go basics
[AutoSave] Saving...
Error: todo not found: #999

Removed todo #2
[AutoSave] Saving...

All todos:
  [X] #1: Learn Go basics
  [ ] #3: Practice concurrency

Summary: 2 total, 1 done, 1 pending

[AutoSave] Shutting down.
```

> Note: The `[AutoSave]` lines may interleave slightly differently due to goroutine scheduling — that's fine.

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge21/main.go
```

Verify thread safety:
```bash
go run -race challenge21/main.go
```

## Hints

- ID generator closure:
  ```go
  func makeIDGenerator() func() int {
      id := 0
      return func() int {
          id++
          return id
      }
  }
  ```
- Stringer:
  ```go
  func (t Todo) String() string {
      status := " "
      if t.Done { status = "X" }
      return fmt.Sprintf("[%s] #%d: %s", status, t.ID, t.Title)
  }
  ```
- Removing from a slice by index `i`:
  ```go
  list.items = append(list.items[:i], list.items[i+1:]...)
  ```
- For the save channel, use `ch := make(chan string)` (unbuffered) — each send will block until `autoSave` receives
- Send a signal: `ch <- "save"` after each mutation
- To shut down: `close(ch)` — the `for range ch` loop in autoSave will exit
- Use `time.Sleep(50 * time.Millisecond)` after sending to let the autosave goroutine print before the next operation (for cleaner output)
