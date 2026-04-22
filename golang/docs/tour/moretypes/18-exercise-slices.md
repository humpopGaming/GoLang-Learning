# Exercise: Slices

## Challenge

Implement the `Pic` function that returns a slice of slices of `uint8` representing a 2D image. The function should allocate a `[][]uint8` with `dy` rows, each row having `dx` elements of type `uint8`.

Each integer value represents a grayscale pixel (0 = black, 255 = white). Choose your own formula to generate interesting patterns.

Possible formulas to try:
- `(x+y)/2`
- `x*y`
- `x^y` (XOR)
- `x*log(y)` (requires math package)
- `x%(y+1)`

### Requirements

1. Create a 2D slice with dimensions `dx` by `dy`
2. Fill each pixel with a computed value based on x and y coordinates
3. Return the completed 2D slice
4. The provided test harness will display your image

### Go Template

```go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// TODO: Implement this function
	// 1. Create outer slice of length dy
	// 2. For each row, create inner slice of length dx
	// 3. For each pixel at (x,y), compute a value
	// 4. Return the 2D slice
}

func main() {
	pic.Show(Pic)
}
```

## Key Concepts to Practice

This exercise reinforces several slice concepts from the Go Tour:

1. **Creating slices with make**: Use `make([][]uint8, dy)` for outer slice
2. **Slices of slices**: Creating and working with 2D slice structures
3. **Slice initialization**: Creating each row with `make([]uint8, dx)`
4. **Range iteration**: Using `for i := range` to iterate with index
5. **Nested loops**: Processing 2D structures with nested range loops

## C# Equivalent Exercise

In C#, you would implement this using jagged arrays or List of Lists:

### C# Template

```csharp
using System;
using System.Drawing;
using System.Drawing.Imaging;

class Program
{
    static byte[][] Pic(int dx, int dy)
    {
        // TODO: Implement this function
        // 1. Create outer array of length dy
        // 2. For each row, create inner array of length dx
        // 3. For each pixel at (x,y), compute a value
        // 4. Return the jagged array
        
        return null;
    }

    static void Main()
    {
        var image = Pic(256, 256);
        DisplayImage(image);
    }

    static void DisplayImage(byte[][] pixels)
    {
        int height = pixels.Length;
        int width = pixels[0].Length;
        
        Bitmap bitmap = new Bitmap(width, height);
        for (int y = 0; y < height; y++)
        {
            for (int x = 0; x < width; x++)
            {
                byte value = pixels[y][x];
                Color color = Color.FromArgb(value, value, value);
                bitmap.SetPixel(x, y, color);
            }
        }
        
        bitmap.Save("output.png", ImageFormat.Png);
        Console.WriteLine("Image saved to output.png");
    }
}
```

## Hints

<details>
<summary>Hint 1: Allocating the 2D structure</summary>

**Go**:
```go
picture := make([][]uint8, dy)  // Outer slice
for i := range picture {
    picture[i] = make([]uint8, dx)  // Inner slices
}
```

**C#**:
```csharp
byte[][] picture = new byte[dy][];  // Outer array
for (int i = 0; i < dy; i++)
{
    picture[i] = new byte[dx];  // Inner arrays
}
```
</details>

<details>
<summary>Hint 2: Filling the pixels</summary>

**Go**:
```go
for y := range picture {
    for x := range picture[y] {
        picture[y][x] = uint8((x + y) / 2)  // Example formula
    }
}
```

**C#**:
```csharp
for (int y = 0; y < picture.Length; y++)
{
    for (int x = 0; x < picture[y].Length; x++)
    {
        picture[y][x] = (byte)((x + y) / 2);  // Example formula
    }
}
```
</details>

<details>
<summary>Hint 3: Complete Go Solution</summary>

```go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    // Allocate outer slice
    picture := make([][]uint8, dy)
    
    // Allocate and fill each row
    for y := range picture {
        picture[y] = make([]uint8, dx)
        for x := range picture[y] {
            // Try different formulas:
            // picture[y][x] = uint8((x + y) / 2)
            // picture[y][x] = uint8(x * y)
            picture[y][x] = uint8(x ^ y)  // XOR creates interesting patterns
        }
    }
    
    return picture
}

func main() {
    pic.Show(Pic)
}
```
</details>

<details>
<summary>Hint 4: Complete C# Solution</summary>

```csharp
using System;
using System.Drawing;
using System.Drawing.Imaging;

class Program
{
    static byte[][] Pic(int dx, int dy)
    {
        // Allocate jagged array
        byte[][] picture = new byte[dy][];
        
        // Allocate and fill each row
        for (int y = 0; y < dy; y++)
        {
            picture[y] = new byte[dx];
            for (int x = 0; x < dx; x++)
            {
                // Try different formulas:
                // picture[y][x] = (byte)((x + y) / 2);
                // picture[y][x] = (byte)(x * y);
                picture[y][x] = (byte)(x ^ y);  // XOR creates interesting patterns
            }
        }
        
        return picture;
    }

    static void Main()
    {
        var image = Pic(256, 256);
        DisplayImage(image);
    }

    static void DisplayImage(byte[][] pixels)
    {
        int height = pixels.Length;
        int width = pixels[0].Length;
        
        Bitmap bitmap = new Bitmap(width, height);
        for (int y = 0; y < height; y++)
        {
            for (int x = 0; x < width; x++)
            {
                byte value = pixels[y][x];
                Color color = Color.FromArgb(value, value, value);
                bitmap.SetPixel(x, y, color);
            }
        }
        
        bitmap.Save("output.png", ImageFormat.Png);
        Console.WriteLine("Image saved to output.png");
    }
}
```
</details>

## Learning Objectives

After completing this exercise, you should understand:

- How to create and manipulate 2D slices (Go) or jagged arrays (C#)
- The difference between allocation and initialization
- Using nested loops to process 2D data structures
- How to use range (Go) vs for loops (C#) for indexed iteration
- The memory layout of slices of slices vs rectangular arrays

## Comparison Notes

**Go approach**: Slices of slices are the natural way to represent 2D data. Simple, uniform syntax using `make` and `range`.

**C# approach**: Multiple options (jagged arrays, rectangular arrays, List<List<T>>). Jagged arrays are closest to Go slices but less commonly used than rectangular arrays for uniform grids.

For image data specifically:
- **Go**: 2D slices are appropriate
- **C#**: Often use `Bitmap` class or single-dimensional arrays with calculated indices for performance

Both languages make it straightforward to work with 2D data, but Go's slice model is simpler and more uniform, while C# offers more options optimized for different use cases.
