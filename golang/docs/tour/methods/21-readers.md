# Readers

## Go Concept

The `io` package specifies the **io.Reader** interface, which represents the read end of a stream of data. The Go standard library contains many implementations of this interface, including files, network connections, compressors, ciphers, and others.

The `io.Reader` interface has a single method:

```go
func (T) Read(b []byte) (n int, err error)
```

`Read` populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an `io.EOF` error when the stream ends.

The Reader interface is one of Go's most important interfaces, forming the basis for all I/O operations. Its simplicity enables powerful composition and allows any data source to be treated uniformly.

### Go Example

```go
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// strings.Reader implements io.Reader
	r := strings.NewReader("Hello, Reader!")

	// Read in small chunks
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// Output:
// n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
// b[:n] = "Hello, R"
// n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
// b[:n] = "eader!"
// n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
// b[:n] = ""

// Custom Reader example
type AlphabetReader struct {
	current byte
}

func (a *AlphabetReader) Read(b []byte) (int, error) {
	if a.current >= 'z' {
		return 0, io.EOF
	}
	
	n := 0
	for i := 0; i < len(b) && a.current < 'z'; i++ {
		a.current++
		b[i] = a.current
		n++
	}
	return n, nil
}

// Using helper functions
func demonstrateReaderUtilities() {
	r := strings.NewReader("Hello, World!")
	
	// io.ReadAll reads all data until EOF
	data, err := io.ReadAll(r)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(data)) // Hello, World!
	
	// io.Copy copies from reader to writer
	r2 := strings.NewReader("Copy this")
	written, err := io.Copy(io.Discard, r2)
	fmt.Printf("Copied %d bytes\n", written)
	
	// io.LimitReader limits how much can be read
	r3 := strings.NewReader("Limited reading")
	limited := io.LimitReader(r3, 7)
	data3, _ := io.ReadAll(limited)
	fmt.Println(string(data3)) // "Limited"
}
```

## C# Equivalent

C# provides the **Stream** abstract class and related types for I/O operations. The most similar interface to io.Reader is the abstract `Stream` class, though C# also has `IEnumerable<byte>` for sequence operations.

**Key C# I/O types:**
- **Stream**: Abstract base class for byte-oriented I/O (most similar to io.Reader)
- **StreamReader**: Text-oriented reading with encoding support
- **IAsyncEnumerable<byte>**: Async sequence of bytes (modern C#)

C# streams are more feature-rich than Go readers, supporting seeking, writing, and async operations in one abstraction. However, this makes them more complex and less composable than Go's focused Reader interface.

### C# Example

```csharp
using System;
using System.IO;
using System.Text;
using System.Threading.Tasks;

class Program
{
    static void Main()
    {
        // MemoryStream is similar to strings.Reader
        byte[] data = Encoding.UTF8.GetBytes("Hello, Reader!");
        using var stream = new MemoryStream(data);
        
        // Read in small chunks
        byte[] buffer = new byte[8];
        int bytesRead;
        while ((bytesRead = stream.Read(buffer, 0, buffer.Length)) > 0)
        {
            Console.WriteLine($"n = {bytesRead}");
            Console.WriteLine($"b = [{string.Join(", ", buffer)}]");
            Console.WriteLine($"b[:n] = \"{Encoding.UTF8.GetString(buffer, 0, bytesRead)}\"");
        }
    }
}

// Custom Stream implementation
class AlphabetStream : Stream
{
    private byte current = (byte)'a' - 1;
    
    public override int Read(byte[] buffer, int offset, int count)
    {
        if (current >= 'z')
            return 0; // End of stream
        
        int bytesRead = 0;
        for (int i = 0; i < count && current < 'z'; i++)
        {
            current++;
            buffer[offset + i] = current;
            bytesRead++;
        }
        return bytesRead;
    }
    
    // Required Stream members (read-only stream)
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

// Using Stream utilities
static async Task DemonstrateStreamUtilities()
{
    // Read all data from stream
    var data = Encoding.UTF8.GetBytes("Hello, World!");
    using var ms = new MemoryStream(data);
    using var reader = new StreamReader(ms);
    string content = await reader.ReadToEndAsync();
    Console.WriteLine(content); // Hello, World!
    
    // Copy stream to another
    var source = new MemoryStream(Encoding.UTF8.GetBytes("Copy this"));
    var destination = new MemoryStream();
    await source.CopyToAsync(destination);
    Console.WriteLine($"Copied {destination.Length} bytes");
    
    // Stream with modern async enumerable
    await foreach (var chunk in ReadChunksAsync(ms, 4))
    {
        Console.WriteLine($"Chunk: {Encoding.UTF8.GetString(chunk)}");
    }
}

static async IAsyncEnumerable<byte[]> ReadChunksAsync(Stream stream, int chunkSize)
{
    byte[] buffer = new byte[chunkSize];
    int bytesRead;
    while ((bytesRead = await stream.ReadAsync(buffer, 0, buffer.Length)) > 0)
    {
        byte[] chunk = new byte[bytesRead];
        Array.Copy(buffer, chunk, bytesRead);
        yield return chunk;
    }
}
```

## Key Differences

- **Abstraction**: Go Reader is a minimal interface (one method); C# Stream is an abstract class with many methods
- **Capabilities**: Go Reader is read-only; C# Stream can support read, write, seek
- **Error Handling**: Go returns (n, err); C# throws exceptions (or returns -1/0)
- **EOF Signaling**: Go uses io.EOF error; C# returns 0 bytes read
- **Async Support**: Go uses same interface (goroutines handle async); C# has separate async methods
- **Composition**: Go Reader composes easily with decorators; C# requires inheriting Stream
- **Type System**: Go uses interface; C# uses abstract class (can't implement multiple)
- **Simplicity**: Go Reader is simpler and easier to implement; C# Stream is more feature-rich
- **Memory**: Go provides byte slice to fill; C# takes buffer, offset, count parameters
- **Return Values**: Go returns bytes read and error; C# returns bytes read, throws on error
- **Partial Reads**: Both support partial reads; Go's error model makes this more explicit
- **Standard Library**: Both have extensive ecosystem of Reader/Stream implementations
- **Testing**: Go Reader easier to mock (interface); C# Stream requires more boilerplate
- **Modern Patterns**: C# added IAsyncEnumerable for sequence operations; Go uses channels for streaming
- **Performance**: Similar performance characteristics; both allow efficient chunk processing

## Common Patterns

**Go Reader Patterns:**
```go
// Read all at once
data, err := io.ReadAll(reader)

// Copy reader to writer
io.Copy(writer, reader)

// Buffered reading
bufReader := bufio.NewReader(reader)

// Chaining readers (decorators)
compressed := gzip.NewReader(file)
```

**C# Stream Patterns:**
```csharp
// Read all at once
using var reader = new StreamReader(stream);
string data = await reader.ReadToEndAsync();

// Copy stream to stream
await sourceStream.CopyToAsync(destStream);

// Buffered reading
using var buffered = new BufferedStream(stream);

// Chaining streams (decorators)
using var compressed = new GZipStream(file, CompressionMode.Decompress);
```

Both languages support similar patterns despite different underlying abstractions.
