# Go functional

A library providing simple utilities helpful in everyday programming written
in a functional manner -- curried, iteratee-first data-last.

All functions use generics, so they require Go ver. 1.18+.

## Functions

```go
func All[T any](f func(val T) bool) func(src []T) bool
```

Runs a function `f() bool` on all elements of slice `src[]` until one
returns `false`. Returns `true` if all the `src[]` elements meet the `f()`
criteria, `false` otherwise.

Example:

```go
foo := []int{1, 2, 3, 4, 5}
allNotZero := func(val int) bool {
    return val != 0
}
functional.All(allNotZero)(foo) // returns true
```

---

```go
func Any[T any](f func(val T) bool) func(src []T) bool
```

Runs a function `f() bool` on all ments of slice `src[]`
until one returns `true`.
Returns `true` if at least one of the ments in the `src[]`
meet the criteria, `false` otherwise.

Example:

```go
foo := []int{1, 2, 3, 4, 5}
isEven := func(val int) bool {
    return val % 2 == 0
}
functional.Any(isEven)(foo) // returns true
```

---

```go
func Compose[T any](f ...func(T) T) func(T) T
```
Performs right-to-left function composition.

Example:

```go
square := func(x int) int {
    return x * x
}
half := func(x int) int {
    return x / 2
}
add1 := func(x int) int {
    return x + 1
}
functional.Compose(add1, half, square)(10) // returns 51
```
---
```go
func Filter[T any](f func(val T) bool) func(src []T) []T
```

Runs a `f() bool` on all the elements of the `src[]` slice and returns a new
slice containing only the elements that returned `true`.

Example:
```go
foo := []int{1, 2, 3, 4, 5}
isEven := func(val int) bool {
    return val%2 == 0
}
functional.Filter(isEven)(foo) // returns 2, 4
```

---

```go
func Includes[T comparable](val T) func(src []T) bool
```
Checks if slice `src[]` contains a `val`. Returns `bool` result.

Example:
```go
foo := []int{1, 2, 3, 4, 5, 6, 7, 8}
functional.Includes(5)(foo) // returns true
```

---


```go
func Map[T any](f func(val T) T) func(src []T) []T
```
Runs the function `f()` on every element in slice `src[]`. Returns a
processed copy of `src[]` slice.

Example:
```go
foo := []int{1, 2, 3, 4, 5}
mapDouble := func(val int) int {
    return val * 2
}
functional.Map(mapDouble)(foo) // returns 2, 4, 6, 8, 10
```

---

```go
func Pipe[T any](f ...func(T) T) func(T) T
```
Performs left-to-right function composition.

Example:
```go
square := func(x int) int {
    return x * x
}
half := func(x int) int {
    return x / 2
}
add1 := func(x int) int {
    return x + 1
}
functional.Pipe(square, half, add1)(10) // returns 51
```


---

```go
func Reduce[T any](f func(result T, val T) T) func(begin T) func(src []T) T
```

Reduces the `src[]` to a single value, by running it through `f()` starting
from the `begin` value.

Example:
```go
foo := []int{1, 2, 3, 4, 5}
reduceSum := func(result int, val int) int {
    return result + val
}
functional.Reduce(reduceSum)(10)(foo) // returns 25
```

---

```go
func Uniq[T comparable](src []T) []T
```

Returns a copy of the slice `src[]` with all the duplicate values removed.

**Note:** The implementation relies on `Includes` function.

Example:

```go
foo := []int{1, 1, 2, 3, 3, 4, 5}
functional.Uniq(foo) // returns 1, 2, 3, 4, 5
```

---

## TODO

More functions, especially the most interesting and useful in [Ramda](https://ramdajs.com/).

## License
MIT