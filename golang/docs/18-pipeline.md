# Challenge 18: Pipeline

## Objective

Build a number-processing pipeline using channels — goroutines connected by channels, where each stage takes input from one channel and sends output to another.

## What You'll Learn

- Channels: typed conduits for sending and receiving values between goroutines
- Creating channels with `make(chan T)`
- Sending to a channel: `ch <- value`
- Receiving from a channel: `value := <-ch`
- Buffered channels: `make(chan T, bufferSize)`
- Closing a channel with `close(ch)`
- Receiving with `for v := range ch` — loops until the channel is closed

## Tour Reference — Read These First

1. [Channels](tour/concurrency/02-channels.html) — Create, send, receive; channels synchronize goroutines
2. [Buffered Channels](tour/concurrency/03-buffered-channels.html) — `make(chan int, 100)` — sends block when buffer is full
3. [Range and Close](tour/concurrency/04-range-and-close.html) — `close(ch)` signals no more values; `for v := range ch` drains it

## What to Build

A three-stage pipeline:
1. **Generator**: produces numbers 1–10 and sends them into a channel
2. **Doubler**: reads numbers from one channel, doubles them, sends to another channel
3. **Printer**: reads final numbers from a channel and prints them

Each stage runs in its own goroutine. Channels connect them.

## File to Create

```
challenge18/main.go
```

## Requirements

1. Write `func generate(out chan<- int)` that sends integers 1 through 10 into `out`, then closes the channel
2. Write `func double(in <-chan int, out chan<- int)` that reads each value from `in`, doubles it, and sends the result to `out`, then closes `out`
3. Write `func print(in <-chan int, done chan<- bool)` that reads each value from `in` and prints it, then sends `true` to `done` when finished
4. In `main()`:
   - Create three channels: `numbers`, `doubled`, and `done`
   - Start `generate`, `double`, and `print` as goroutines, wiring the channels together
   - Wait for `done` to receive a value before exiting
5. **Bonus**: Also demonstrate a **buffered channel** — create one with capacity 3, send 3 values without a receiver, then receive all 3

## Expected Output

```
Pipeline output:
2
4
6
8
10
12
14
16
18
20

Buffered channel demo:
Sent 3 values to a buffered channel (capacity 3) without blocking
Received: 100
Received: 200
Received: 300
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge18/main.go
```

## Hints

- Channel direction syntax:
  - `chan<- int` — send-only channel
  - `<-chan int` — receive-only channel
  - `chan int` — bidirectional
- Generator pattern:
  ```go
  func generate(out chan<- int) {
      for i := 1; i <= 10; i++ {
          out <- i
      }
      close(out)
  }
  ```
- Reading until closed: `for v := range in { ... }`
- Only the **sender** should close a channel, never the receiver
- Buffered channel: `ch := make(chan int, 3)` can hold 3 values before blocking
- The `done` channel is a simple synchronization trick — `main` blocks on `<-done` until the print stage finishes
