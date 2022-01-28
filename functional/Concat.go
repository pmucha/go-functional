package functional

// Returns a slice containing all the elements of `src[]`
// followed by the elements of all slices provided as `with[]`.
//
// Example:
// 	src1 := []int{1, 1, 2, 3, 3, 4, 5}
// 	src2 := []int{2, 6, 9}
// 	with1 := []int{7, 9, 10}
// 	with2 := []int{9, 11, 12}
// 	functional.Concat(src1, src2)(with1, with2)
// 		// returns [1 1 2 3 3 4 5 2 6 9 7 9 10 9 11 12]
func Concat[T comparable](src ...[]T) func(with ...[]T) []T {
	length := 0

	for _, v := range src {
		length += len(v)
	}

	return func(with ...[]T) []T {
		for _, v := range with {
			length += len(v)
		}

		result := make([]T, 0, length)

		for _, v := range src {
			result = append(result, v...)
		}

		for _, v := range with {
			result = append(result, v...)
		}

		return result
	}
}
