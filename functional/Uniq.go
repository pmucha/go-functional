package functional

// Returns a copy of the slice `src[]` with all
// the duplicate values removed. Also returns an error if occurred.
//
// Example:
// 	foo := []int{1, 1, 2, 3, 3, 4, 5}
// 	functional.Uniq(foo) // returns [1 2 3 4 5], nil
func Uniq[T comparable](src []T) ([]T, error) {
	var zero []T
	var err error
	result := make([]T, 0, len(src))

	for _, v := range src {
		inc, err := Includes(v)(result)
		if !inc {
			result = append(result, v)
		} else if err != nil {
			return zero, err
		}
	}

	return result, err
}
