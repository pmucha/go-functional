package functional

// Returns a slice containing all the elements of
// provided `src[]` slices.
//
// Example:
// 	src1 := []int{1, 1, 2, 3, 3, 4, 5}
// 	src2 := []int{2, 6, 9}
// 	functional.Concat(src1, src2)
// 		// returns [1 1 2 3 3 4 5 2 6 9], nil
func Concat[T comparable](src ...[]T) ([]T, error) {
	length := 0

	for _, v := range src {
		length += len(v)
	}

	result := make([]T, 0, length)

	for _, v := range src {
		result = append(result, v...)
	}

	return result, nil
}
