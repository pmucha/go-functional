# Go functional

A library providing simple utilities helpful in everyday programming written
in a functional manner. They are curried if possible and reasonable.
They are built as "iteratee-first data-last", meaning:
`[do-what] -- [do how] -- [to what]` returning `result, error`.
For example: do `Map`, do `func`, to `slice`, returning `new slice` and possible error.

Most functions use generics, so they require Go ver. 1.18+.

## Installation

To download the module use:
```
go get github.com/pmucha/go-functional/functional
```

To use it inside your program:
```go
import "github.com/pmucha/go-functional/functional"
```

This way all the functions will be available as
`functional.FunctionName()`. Use:
```go
import f "github.com/pmucha/go-functional/functional"
```
to make them available through an alias: `f.FunctionName()`.

## Functions

### All

```go
func All[T any](f func(val T) (bool, error)) func(src []T) (bool, error) {
```

Runs a function `f()` on all elements of slice `src[]`
until one returns `false`.
Returns `true` if all the `src[]` elements meet the `f()`
criteria, `false` otherwise. Also returns an error if occured.

Example:

```go
foo := []int{1, 2, 3, 4, 5}
notZero := func(val int) (bool, nil) {
    return val != 0, nil
}
functional.All(notZero)(foo) // returns true, nil
```

---

### Any

```go
func Any[T any](f func(val T) bool) func(src []T) bool
```

Runs a function `f()` on all elements of slice `src[]`
until one returns `true`.
Returns `true` if at least one of the elements in the `src[]`
meet the criteria, `false` otherwise.
Also returns an error if occured.

Example:

```go
foo := []int{1, 2, 3, 4, 5}
isEven := func(val int) (bool, error) {
    return val % 2 == 0, nil
}
functional.Any(isEven)(foo) // returns true, nil
```

---

### Compose

```go
func Compose[T any](f ...func(T) (T, error)) func(T) (T, error) {
```
Performs right-to-left function composition.

Example:

```go
square := func(x int) (int, error) {
	return x * x, nil
}
half := func(x int) (int, error) {
	return x / 2, nil
}
add1 := func(x int) (int, error) {
	return x + 1, nil
}
functional.Compose(add1, half, square)(10) // returns 51, nil
```

---

### Concat

```go
func Concat[T comparable](src ...[]T) ([]T, error)
```

Returns a slice containing all the elements of
provided `src[]` slices.

Example:

```go
src1 := []int{1, 1, 2, 3, 3, 4, 5}
src2 := []int{2, 6, 9}
functional.Concat(src1, src2)
	// returns [1 1 2 3 3 4 5 2 6 9], nil
```

---

### Contains

```go
func Contains[T comparable](src []T) func(val T) (bool, error) {
```
Checks if slice `src[]` contains a `val`.
Returns `bool` result and an error if occured.
**Note**: It works just like `Includes` except
it accepts the arguments in reverse order.

Example:

```go
foo := []int{1, 2, 3, 4, 5, 6, 7, 8}
functional.Contains(foo)(5) // returns true, nil
```

---

### Filter
```go
func Filter[T any](f func(val T) (bool, error)) func(src []T) ([]T, error)
```
Runs a `f()` on all the elements of the `src[]` slice
and returns a new slice containing only the elements
that returned `true`. It also returns an error if occured.

Example:

```go
foo := []int{1, 2, 3, 4, 5}
isEven := func(val int) (bool, error) {
	return val%2 == 0, nil
}
functional.Filter(isEven)(foo) // returns [2, 4], nil
```

---

### Includes

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

### Keys

```go
func Keys[K comparable, V any](m map[K]V) []K
```

Returns a slice containing the keys of the provided map `m`.

Example:

```go
m := map[string]int{
    "Matthew": 1,
    "Mark": 2,
    "Luke": 3,
    "John": 4,
}

functional.Keys(m) // returns ["Matthew" "Mark" "Luke" "John"]
```

**Note:** Due to current Go generics' limitations the usability of this
function in composing is limited.

---

### Map

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
functional.Map(mapDouble)(foo) // returns [2 4 6 8 10]
```

---

### Pipe

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

### Reduce

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

### ToAny

```go
func ToAny[T any](s []T) []any
```

Converts any slice of type `[]T` to `[]any` which might be helpful for
functions such as `ZipMap()`.

Example:

```go
strings := []string{"foo", "bar", "baz"}
ToAny(strings) // returns []interface [foo bar baz]
```

---

### Union

```go
func Union[T comparable](src []T) func(with ...[]T) []T
```

Returns a unique slice containing the elements of `src[]` and the elements
of all slices provided as `with[]`.

**Note:** The implementation relies on the `Uniq` function.

Example:

```go
foo := []int{1, 1, 2, 3, 3, 4, 5}
bar1 := []int{2, 6, 9}
bar2 := []int{7, 9, 10}
bar3 := []int{9, 11, 12}
functional.Union(foo)(bar1, bar2, bar3) // returns [1 2 3 4 5 6 9 7 10 11 12]
```

---
### Uniq

```go
func Uniq[T comparable](src []T) []T
```

Returns a copy of the slice `src[]` with all the duplicate values removed.

**Note:** The implementation relies on the `Includes` function.

Example:

```go
foo := []int{1, 1, 2, 3, 3, 4, 5}
functional.Uniq(foo) // returns [1 2 3 4 5]
```

---

### Values

```go
func Values[K comparable, V any](m map[K]V) []V
```

Returns a slice containing the keys of the provided map `m`.

Example:

```go
m := map[int]string{
    1: "Matthew",
    2: "Mark",
    3: "Luke",
    4: "John",
}

functional.Values(m) // returns ["Matthew" "Mark" "Luke" "John"]
```

**Note:** Due to current Go generics' limitations the usability of this
function in composing is limited.

---

### ZipMap

```go
func ZipMap[T comparable](keys []T) func(vals []any) map[T]any
```

Combines provided slices into a `map` which has `keys[]` elements as keys,
and `vals[]` elements as values (similar to `ZipObj` in _Ramda_).
If the `keys[]` is empty, the result will be an empty map.
If the `keys[]` isn't unique, the subsequent values will
overwrite the previously set ones. If the `values[]` is shorter than the
`keys[]`, the exceeding elements won't be used.

Example:

```go
keys := []int{-1, 0, 1, 2, 2, 3}
vals := []string{"minus one", "zero", "one", "two", "three"}

functional.ZipMap(keys)(vals)
    // returns [-1: "minus one", 0: "zero", 1: "one", 2: "three"]
```

## TODO

More functions, especially the most interesting and useful in [Ramda](https://ramdajs.com/).

## License
MIT