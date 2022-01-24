package functional

// Runs the function `f()` on every element in slice `src[]`.
// Returns a processed copy of `src[]` slice.
//
// Example:
// 	foo := []int{1, 2, 3, 4, 5}
// 	mapDouble := func(val int) int {
// 		return val * 2
// 	}
// 	functional.Map(mapDouble)(foo) // returns 2, 4, 6, 8, 10
func Map[T any](f func(val T) T) func(src []T) []T {
	return func(src []T) []T {
		result := make([]T, len(src))
		copy(result, src)

		for k, v := range result {
			result[k] = f(v)
		}

		return result
	}
}
