# Exercise: rot13Reader

## Challenge

A common pattern is an `io.Reader` that wraps another `io.Reader`, modifying the stream in some way.

Implement a `rot13Reader` that implements `io.Reader` and reads from an `io.Reader`, applying the ROT13 substitution cipher to the stream.

The `rot13Reader` type is provided for you. Make it an `io.Reader` by implementing its `Read` method.

## Background: ROT13

ROT13 is a simple letter substitution cipher that replaces a letter with the letter 13 positions after it in the alphabet. It's a special case of the Caesar cipher.

- A ↔ N, B ↔ O, C ↔ P, ..., M ↔ Z
- Only letters are transformed; numbers and punctuation remain unchanged

Example: `"Hello"` becomes `"Uryyb"`

## Starter Code

```go
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	// TODO: Read from rot.r, apply ROT13, write to b
	return 0, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
```

## Expected Output

```
You cracked the code!
```

## Hints

- Call `rot.r.Read(b)` to read from the underlying reader
- Transform each byte in place after reading
- For letters A-Z: `(char - 'A' + 13) % 26 + 'A'`
- For letters a-z: `(char - 'a' + 13) % 26 + 'a'`
- Leave non-letter characters unchanged
- Return the same values (n, err) that the underlying reader returned

## Solution

<details>
<summary>Click to reveal solution</summary>

```go
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'Z':
		return 'A' + (b-'A'+13)%26
	case 'a' <= b && b <= 'z':
		return 'a' + (b-'a'+13)%26
	default:
		return b
	}
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r) // Output: You cracked the code!
}
```

**Explanation:**
1. Read from the underlying reader into byte slice `b`
2. Transform each byte using ROT13 cipher
3. Return the same (n, err) from the underlying reader
4. This is the **decorator pattern** - wrapping a reader to add behavior

</details>

## C# Equivalent Exercise

```csharp
using System;
using System.IO;
using System.Text;

// Implement a Stream that wraps another Stream and applies ROT13
class Rot13Stream : Stream
{
    private readonly Stream innerStream;
    
    public Rot13Stream(Stream stream)
    {
        innerStream = stream;
    }
    
    private static byte Rot13(byte b)
    {
        // TODO: Implement ROT13 transformation
        return b;
    }
    
    public override int Read(byte[] buffer, int offset, int count)
    {
        // TODO: Read from innerStream, apply ROT13, write to buffer
        return 0;
    }
    
    // Required Stream members
    public override bool CanRead => innerStream.CanRead;
    public override bool CanSeek => false;
    public override bool CanWrite => false;
    public override long Length => throw new NotSupportedException();
    public override long Position 
    { 
        get => throw new NotSupportedException(); 
        set => throw new NotSupportedException(); 
    }
    
    public override void Flush() => innerStream.Flush();
    public override long Seek(long offset, SeekOrigin origin) => throw new NotSupportedException();
    public override void SetLength(long value) => throw new NotSupportedException();
    public override void Write(byte[] buffer, int offset, int count) => throw new NotSupportedException();
    
    protected override void Dispose(bool disposing)
    {
        if (disposing)
            innerStream?.Dispose();
        base.Dispose(disposing);
    }
}

class Program
{
    static void Main()
    {
        byte[] data = Encoding.ASCII.GetBytes("Lbh penpxrq gur pbqr!");
        using var source = new MemoryStream(data);
        using var rot13 = new Rot13Stream(source);
        using var reader = new StreamReader(rot13);
        
        Console.WriteLine(reader.ReadToEnd());
        // Should output: You cracked the code!
    }
}
```

**C# Solution:**

```csharp
private static byte Rot13(byte b)
{
    if (b >= 'A' && b <= 'Z')
        return (byte)('A' + (b - 'A' + 13) % 26);
    if (b >= 'a' && b <= 'z')
        return (byte)('a' + (b - 'a' + 13) % 26);
    return b;
}

public override int Read(byte[] buffer, int offset, int count)
{
    int bytesRead = innerStream.Read(buffer, offset, count);
    for (int i = 0; i < bytesRead; i++)
    {
        buffer[offset + i] = Rot13(buffer[offset + i]);
    }
    return bytesRead;
}
```

## Extension Challenges

1. **Reversible**: Make ROT13 work both for encoding and decoding (it's symmetric!)

2. **Caesar Cipher**: Generalize to accept any shift value, not just 13

3. **Write Support**: Add a Write method that also applies ROT13

4. **Buffering**: Add buffering to improve performance for small reads

5. **Other Transforms**: Implement other transformations (uppercase, reverse, etc.)

6. **Chaining**: Chain multiple transforming readers together

## Real-World Uses of Reader Decorators

```go
// Compression
compressed := gzip.NewReader(file)

// Encryption  
encrypted := cipher.StreamReader{S: stream, R: file}

// Buffering
buffered := bufio.NewReader(file)

// Limiting
limited := io.LimitReader(file, 1024)

// Chaining multiple decorators
reader := bufio.NewReader(
    gzip.NewReader(
        cipher.StreamReader{S: stream, R: file}))
```

```csharp
// C# equivalent
using var file = File.OpenRead("data.txt");
using var gzip = new GZipStream(file, CompressionMode.Decompress);
using var crypto = new CryptoStream(gzip, transform, CryptoStreamMode.Read);
using var buffered = new BufferedStream(crypto);
```

## Key Learning Points

- **Decorator Pattern**: Wrap one reader to add behavior without modifying it
- **Composability**: Readers can be chained together easily
- **Transparency**: The decorator implements the same interface as what it wraps
- **Single Responsibility**: Each reader does one thing (decompress, decrypt, transform)
- **Standard Library**: Many standard readers use this pattern
- **Testing**: Easy to test by using strings.Reader for input
- Both Go and C# support this pattern, though C# requires more boilerplate
- The pattern enables building complex processing pipelines from simple components
- This is a core design pattern in Go's I/O system
