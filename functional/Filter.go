package functional

// Runs a `f()` on all the elements of the `src[]` slice
// and returns a new slice containing only the elements
// that returned `true`. It also returns an error if occured.
//
// Example:
//	foo := []int{1, 2, 3, 4, 5}
//	isEven := func(val int) (bool, error) {
//		return val%2 == 0, nil
//	}
// 	functional.Filter(isEven)(foo) // returns [2, 4], nil
func Filter[T any](f func(val T) (bool, error)) func(src []T) ([]T, error) {
	return func(src []T) ([]T, error) {
		result := make([]T, 0, len(src))

		var returnErr error

		for _, v := range src {
			if fv, err := f(v); fv {
				result = append(result, v)
			} else if err != nil {
				returnErr = err
			}
		}

		return result, returnErr
	}
}
