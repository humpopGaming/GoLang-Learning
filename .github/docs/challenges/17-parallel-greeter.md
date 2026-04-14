# Challenge 17: Parallel Greeter

## Objective

Learn goroutines by writing a program that greets multiple people in parallel, seeing how concurrent execution interleaves output.

## What You'll Learn

- Starting a goroutine with the `go` keyword: `go functionName(args)`
- Goroutines are lightweight threads managed by the Go runtime
- Goroutines run concurrently — their output interleaves unpredictably
- The main goroutine must stay alive for other goroutines to finish
- Using `time.Sleep` as a simple (temporary) way to wait for goroutines

## Tour Reference — Read These First

1. [Goroutines](https://go.dev/tour/concurrency/1) — `go f(x)` starts a new goroutine running `f(x)`

## What to Build

A program that:
1. Greets several people concurrently using goroutines
2. Shows that goroutine output is interleaved (non-deterministic order)
3. Counts from 1 to 5 in parallel from two goroutines to demonstrate interleaving

## File to Create

```
challenge17/main.go
```

## Requirements

1. Write `func greet(name string)` that prints `"Hello, <name>! (from goroutine)"` three times with a short delay between each (use `time.Sleep(100 * time.Millisecond)`)
2. Write `func countUp(label string)` that prints `"<label>: 1"`, `"<label>: 2"`, ... `"<label>: 5"` with a short delay between each
3. In `main()`:
   - Start `greet("Alice")`, `greet("Bob")`, and `greet("Charlie")` as goroutines using `go`
   - Wait long enough for them to finish (use `time.Sleep(500 * time.Millisecond)`)
   - Print a separator line
   - Start `countUp("Counter A")` and `countUp("Counter B")` as goroutines
   - Wait for them to finish
   - Print `"All done!"`

## Expected Output

Output will vary each run because goroutines are concurrent! Something like:

```
Hello, Charlie! (from goroutine)
Hello, Alice! (from goroutine)
Hello, Bob! (from goroutine)
Hello, Alice! (from goroutine)
Hello, Charlie! (from goroutine)
Hello, Bob! (from goroutine)
Hello, Bob! (from goroutine)
Hello, Alice! (from goroutine)
Hello, Charlie! (from goroutine)
---
Counter A: 1
Counter B: 1
Counter B: 2
Counter A: 2
Counter A: 3
Counter B: 3
Counter A: 4
Counter B: 4
Counter A: 5
Counter B: 5
All done!
```

> The order of lines will be different each time you run it. That's the whole point!

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge17/main.go
```

## Hints

- Starting a goroutine: `go greet("Alice")` — this returns immediately, and `greet` runs in the background
- Import `"time"` for `time.Sleep` and `time.Millisecond`
- `time.Sleep` is a crude way to wait — in Challenge 18 you'll learn channels, which are the proper way
- If your program ends before the goroutines finish, you won't see their output — that's why we sleep in `main()`
- The `go` keyword works with any function call: `go myFunc(arg1, arg2)`
