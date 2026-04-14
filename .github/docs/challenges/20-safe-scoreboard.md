# Challenge 20: Safe Scoreboard

## Objective

Build a scoreboard that multiple goroutines update concurrently, using `sync.Mutex` to prevent race conditions on shared data.

## What You'll Learn

- Race conditions: what goes wrong when goroutines access shared data without synchronization
- `sync.Mutex` — mutual exclusion lock: `Lock()` and `Unlock()`
- Using `defer mu.Unlock()` to ensure the lock is always released
- Combining structs with a mutex for thread-safe data structures

## Tour Reference — Read These First

1. [sync.Mutex](https://go.dev/tour/concurrency/9) — Mutual exclusion for safe concurrent access to shared data

## What to Build

A `Scoreboard` struct that safely tracks scores for players, even when multiple goroutines update it at the same time.

## File to Create

```
challenge20/main.go
```

## Requirements

1. Define `type Scoreboard struct` with:
   - `mu sync.Mutex` — the mutex
   - `scores map[string]int` — player name to score
2. Write `func NewScoreboard() *Scoreboard` that creates and returns an initialized Scoreboard
3. Implement these methods on `*Scoreboard`:
   - `AddPoints(player string, points int)` — locks the mutex, adds points, unlocks
   - `GetScore(player string) int` — locks, reads the score, unlocks (use `defer` to unlock)
   - `GetAllScores() map[string]int` — returns a **copy** of the scores map (to avoid leaking the mutex-protected data)
4. In `main()`:
   - Create a scoreboard
   - Launch 3 goroutines, one for each player (`"Alice"`, `"Bob"`, `"Charlie"`)
   - Each goroutine adds 10 points, 100 times (total 1000 points per player)
   - Wait for all goroutines to finish (use `sync.WaitGroup`)
   - Print each player's final score — they should all be exactly 1000

## Expected Output

```
Starting 3 goroutines, each adding 10 points × 100 times...

Final scores:
  Alice: 1000
  Bob: 1000
  Charlie: 1000

All scores add up correctly!
```

> Without the mutex, you would get inconsistent/wrong totals due to race conditions.

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge20/main.go
```

You can also run with the race detector to verify there are no races:
```bash
go run -race challenge20/main.go
```

## Hints

- Mutex usage:
  ```go
  func (sb *Scoreboard) AddPoints(player string, points int) {
      sb.mu.Lock()
      sb.scores[player] += points
      sb.mu.Unlock()
  }
  ```
- Using defer:
  ```go
  func (sb *Scoreboard) GetScore(player string) int {
      sb.mu.Lock()
      defer sb.mu.Unlock()
      return sb.scores[player]
  }
  ```
- `sync.WaitGroup` for waiting:
  ```go
  var wg sync.WaitGroup
  wg.Add(3)
  go func() {
      defer wg.Done()
      // ... work ...
  }()
  wg.Wait()
  ```
- Import `"sync"` for both `sync.Mutex` and `sync.WaitGroup`
- To copy a map: create a new map and iterate over the original, copying each key-value pair
- `go run -race` is Go's built-in race condition detector — it will catch unsynchronized access
