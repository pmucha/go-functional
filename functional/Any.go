package functional

// Runs a function `f() bool` on all elements of slice `src[]`
// until one returns `true`.
// Returns `true` if at least one of the elements in the `src[]`
// meet the criteria, `false` otherwise.
//
// Example:
//	foo := []int{1, 2, 3, 4, 5}
// 	isEven := func(val int) bool {
// 		return val % 2 == 0
// 	}
// 	functional.Any(isEven)(foo) // returns true
func Any[T any](f func(val T) bool) func(src []T) bool {
	return func(src []T) bool {
		for _, v := range src {
			if f(v) {
				return true
			}
		}
		return false
	}
}
