package functional

// Checks if slice `src[]` contains a `val`.
// Returns `bool` result and an error if occured.
//
// **Note**: It works just like `Includes` except
// it accepts the arguments in reverse order.
//
// Example:
// 	foo := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	functional.Contains(foo)(5) // returns true, nil
func Contains[T comparable](src []T) func(val T) (bool, error) {
	return func(val T) (bool, error) {
		for _, v := range src {
			if val == v {
				return true, nil
			}
		}
		return false, nil
	}
}
