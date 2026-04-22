# sync.Mutex

## Go Concept

**Mutual exclusion** (mutex) is used to protect shared data from concurrent access. While Go encourages communication through channels, sometimes you need to ensure that only one goroutine can access a variable at a time to avoid conflicts.

Go's standard library provides `sync.Mutex` with two methods:
- `Lock()` - Acquire the mutex (blocks if already locked)
- `Unlock()` - Release the mutex

### When to Use Mutex vs Channels

**Use channels when:**
- Communicating data between goroutines
- Coordinating goroutine execution
- Implementing pipelines or workflows

**Use mutex when:**
- Protecting shared state (counters, caches, maps)
- Multiple goroutines need to read/modify the same data
- The data structure itself is the important thing (not the communication)

Go philosophy: "Don't communicate by sharing memory; share memory by communicating." However, sometimes mutexes are the simpler solution.

### Go Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	
	// Launch 1000 goroutines that increment the same key
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey")) // Output: 1000
}
```

### sync.RWMutex

Go also provides `sync.RWMutex` for **reader-writer locks**:
- Multiple readers can hold the lock simultaneously
- Only one writer can hold the lock
- Writers block both readers and other writers

```go
type Cache struct {
	mu    sync.RWMutex
	items map[string]string
}

func (c *Cache) Get(key string) string {
	c.mu.RLock()         // Read lock
	defer c.mu.RUnlock()
	return c.items[key]
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()          // Write lock
	defer c.mu.Unlock()
	c.items[key] = value
}
```

## C# Equivalent

C# provides several synchronization primitives:
- **`lock` keyword** - Most common, similar to mutex
- **`Monitor`** - Lower-level than lock
- **`SemaphoreSlim`** - More flexible
- **`ReaderWriterLockSlim`** - Similar to RWMutex

### **Paradigm Context:**

**Go:**
- Mutexes are **straightforward** - lock/unlock
- Used sparingly; channels are preferred
- Goroutines are so cheap that blocking is fine

**C#:**
- Multiple locking mechanisms available
- More common than channels
- Must be careful with async code (lock doesn't work with await)
- `SemaphoreSlim` needed for async scenarios

### C# Example (using lock)

```csharp
using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;

class SafeCounter
{
    private readonly object _lock = new object();
    private readonly Dictionary<string, int> _values = new Dictionary<string, int>();

    public void Inc(string key)
    {
        lock (_lock) // Lock so only one thread can access at a time
        {
            if (!_values.ContainsKey(key))
                _values[key] = 0;
            _values[key]++;
        }
    }

    public int Value(string key)
    {
        lock (_lock)
        {
            return _values.ContainsKey(key) ? _values[key] : 0;
        }
    }
}

class Program
{
    static async Task Main()
    {
        var counter = new SafeCounter();
        
        // Launch 1000 tasks that increment the same key
        var tasks = new Task[1000];
        for (int i = 0; i < 1000; i++)
        {
            tasks[i] = Task.Run(() => counter.Inc("somekey"));
        }
        
        await Task.WhenAll(tasks);
        
        Console.WriteLine(counter.Value("somekey")); // Output: 1000
    }
}
```

### C# Example (using ReaderWriterLockSlim)

```csharp
using System;
using System.Collections.Generic;
using System.Threading;

class Cache
{
    private readonly ReaderWriterLockSlim _lock = new ReaderWriterLockSlim();
    private readonly Dictionary<string, string> _items = new Dictionary<string, string>();

    public string Get(string key)
    {
        _lock.EnterReadLock();
        try
        {
            return _items.ContainsKey(key) ? _items[key] : null;
        }
        finally
        {
            _lock.ExitReadLock();
        }
    }

    public void Set(string key, string value)
    {
        _lock.EnterWriteLock();
        try
        {
            _items[key] = value;
        }
        finally
        {
            _lock.ExitWriteLock();
        }
    }
}
```

### C# Async Example (using SemaphoreSlim)

The `lock` keyword doesn't work with async/await. Use `SemaphoreSlim` instead:

```csharp
using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;

class AsyncSafeCounter
{
    private readonly SemaphoreSlim _semaphore = new SemaphoreSlim(1, 1);
    private readonly Dictionary<string, int> _values = new Dictionary<string, int>();

    public async Task IncAsync(string key)
    {
        await _semaphore.WaitAsync(); // Like Lock()
        try
        {
            if (!_values.ContainsKey(key))
                _values[key] = 0;
            _values[key]++;
        }
        finally
        {
            _semaphore.Release(); // Like Unlock()
        }
    }

    public async Task<int> ValueAsync(string key)
    {
        await _semaphore.WaitAsync();
        try
        {
            return _values.ContainsKey(key) ? _values[key] : 0;
        }
        finally
        {
            _semaphore.Release();
        }
    }
}
```

## Key Comparison

| Feature | Go (sync.Mutex) | C# (lock) | C# (SemaphoreSlim) |
|---------|-----------------|-----------|-------------------|
| **Basic syntax** | `mu.Lock()` / `mu.Unlock()` | `lock (obj) { }` | `await sem.WaitAsync()` / `sem.Release()` |
| **Async support** | N/A (blocking is cheap) | No | Yes |
| **Defer unlock** | `defer mu.Unlock()` | Automatic with `lock` | Use try/finally |
| **Read-write** | `sync.RWMutex` | `ReaderWriterLockSlim` | N/A |
| **Philosophy** | Use channels when possible | Use locks commonly | Use for async code |

## Comparison of C# Locking Mechanisms

| Mechanism | Use When | Async Support | Performance |
|-----------|----------|---------------|-------------|
| **lock keyword** | Simple thread synchronization | No | Fast |
| **Monitor** | Need fine control (Wait/Pulse) | No | Fast |
| **SemaphoreSlim** | Async code or limiting concurrency | Yes | Medium |
| **ReaderWriterLockSlim** | Many reads, few writes | No | Fast for reads |
| **Concurrent collections** | Simple scenarios | Implicit | Very fast |

## Common Patterns

### Pattern 1: Defer Unlock (Go)
```go
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock() // Ensures unlock even if panic
	return c.v[key]
}
```

**C# Equivalent:**
```csharp
public int Value(string key)
{
    lock (_lock) // Automatically unlocks
    {
        return _values[key];
    }
}
```

### Pattern 2: Async Locking (C# only)
```csharp
public async Task<int> IncrementAsync(string key)
{
    await _semaphore.WaitAsync();
    try
    {
        // Critical section
        return ++_values[key];
    }
    finally
    {
        _semaphore.Release();
    }
}
```

### Pattern 3: Concurrent Collections (C# Alternative)

C# has **concurrent collections** that handle locking internally:

```csharp
using System.Collections.Concurrent;

class SafeCounter
{
    private readonly ConcurrentDictionary<string, int> _values = 
        new ConcurrentDictionary<string, int>();

    public void Inc(string key)
    {
        _values.AddOrUpdate(key, 1, (k, v) => v + 1);
    }

    public int Value(string key)
    {
        return _values.GetValueOrDefault(key);
    }
}
```

## Race Condition Example

Without mutex (Go):
```go
// UNSAFE - race condition!
type UnsafeCounter struct {
	v map[string]int
}

func (c *UnsafeCounter) Inc(key string) {
	c.v[key]++ // Multiple goroutines can corrupt this!
}
```

Run with `go run -race` to detect race conditions.

Without lock (C#):
```csharp
// UNSAFE - race condition!
class UnsafeCounter
{
    private Dictionary<string, int> _values = new Dictionary<string, int>();

    public void Inc(string key)
    {
        _values[key]++; // Multiple threads can corrupt this!
    }
}
```

## When to Use What

### Go Decision Tree:
1. **Need to communicate?** → Use channels
2. **Protecting shared state?** → Use mutex
3. **Many reads, few writes?** → Use RWMutex
4. **Simple counter?** → Use atomic operations (sync/atomic)

### C# Decision Tree:
1. **Async code?** → Use SemaphoreSlim or async-safe patterns
2. **Simple sync code?** → Use lock keyword
3. **Many reads, few writes?** → Use ReaderWriterLockSlim
4. **Simple collections?** → Use ConcurrentDictionary or concurrent collections
5. **Complex coordination?** → Consider Channel<T> or async patterns

## Best Practices

**Go:**
- Prefer channels for communication
- Use mutexes for protecting shared state
- Always use `defer` to unlock
- Keep critical sections small
- Run with `-race` flag during development

**C#:**
- Use lock for simple synchronization
- Use SemaphoreSlim with async code
- Never use `lock` with `await` inside
- Consider concurrent collections first
- Keep critical sections small
- Avoid nested locks (deadlock risk)
