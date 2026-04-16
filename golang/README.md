# Go Language Learning

This section contains my journey learning Go through the [Tour of Go](https://go.dev/tour/list) interactive challenges.

## Overview

The Tour of Go challenges are a series of 21 hands-on exercises that cover the core concepts of the Go programming language. Each challenge builds on the previous one, helping solidify understanding through practical implementation.

## Progress

**Completed Challenges:** 10/21

### ✅ Completed (Challenges 1-10)

- **Phase 1: Basics** - Packages, Variables, Functions
  - [Challenge 01](challenge01/main.go) - Hello Playground (packages, imports, exported names)
  - [Challenge 02](challenge02/main.go) - Temperature Converter (functions, parameters, returns)
  - [Challenge 03](challenge03/main.go) - Currency Exchange (variables, constants)

- **Phase 2: Flow Control**
  - [Challenge 04](challenge04/main.go) - FizzBuzz (for loops, if/else)
  - [Challenge 05](challenge05/main.go) - Grade Calculator (switch statements)
  - [Challenge 06](challenge06/main.go) - Countdown Timer (defer)

- **Phase 3: Composite Types**
  - [Challenge 07](challenge07/main.go) - Pointer Swap (pointers)
  - [Challenge 08](challenge08/main.go) - Contact Card (structs)
  - [Challenge 09](challenge09/main.go) - Shopping List (slices, arrays)
  - [Challenge 10](challenge10/main.go) - Phonebook (maps, closures)

### 🔲 Remaining (Challenges 11-21)

- **Phase 4: Methods & Interfaces** (11-14)
- **Phase 5: Generics** (15-16)
- **Phase 6: Concurrency** (17-20)
- **Capstone** (21)

See [detailed challenge documentation](docs/README.md) for full descriptions and expected outputs.

## Quick Start

### Prerequisites

- Go 1.21+ installed
- VS Code with Go extension (recommended)

### Running a Challenge

From the repository root:

```bash
go run golang/challenge01/main.go
```

Or from within the `golang/` directory:

```bash
cd golang
go run challenge01/main.go
```

### Tour of Go

Run the Tour of Go locally to follow along with the concepts:

```bash
go install golang.org/x/website/tour@latest
tour
```

Visit: http://127.0.0.1:3999

## Challenge Structure

Each challenge folder contains:
- `main.go` - The implementation file
- Focuses on specific Go concepts
- Designed to be simpler than Tour of Go exercises for learning reinforcement

## Resources

- [Tour of Go](https://go.dev/tour/list)
- [Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Challenge Details](docs/README.md)

## Next Steps

Continue with Challenge 11 (Shape Calculator) to learn about methods and interfaces.
