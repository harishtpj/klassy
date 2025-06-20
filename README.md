# Klassy

[![Go Reference](https://pkg.go.dev/badge/github.com/harishtpj/klassy.svg)](https://pkg.go.dev/github.com/harishtpj/klassy)
[![Go Report Card](https://goreportcard.com/badge/github.com/harishtpj/klassy)](https://goreportcard.com/report/github.com/harishtpj/klassy)

Klassy provides object-oriented style wrappers for Go's standard library `strings`, `slices`, and `maps` packages with chainable methods and fluent API design.

## Features

- **Method Chaining**: Fluent API for readable code
- **Type Safety**: Leverages Go generics for type-safe operations
- **Standard Library Compatible**: Wraps existing `strings` and `slices` functions
- **Zero Dependencies**: Built only on Go standard library

## Installation

```bash
go get github.com/harishtpj/klassy
```

## Quick Start

### String Operations

```go
package main

import (
    "fmt"
    "github.com/harishtpj/klassy/String"
)

func main() {
    text := String.New("  Hello, World!  ")
    
    result := text.
        TrimSpace().
        ToUpper().
        Replace("WORLD", "GO", -1)
    
    fmt.Println(result.Value()) // "HELLO, GO!"
}
```

### Slice Operations

```go
package main

import (
    "fmt"
    "github.com/harishtpj/klassy/Slice"
)

func main() {
    numbers := Slice.New([]int{1, 2, 3, 4, 5})
    
    // Transform to strings and join
    result := numbers.Map(func(x int) any { return x * 2 })
    fmt.Println(result.Items) // [2 4 6 8 10]
    
    // Type-safe transformation
    doubled := Slice.MapTo(numbers, func(x int) int { return x * 2 })
    fmt.Println(doubled.Items) // [2 4 6 8 10]
}
```

## Packages

### String Package

Provides a `String` type with chainable string operations. It supports all functions that the stdlib's strings package support:

```go
s := String.New("hello world")
result := s.ToUpper().Replace("WORLD", "GO", 1)
```

### Slice Package

Provides a generic `Slice[T]` type with chainable slice operations:

```go
slice := Slice.New([]int{1, 2, 3})
slice.Append(4, 5, 6)
found := slice.Contains(3)
```

**Key Methods:**
- `Append`, `Push`, `Concat`
- `Contains`, `ContainsFunc`
- `Map`, `MapTo` - Transform elements with type safety
- `At`, `Length`, `Clone`
- `All`, `Backward` - Iterator support

**Generic Functions:**
- `MapTo[T, U any](slice Slice[T], fn func(T) U) Slice[U]` - Type-safe transformations

## Examples

### String Processing Pipeline

```go
text := String.New("  The Quick Brown Fox  ")
processed := text.
    TrimSpace().
    ToLower().
    Replace("fox", "dog", 1).
    Fields()

fmt.Println(processed) // ["the", "quick", "brown", "dog"]
```

### Slice Transformations

```go
words := Slice.New([]string{"hello", "world", "go"})
lengths := Slice.MapTo(words, func(s string) int { return len(s) })
fmt.Println(lengths.Items) // [5 5 2]

// Check if any word has length > 4
hasLongWord := lengths.ContainsFunc(func(n int) bool { return n > 4 })
fmt.Println(hasLongWord) // true
```

## Design Philosophy

Klassy follows these principles:

1. **Fluent API**: Methods return the same type for chaining
2. **Immutability**: String operations return new instances
3. **Type Safety**: Leverage Go's type system and generics
4. **Familiar API**: Mirror standard library function names
5. **Performance**: Minimal overhead over standard library

## API Reference

For detailed API documentation, visit [pkg.go.dev](https://pkg.go.dev/github.com/harishtpj/klassy).

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

#### Copyright © 2025 [M.V.Harish Kumar](https://github.com/harishtpj). <br>

#### This project is [MIT](https://github.com/harishtpj/klassy/blob/ec02a753a71741aac41f57d55c4bc4a66a724291/LICENSE) licensed.

## Acknowledgments

- Inspired by method chaining patterns in other languages
- Built on Go's excellent standard library
- Leverages Go 1.18+ generics for type safety

## TODO
- Improve Slice API
- Create Map API
