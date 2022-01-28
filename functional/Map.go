package functional

// Runs the function `f()` on every element in slice `src[]`.
// Returns a processed copy of `src[]` slice.
// Also returns an error if occurred.
//
// Example:
// 	foo := []int{1, 2, 3, 4, 5}
// 	mapDouble := func(val int) (int, error) {
// 		return val * 2, nil
// 	}
// 	functional.Map(mapDouble)(foo) // returns [2 4 6 8 10], nil
func Map[T any](f func(val T) (T, error)) func(src []T) ([]T, error) {
	return func(src []T) ([]T, error) {
		var err error

		result := make([]T, len(src))
		copy(result, src)

		for k, v := range result {
			var fv T
			fv, err = f(v)
			result[k] = fv
		}

		return result, err
	}
}
