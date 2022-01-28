package functional

// Reduces the `src[]` to a single value, by running it through `f()`
// starting from the `begin` value. Returns this value and an error
// if occurred.
//
// Example:
//	foo := []int{1, 2, 3, 4, 5}
// 	sum := func(result int, val int) (int, error) {
// 		return result + val
// 	}
// 	functional.Reduce(sum)(10)(foo) // returns 25, nil
func Reduce[T any](f func(result T, val T) (T, error)) func(begin T) func(src []T) (T, error) {
	return func(begin T) func(src []T) (T, error) {
		result := begin

		return func(src []T) (T, error) {
			var zero T
			var err error
			for _, v := range src {
				result, err = f(result, v)
				if err != nil {
					return zero, err
				}
			}

			return result, nil
		}
	}
}
