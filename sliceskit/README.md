# sliceskit

`sliceskit` package contains useful slice utility functions to boost writing Go applications. It provides functional programming utilities similar to JavaScript's Array methods but with Go's type safety and performance.

## Available Functions

### Core Functions

- [x] [Any](./any.go) - Check if any element satisfies the predicate
- [x] [Every](./every.go) - Check if all elements satisfy the predicate
- [x] [Chunk](./chunk.go) - Split a slice into smaller chunks
- [x] [Find](./find.go) - Find the first element that satisfies the predicate
- [x] [FindPtr](./find.go) - Find the first pointer in a slice of pointers that satisfies the predicate
- [x] [Filter](./filter.go) - Filter elements based on a predicate
- [x] [Map](./map.go) - Transform elements using a mapping function
- [x] [Reduce](./reduce.go) - Reduce a slice to a single value

## Function Details

### Any

```go
func Any[Slice ~[]E, E any](s Slice, anyFunc func(E) bool) bool
```

Returns `true` if at least one element in the slice satisfies the predicate, `false` otherwise.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
hasEven := sliceskit.Any(numbers, func(n int) bool { return n%2 == 0 })
// hasEven = true
```

### Every

```go
func Every[Slice ~[]E, E any](s Slice, everyFunc func(E) bool) bool
```

Returns `true` if all elements in the slice satisfy the predicate, `false` otherwise.

**Example:**

```go
numbers := []int{2, 4, 6, 8}
allEven := sliceskit.Every(numbers, func(n int) bool { return n%2 == 0 })
// allEven = true
```

### Chunk

```go
func Chunk[Slice ~[]E, E any](s Slice, size int) [][]E
```

Splits a slice into smaller chunks of the specified size. Returns an empty slice if size is zero or negative.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5, 6}
chunks := sliceskit.Chunk(numbers, 2)
// chunks = [[1, 2], [3, 4], [5, 6]]
```

### Find

```go
func Find[Slice ~[]E, E any](s Slice, findFunc func(E) bool) (E, bool)
```

Returns the first element that satisfies the predicate and a boolean indicating if found.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
found, exists := sliceskit.Find(numbers, func(n int) bool { return n > 3 })
// found = 4, exists = true
```

### FindPtr

```go
func FindPtr[Slice ~[]*E, E any](s Slice, findFunc func(*E) bool) *E
```

Returns the first pointer in a slice of pointers that satisfies the predicate, or `nil` if not found.

**Example:**

```go
users := []*User{{Name: "Alice"}, {Name: "Bob"}}
found := sliceskit.FindPtr(users, func(u *User) bool { return u.Name == "Bob" })
// found points to the Bob user, or nil if not found
```

### Filter

```go
func Filter[Slice ~[]E, E any](s Slice, filterFunc func(E) bool) Slice
func FilterWithIndex[Slice ~[]E, E any](s Slice, filterFunc func(E, int) bool) Slice
func FilterWithFuncErr[Slice ~[]E, E any](s Slice, filterFunc func(E) (bool, error)) (Slice, error)
func FilterWithIndexAndFuncErr[Slice ~[]E, E any](s Slice, filterFunc func(E, int) (bool, error)) (Slice, error)
```

Returns a new slice containing only the elements that satisfy the predicate.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5, 6}
evens := sliceskit.Filter(numbers, func(n int) bool { return n%2 == 0 })
// evens = [2, 4, 6]
```

### Map

```go
func Map[Slice ~[]E, T any, E any](s Slice, mapFunc func(E) T) []T
func MapWithIndex[Slice ~[]E, T any, E any](s Slice, mapFunc func(E, int) T) []T
func MapWithFuncErr[Slice ~[]E, T any, E any](s Slice, mapFunc func(E) (T, error)) ([]T, error)
func MapWithIndexAndFuncErr[Slice ~[]E, T any, E any](s Slice, mapFunc func(E, int) (T, error)) ([]T, error)
```

Transforms each element in the slice using the provided function.

**Example:**

```go
numbers := []int{1, 2, 3, 4}
squares := sliceskit.Map(numbers, func(n int) int { return n * n })
// squares = [1, 4, 9, 16]
```

### Reduce

```go
func Reduce[Slice ~[]E, U any, E any](s Slice, reduceFunc func(prev U, current E) U, initial U) U
func ReduceWithIndex[Slice ~[]E, U any, E any](s Slice, reduceFunc func(prev U, current E, index int) U, initial U) U
func ReduceWithFuncErr[Slice ~[]E, U any, E any](s Slice, reduceFunc func(prev U, current E) (U, error), initial U) (U, error)
func ReduceWithIndexAndFuncErr[Slice ~[]E, U any, E any](s Slice, reduceFunc func(prev U, current E, index int) (U, error), initial U) (U, error)
```

Reduces a slice to a single value by applying a function to each element and accumulating the result.

**Example:**

```go
numbers := []int{1, 2, 3, 4, 5}
sum := sliceskit.Reduce(numbers, func(prev, current int) int { return prev + current }, 0)
// sum = 15
```

## Features

- **Type Safe**: All functions use Go generics for compile-time type safety
- **Nil Safe**: Functions handle nil slices gracefully
- **Error Handling**: Variants with error handling for robust applications
- **Index Support**: Many functions have variants that provide element indices
- **Performance**: Optimized for Go's slice operations
- **Comprehensive Testing**: All functions have extensive test coverage

## Installation

```bash
go get github.com/umefy/godash/sliceskit
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/umefy/godash/sliceskit"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // Filter even numbers
    evens := sliceskit.Filter(numbers, func(n int) bool { return n%2 == 0 })

    // Map to squares
    squares := sliceskit.Map(evens, func(n int) int { return n * n })

    // Sum all squares
    sum := sliceskit.Reduce(squares, func(prev, current int) int { return prev + current }, 0)

    fmt.Printf("Sum of squares of even numbers: %d\n", sum)
}
```
