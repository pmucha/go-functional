package functional

// Runs a function `f()` on all elements of slice `src[]`
// until one returns `false`.
// Returns `true` if all the `src[]` elements meet the `f()`
// criteria, `false` otherwise. Also returns an error if occured.
//
// Example:
//	foo := []int{1, 2, 3, 4, 5}
// 	notZero := func(val int) bool {
// 		return val != 0
// 	}
// 	functional.All(notZero)(foo) // returns true, nil
func All[T any](f func(val T) (bool, error)) func(src []T) (bool, error) {
	return func(src []T) (bool, error) {
		for _, v := range src {
			if result, err := f(v); !result || err != nil {
				return false, err
			}
		}
		return true, nil
	}
}
