# Challenge 19: Timeout Racer

## Objective

Build a program that races multiple slow operations and picks the result from whichever finishes first, using `select` to wait on multiple channels simultaneously.

## What You'll Learn

- `select` lets a goroutine wait on multiple channel operations
- The first ready case runs; if multiple are ready, one is chosen at random
- `default` case runs if no other case is ready (non-blocking receive/send)
- Using `time.After` for timeouts
- Combining goroutines, channels, and select for real concurrent patterns

## Tour Reference — Read These First

1. [Select](http://127.0.0.1:3999/tour/concurrency/5) — Wait on multiple channels; first ready case wins
2. [Default Selection](http://127.0.0.1:3999/tour/concurrency/6) — `default` makes a select non-blocking

## What to Build

A program that:
1. Simulates two "servers" that respond at different speeds
2. Races them — whoever responds first wins
3. Demonstrates a timeout using `time.After`
4. Shows a non-blocking channel check using `default`

## File to Create

```
challenge19/main.go
```

## Requirements

1. Write `func slowServer(name string, delay time.Duration, ch chan<- string)` that sleeps for `delay` then sends `"Response from <name>"` to `ch`
2. Write `func race()` that:
   - Creates two channels
   - Starts two `slowServer` goroutines: `"Server A"` with 200ms delay and `"Server B"` with 100ms delay
   - Uses `select` to receive from whichever channel responds first
   - Prints the winner
3. Write `func raceWithTimeout()` that:
   - Does the same as `race()` but both servers are slow (500ms and 800ms)
   - Adds a `time.After(300 * time.Millisecond)` as a timeout case in the `select`
   - If the timeout fires first, prints `"Timed out! No server responded in time."`
4. Write `func nonBlockingCheck()` that:
   - Creates a channel
   - Uses `select` with a `default` case to try receiving without blocking
   - Prints `"No value ready"` from the default case
   - Then sends a value to a **buffered channel** and tries again — this time it should receive the value
5. Call all three functions from `main()`

## Expected Output

```
--- Race ---
Winner: Response from Server B

--- Race with Timeout ---
Timed out! No server responded in time.

--- Non-blocking Check ---
No value ready
Got value: 42
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge19/main.go
```

## Hints

- Select syntax:
  ```go
  select {
  case msg := <-ch1:
      fmt.Println(msg)
  case msg := <-ch2:
      fmt.Println(msg)
  case <-time.After(300 * time.Millisecond):
      fmt.Println("Timeout!")
  }
  ```
- `time.After(d)` returns a channel that receives a value after duration `d`
- Non-blocking check with default:
  ```go
  select {
  case v := <-ch:
      fmt.Println("Got:", v)
  default:
      fmt.Println("No value ready")
  }
  ```
- Import `"time"` for `time.Duration`, `time.Millisecond`, `time.After`, `time.Sleep`
- Server B wins the first race because it has a shorter delay (100ms vs 200ms)
