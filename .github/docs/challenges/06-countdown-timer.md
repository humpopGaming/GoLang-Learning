# Challenge 06: Countdown Timer

## Objective

Build a countdown program that demonstrates `defer` and stacking defers — showing how deferred function calls execute in LIFO (last-in-first-out) order when a function returns.

## What You'll Learn

- `defer` delays a function call until the surrounding function returns
- Deferred call arguments are evaluated immediately (not when the deferred call executes)
- Multiple defers stack — they execute in last-in-first-out (LIFO) order
- Practical use: cleanup actions that should always run

## Tour Reference — Read These First

1. [Defer](http://127.0.0.1:3999/tour/flowcontrol/12) — A defer statement defers execution until the surrounding function returns
2. [Stacking defers](http://127.0.0.1:3999/tour/flowcontrol/13) — Deferred calls are pushed onto a stack (LIFO)

## What to Build

A program with two functions:
1. A greeting function that uses `defer` to print a goodbye message after the greeting
2. A countdown function that uses deferred calls in a loop to print numbers in reverse order

## File to Create

```
challenge06/main.go
```

## Requirements

1. Write `func greet(name string)` that:
   - Defers a print of `"Goodbye, <name>!"` as the first line
   - Then prints `"Hello, <name>!"`
   - Then prints `"How are you, <name>?"`
   - When the function returns, the deferred goodbye should print last
2. Write `func countdown(from int)` that:
   - Prints `"Starting countdown..."`
   - Uses a `for` loop from 1 to `from`, deferring `fmt.Println(i)` in each iteration
   - After the loop, prints `"Go!"`
   - When the function ends, the deferred prints should fire in reverse order (LIFO): `from, from-1, ..., 1`
3. In `main()`:
   - Call `greet("Gopher")`
   - Print a blank line
   - Call `countdown(5)`

## Expected Output

```
Hello, Gopher!
How are you, Gopher?
Goodbye, Gopher!

Starting countdown...
Go!
5
4
3
2
1
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge06/main.go
```

## Hints

- `defer fmt.Println("Goodbye")` — the Println runs when the function returns, not when this line is reached
- In the countdown loop, each `defer fmt.Println(i)` captures the current value of `i` at the time of the defer call
- Stacking means: defer 1, defer 2, defer 3 → executes as 3, 2, 1
- The "Go!" line prints before the deferred numbers because defers run on function exit
- Think of `defer` like saying "remind me to do this before I leave"
