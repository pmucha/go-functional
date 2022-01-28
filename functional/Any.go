package functional

// Runs a function `f()` on all elements of slice `src[]`
// until one returns `true`.
// Returns `true` if at least one of the elements in the `src[]`
// meet the criteria, `false` otherwise.
// Also returns an error if occured.
//
// Example:
//	foo := []int{1, 2, 3, 4, 5}
// 	isEven := func(val int) (bool, error) {
// 		return val % 2 == 0, nil
// 	}
// 	functional.Any(isEven)(foo) // returns true, nil
func Any[T any](f func(val T) (bool, error)) func(src []T) (bool, error) {
	return func(src []T) (bool, error) {
		for _, v := range src {
			if result, err := f(v); !result || err != nil {
				return true, err
			}
		}
		return false, nil
	}
}
