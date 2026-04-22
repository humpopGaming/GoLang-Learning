# Exercise: Readers

## Challenge

Implement a `Reader` type that emits an infinite stream of the ASCII character `'A'`.

## Starter Code

```go
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
```

## Expected Output

```
OK!
```

## Hints

- The `Read` method should fill the byte slice with the character `'A'`
- Return the number of bytes written (which is `len(b)`)
- Return `nil` for the error (the stream is infinite, so no EOF)
- Remember: `'A'` is ASCII value 65

## Solution

<details>
<summary>Click to reveal solution</summary>

```go
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
```

**Explanation:**
- Loop through the byte slice `b`
- Set each byte to 'A' (ASCII 65)
- Return the number of bytes filled (`len(b)`)
- Return `nil` for error since this reader never ends
- The `reader.Validate` function tests your implementation

</details>

## C# Equivalent Exercise

```csharp
using System;
using System.IO;

// Implement a Stream that emits infinite 'A' characters
class MyStream : Stream
{
    // TODO: Implement Read method
    public override int Read(byte[] buffer, int offset, int count)
    {
        // Fill buffer with 'A' (ASCII 65)
        return 0;
    }
    
    // Required Stream properties (read-only stream)
    public override bool CanRead => true;
    public override bool CanSeek => false;
    public override bool CanWrite => false;
    public override long Length => throw new NotSupportedException();
    public override long Position 
    { 
        get => throw new NotSupportedException(); 
        set => throw new NotSupportedException(); 
    }
    
    public override void Flush() { }
    public override long Seek(long offset, SeekOrigin origin) => throw new NotSupportedException();
    public override void SetLength(long value) => throw new NotSupportedException();
    public override void Write(byte[] buffer, int offset, int count) => throw new NotSupportedException();
}

class Program
{
    static void Main()
    {
        using var stream = new MyStream();
        byte[] buffer = new byte[10];
        int bytesRead = stream.Read(buffer, 0, buffer.Length);
        
        Console.WriteLine($"Read {bytesRead} bytes");
        Console.WriteLine($"Content: {System.Text.Encoding.ASCII.GetString(buffer)}");
        // Should print: AAAAAAAAAA
    }
}
```

**C# Solution:**

```csharp
public override int Read(byte[] buffer, int offset, int count)
{
    for (int i = 0; i < count; i++)
    {
        buffer[offset + i] = (byte)'A';
    }
    return count;
}
```

**Alternative: Using Span<byte> (modern C#):**

```csharp
public override int Read(Span<byte> buffer)
{
    buffer.Fill((byte)'A');
    return buffer.Length;
}
```

## Extension Challenges

1. **Alphabet Reader**: Create a reader that cycles through 'A' to 'Z' repeatedly

2. **Configurable Reader**: Accept a character in the constructor and emit that character

3. **Counting Reader**: Emit characters and track how many bytes have been read

4. **Limited Reader**: Add a maximum number of bytes to emit before returning EOF

5. **Pattern Reader**: Emit a pattern like "ABC" repeatedly

## Testing Your Reader

```go
package main

import (
	"fmt"
	"io"
)

func testReader(r io.Reader) {
	b := make([]byte, 8)
	for i := 0; i < 3; i++ {
		n, err := r.Read(b)
		fmt.Printf("Read %d bytes: %s\n", n, b[:n])
		if err == io.EOF {
			break
		}
	}
}

func main() {
	r := MyReader{}
	testReader(r)
	// Output:
	// Read 8 bytes: AAAAAAAA
	// Read 8 bytes: AAAAAAAA
	// Read 8 bytes: AAAAAAAA
}
```

## Key Learning Points

- `io.Reader` interface requires only one method: `Read([]byte) (int, error)`
- The Reader fills the provided slice; it doesn't allocate
- Return value is (bytes written, error)
- `io.EOF` signals end of stream
- Infinite readers return `nil` error
- The same pattern works for files, networks, compression, encryption, etc.
- C# Stream requires implementing many more methods, even if unused
- Go's minimal interface makes it easy to create custom readers
- Reader composability enables powerful data processing pipelines
