package functional

// Returns a unique slice containing the elements of `src[]`
// and the elements of all slices provided as `with[]`.
//
// **Note:** The implementation relies on `Uniq` function.
//
// Example:
// 	foo := []int{1, 1, 2, 3, 3, 4, 5}
// 	bar1 := []int{2, 6, 9}
// 	bar2 := []int{7, 9, 10}
// 	bar3 := []int{9, 11, 12}
// 	functional.Union(foo)(bar1, bar2, bar3) // returns [1 2 3 4 5 6 9 7 10 11 12]
func Union[T comparable](src []T) func(with ...[]T) []T {
	length := len(src)

	return func(with ...[]T) []T {
		for _, v := range with {
			length += len(v)
		}

		result := make([]T, 0, length)
		result = append(result, src...)

		for _, v := range with {
			result = append(result, v...)
		}

		return Uniq(result)
	}
}
