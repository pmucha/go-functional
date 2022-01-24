package functional

// Runs a function `f() bool` on all elements of slice `src[]`
// until one returns `false`.
// Returns `true` if all the `src[]` elements meet the `f()`
// criteria, `false` otherwise.
//
// Example:
//	foo := []int{1, 2, 3, 4, 5}
// 	allNotZero := func(val int) bool {
// 		return val != 0
// 	}
// 	functional.All(allNotZero)(foo) // returns true
func All[T any](f func(val T) bool) func(src []T) bool {
	return func(src []T) bool {
		for _, v := range src {
			if !f(v) {
				return false
			}
		}
		return true
	}
}
