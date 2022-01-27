package functional

// Returns a copy of the slice `src[]` with all
// the duplicate values removed.
//
// **Note:** The implementation relies on the `Includes` function.
//
// Example:
// 	foo := []int{1, 1, 2, 3, 3, 4, 5}
// 	functional.Uniq(foo) // returns 1, 2, 3, 4, 5
func Uniq[T comparable](src []T) []T {
	result := make([]T, 0, len(src))

	for _, v := range src {
		if !Includes(v)(result) {
			result = append(result, v)
		}
	}

	return result
}
