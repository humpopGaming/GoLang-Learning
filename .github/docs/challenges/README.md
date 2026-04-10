# Go Tour Interactive Challenges

A series of 21 hands-on challenges that walk you through the [Tour of Go](http://127.0.0.1:3999/tour/list), building small runnable programs for each concept. These are designed to be **simpler than the Tour's own exercises** — the goal is to make you _use_ each concept so it sticks.

## How to Use

1. **Read the challenge file** — each one tells you what to build and links to the Tour pages to read first.
2. **Create the Go file** — each challenge tells you exactly which file to create (e.g. `challenge01/main.go`).
3. **Run it** — use `go run challenge01/main.go` from the `GoPlayground` root folder.
4. **Compare your output** — each challenge shows the expected output.
5. **Move on** — challenges build on each other conceptually, so do them in order.

> **Tip**: You can also open any challenge file in a **separate AI chat session** — it contains everything the AI needs to guide you through the implementation.

## Prerequisites

- Go installed (1.21+)
- VS Code with the Go extension
- The Tour of Go running locally at `http://127.0.0.1:3999`

---

## Phase 1: Basics — Packages, Variables, Functions

| #  | Challenge | File | Concepts |
|----|-----------|------|----------|
| 01 | [Hello Playground](01-hello-playground.md) | `challenge01/main.go` | Packages, imports, exported names |
| 02 | [Temperature Converter](02-temperature-converter.md) | `challenge02/main.go` | Functions, params, returns, types, conversions |
| 03 | [Currency Exchange](03-currency-exchange.md) | `challenge03/main.go` | Variables, constants, short declarations |

## Phase 2: Flow Control

| #  | Challenge | File | Concepts |
|----|-----------|------|----------|
| 04 | [FizzBuzz](04-fizzbuzz.md) | `challenge04/main.go` | `for` loops, `if`/`else` |
| 05 | [Grade Calculator](05-grade-calculator.md) | `challenge05/main.go` | `switch`, conditionless switch |
| 06 | [Countdown Timer](06-countdown-timer.md) | `challenge06/main.go` | `defer`, stacking defers |

## Phase 3: Composite Types

| #  | Challenge | File | Concepts |
|----|-----------|------|----------|
| 07 | [Pointer Swap](07-pointer-swap.md) | `challenge07/main.go` | Pointers, `&`, `*` |
| 08 | [Contact Card](08-contact-card.md) | `challenge08/main.go` | Structs, fields, struct literals |
| 09 | [Shopping List](09-shopping-list.md) | `challenge09/main.go` | Arrays, slices, append, range |
| 10 | [Phonebook](10-phonebook.md) | `challenge10/main.go` | Maps, function values, closures |

## Phase 4: Methods & Interfaces

| #  | Challenge | File | Concepts |
|----|-----------|------|----------|
| 11 | [Shape Calculator](11-shape-calculator.md) | `challenge11/main.go` | Methods, pointer/value receivers |
| 12 | [Animal Sounds](12-animal-sounds.md) | `challenge12/main.go` | Interfaces, Stringer |
| 13 | [Safe Divide](13-safe-divide.md) | `challenge13/main.go` | Errors, type assertions, type switches |
| 14 | [Byte Counter](14-byte-counter.md) | `challenge14/main.go` | `io.Reader` interface |

## Phase 5: Generics

| #  | Challenge | File | Concepts |
|----|-----------|------|----------|
| 15 | [Generic Helpers](15-generic-helpers.md) | `challenge15/main.go` | Generic functions, `comparable` |
| 16 | [Generic Stack](16-generic-stack.md) | `challenge16/main.go` | Generic types |

## Phase 6: Concurrency

| #  | Challenge | File | Concepts |
|----|-----------|------|----------|
| 17 | [Parallel Greeter](17-parallel-greeter.md) | `challenge17/main.go` | Goroutines |
| 18 | [Pipeline](18-pipeline.md) | `challenge18/main.go` | Channels, buffered channels, range/close |
| 19 | [Timeout Racer](19-timeout-racer.md) | `challenge19/main.go` | `select`, default selection |
| 20 | [Safe Scoreboard](20-safe-scoreboard.md) | `challenge20/main.go` | `sync.Mutex` |

## Capstone

| #  | Challenge | File | Concepts |
|----|-----------|------|----------|
| 21 | [Todo CLI](21-todo-cli.md) | `challenge21/main.go` | Everything combined |
