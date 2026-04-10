# Challenge Index

Quick-reference mapping of challenge numbers to files, concepts, and Tour pages.

## By Challenge Number

| # | Name | File | Key Concepts | Tour Pages |
|---|------|------|-------------|------------|
| 01 | Hello Playground | [01-hello-playground.md](../../docs/challenges/01-hello-playground.md) | Packages, imports, exported names | basics/1-3 |
| 02 | Temperature Converter | [02-temperature-converter.md](../../docs/challenges/02-temperature-converter.md) | Functions, multiple returns, types, conversions | basics/4-7, 11-14 |
| 03 | Currency Exchange | [03-currency-exchange.md](../../docs/challenges/03-currency-exchange.md) | Variables, constants, short declarations | basics/8-10, 15-16 |
| 04 | FizzBuzz | [04-fizzbuzz.md](../../docs/challenges/04-fizzbuzz.md) | for loops, if/else, if-short-statement | flowcontrol/1-3, 5-7 |
| 05 | Grade Calculator | [05-grade-calculator.md](../../docs/challenges/05-grade-calculator.md) | switch, conditionless switch | flowcontrol/9-11 |
| 06 | Countdown Timer | [06-countdown-timer.md](../../docs/challenges/06-countdown-timer.md) | defer, stacking defers | flowcontrol/12-13 |
| 07 | Pointer Swap | [07-pointer-swap.md](../../docs/challenges/07-pointer-swap.md) | Pointers, &, *, dereferencing | moretypes/1 |
| 08 | Contact Card | [08-contact-card.md](../../docs/challenges/08-contact-card.md) | Structs, fields, struct literals, pointers to structs | moretypes/2-5 |
| 09 | Shopping List | [09-shopping-list.md](../../docs/challenges/09-shopping-list.md) | Arrays, slices, append, range, make, len/cap | moretypes/6-17 |
| 10 | Phonebook | [10-phonebook.md](../../docs/challenges/10-phonebook.md) | Maps, map literals, mutating maps, closures | moretypes/19-26 |
| 11 | Shape Calculator | [11-shape-calculator.md](../../docs/challenges/11-shape-calculator.md) | Methods, pointer/value receivers | methods/1-8 |
| 12 | Animal Sounds | [12-animal-sounds.md](../../docs/challenges/12-animal-sounds.md) | Interfaces, implicit implementation, Stringer | methods/9-14, 17 |
| 13 | Safe Divide | [13-safe-divide.md](../../docs/challenges/13-safe-divide.md) | error interface, custom errors, type assertions, type switches | methods/15-16, 19-20 |
| 14 | Byte Counter | [14-byte-counter.md](../../docs/challenges/14-byte-counter.md) | io.Reader, Read method, wrapping readers | methods/21-23 |
| 15 | Generic Helpers | [15-generic-helpers.md](../../docs/challenges/15-generic-helpers.md) | Generic functions, comparable, type parameters | generics/1 |
| 16 | Generic Stack | [16-generic-stack.md](../../docs/challenges/16-generic-stack.md) | Generic types, type parameters on structs | generics/2 |
| 17 | Parallel Greeter | [17-parallel-greeter.md](../../docs/challenges/17-parallel-greeter.md) | Goroutines, go keyword | concurrency/1 |
| 18 | Pipeline | [18-pipeline.md](../../docs/challenges/18-pipeline.md) | Channels, buffered channels, range/close | concurrency/2-4 |
| 19 | Timeout Racer | [19-timeout-racer.md](../../docs/challenges/19-timeout-racer.md) | select, default selection, time.After | concurrency/5-6 |
| 20 | Safe Scoreboard | [20-safe-scoreboard.md](../../docs/challenges/20-safe-scoreboard.md) | sync.Mutex, WaitGroup, race conditions | concurrency/9 |
| 21 | Todo CLI | [21-todo-cli.md](../../docs/challenges/21-todo-cli.md) | Capstone — all concepts combined | All |

## By Concept

Use this to find which challenge covers the concept the learner is asking about.

| Concept | Challenge(s) |
|---------|-------------|
| `package main`, imports | 01 |
| Exported names (capital letters) | 01 |
| Functions, parameters, return values | 02 |
| Multiple return values | 02 |
| Named return values | 02 |
| Basic types (int, float64, bool, string) | 02 |
| Zero values | 02, 08 |
| Type conversions | 02 |
| Type inference (`:=`) | 02, 03 |
| `var` declarations | 03 |
| Short variable declarations (`:=`) | 03 |
| Constants (`const`) | 03 |
| Numeric constants | 03 |
| `for` loop (standard) | 04 |
| `for` as while loop | 04 |
| `if`/`else` | 04 |
| `if` with short statement | 04 |
| `switch` | 05 |
| Conditionless switch | 05 |
| `defer` | 06 |
| Stacking defers (LIFO) | 06 |
| Pointers (`*`, `&`) | 07 |
| Structs | 08 |
| Struct fields, dot notation | 08 |
| Pointers to structs | 08 |
| Struct literals | 08 |
| Arrays | 09 |
| Slices | 09 |
| Slice literals, defaults | 09 |
| `len()`, `cap()` | 09 |
| Nil slices | 09 |
| `make()` for slices | 09 |
| `append()` | 09 |
| `range` | 09, 10 |
| Maps | 10 |
| Map literals | 10 |
| Mutating maps (add/update/delete/exists) | 10 |
| Function values | 10 |
| Closures | 10, 21 |
| Methods | 11 |
| Value receivers | 11 |
| Pointer receivers | 11 |
| Method indirection | 11 |
| Interfaces | 12 |
| Implicit interface implementation | 12 |
| Interface values | 12 |
| Empty interface (`interface{}`/`any`) | 12 |
| `fmt.Stringer` | 12, 21 |
| Type assertions | 13 |
| Type switches | 13 |
| `error` interface | 13, 21 |
| Custom error types | 13, 21 |
| `io.Reader` | 14 |
| Generic functions | 15 |
| `comparable` constraint | 15 |
| Generic types | 16 |
| Goroutines | 17 |
| Channels | 18 |
| Buffered channels | 18 |
| `close()`, `range` on channels | 18 |
| `select` | 19 |
| Default selection | 19 |
| `time.After` | 19 |
| `sync.Mutex` | 20, 21 |
| `sync.WaitGroup` | 20 |
| Race conditions | 20 |
