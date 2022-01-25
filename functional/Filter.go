package functional

// Runs a `f() bool` on all the elements of the `src[]` slice
// and returns a new slice containing only the elements
// that returned `true`.
//
// Example:
//	foo := []int{1, 2, 3, 4, 5}
//	isEven := func(val int) bool {
//		return val%2 == 0
//	}
// 	functional.Filter(isEven)(foo) // returns 2, 4
func Filter[T any](f func(val T) bool) func(src []T) []T {
	return func(src []T) []T {
		result := make([]T, 0, len(src))

		for _, v := range src {
			if f(v) {
				result = append(result, v)
			}
		}

		return result
	}
}
