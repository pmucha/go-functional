package functional

// Checks if slice `src[]` contains a `val`.
// Returns `bool` result. Also it returns
// an error if occured.
//
// **Note**: It works just like `Contains` except
// it accepts the arguments in reverse order.
//
// Example:
// 	foo := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	functional.Includes(5)(foo) // returns true, nil
func Includes[T comparable](val T) func(src []T) (bool, error) {
	return func(src []T) (bool, error) {
		for _, v := range src {
			if val == v {
				return true, nil
			}
		}
		return false, nil
	}
}
