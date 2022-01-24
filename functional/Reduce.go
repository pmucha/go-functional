package functional

// Reduces the `src[]` to a single value, by running it through `f()`
// starting from the `begin` value.
//
// Example:
//	foo := []int{1, 2, 3, 4, 5}
// 	reduceSum := func(result int, val int) int {
// 		return result + val
// 	}
// 	functional.Reduce(reduceSum)(10)(foo) // returns 25
func Reduce[T any](f func(result T, val T) T) func(begin T) func(src []T) T {
	return func(begin T) func(src []T) T {
		result := begin

		return func(src []T) T {
			for _, v := range src {
				result = f(result, v)
			}

			return result
		}
	}
}
