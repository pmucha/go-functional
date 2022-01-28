package functional

// Returns a unique slice containing the elements of
// slices provided as `src[]`. Also returns an error
// if occurred.
//
// Example:
// 	src1 := []int{1, 1, 2, 3, 3, 4, 5}
// 	src2 := []int{2, 6, 9}
// 	functional.Union(src1, src2) // returns [1 2 3 4 5 6 9], nil
func Union[T comparable](src ...[]T) ([]T, error) {
	var zero []T
	concatenated, err := Concat(src...)
	if err != nil {
		return zero, err
	}
	return Uniq(concatenated)
}
