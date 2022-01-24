# Go functional

A library providing simple utilities helpful in everyday programming written
in a functional manner -- curried, iteratee-first data-last.

All functions use generics, so they require Go ver. 1.18+.

## Functions

```go
func All[T any](f func(val T) bool) func(src []T) bool
```

All runs a function `f() bool` on all elements of slice `src[]` until one
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
func Includes[T comparable](val T) func(src []T) bool
```
Includes checks if any of the elements in slice `src[]` is equal to `val`.
Returns `bool`

Example:
```go
foo := []int{1, 2, 3, 4, 5, 6, 7, 8}
functional.Includes(5)(foo) // returns true
```

---


```go
func Map[T any](f func(val T) T) func(src []T) []T
```
Map runs the function `f()` on every element in slice `src[]`. Returns a
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

Uniq returns a copy of the slice `src[]` with the duplicate values removed and ordering preserved.

Note: The implementation relies on `Includes` function.

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