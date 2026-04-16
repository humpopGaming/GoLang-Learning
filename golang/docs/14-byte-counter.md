# Challenge 14: Byte Counter

## Objective

Implement a custom `io.Reader` that counts how many bytes it has read, learning the `io.Reader` interface pattern — one of the most important interfaces in Go.

## What You'll Learn

- The `io.Reader` interface: `Read(p []byte) (n int, err error)`
- How `Read` works: it fills a byte slice and returns how many bytes were read
- `io.EOF` signals the end of a stream
- Wrapping one Reader around another (the "decorator" pattern)
- Using `strings.NewReader` to create a Reader from a string

## Tour Reference — Read These First

1. [Readers](https://go.dev/tour/methods/21) — The `io.Reader` interface and how `Read` works
2. [Exercise: Readers](https://go.dev/tour/methods/22) — Implementing a Reader that emits data
3. [Exercise: rot13Reader](https://go.dev/tour/methods/23) — Wrapping a Reader to transform data

## What to Build

A `CountingReader` that wraps any `io.Reader` and counts the total number of bytes read through it. Then use it to count bytes while reading a string.

## File to Create

```
challenge14/main.go
```

## Requirements

1. Define `type CountingReader struct` with two fields:
   - `Reader io.Reader` — the wrapped reader
   - `BytesRead int` — running total of bytes read
2. Implement the `Read(p []byte) (int, error)` method on `*CountingReader`:
   - Call the wrapped `Reader.Read(p)`
   - Add the returned `n` to `BytesRead`
   - Return the same `n` and `err` from the wrapped reader
3. In `main()`:
   - Create a `strings.NewReader` with the text `"Hello, Go Reader interface!"`
   - Wrap it in a `CountingReader`
   - Read from the `CountingReader` in a loop using a small buffer (e.g. 8 bytes), printing what you read each time
   - When `io.EOF` is reached, stop and print the total bytes read

## Expected Output

```
Read 8 bytes: "Hello, G"
Read 8 bytes: "o Reader"
Read 8 bytes: " interfa"
Read 3 bytes: "ce!"
Read 0 bytes: "" (EOF)

Total bytes read: 27
```

## How to Run

```bash
cd c:\repos\GoPlayground
go run challenge14/main.go
```

## Hints

- Your struct:
  ```go
  type CountingReader struct {
      Reader    io.Reader
      BytesRead int
  }
  ```
- The Read method:
  ```go
  func (cr *CountingReader) Read(p []byte) (int, error) {
      n, err := cr.Reader.Read(p)
      cr.BytesRead += n
      return n, err
  }
  ```
- Reading loop pattern:
  ```go
  buf := make([]byte, 8)
  for {
      n, err := cr.Read(buf)
      fmt.Printf("Read %d bytes: %q\n", n, buf[:n])
      if err == io.EOF {
          break
      }
  }
  ```
- Import `"io"`, `"strings"`, and `"fmt"`
- Use `buf[:n]` to only look at the bytes that were actually read
