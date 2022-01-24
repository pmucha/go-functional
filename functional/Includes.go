package functional

// Checks if slice `src[]` contains a `val`.
// Returns `bool` result.
//
// Example:
// 	foo := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	functional.Includes(5)(foo) // returns true
func Includes[T comparable](val T) func(src []T) bool {
	return func(src []T) bool {
		for _, v := range src {
			if val == v {
				return true
			}
		}
		return false
	}
}
