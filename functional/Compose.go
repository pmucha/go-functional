package functional

// Performs right-to-left function composition.
//
// Example:
//
// 	square := func(x int) int {
// 		return x * x
// 	}
// 	half := func(x int) int {
// 		return x / 2
// 	}
// 	add1 := func(x int) int {
// 		return x + 1
// 	}
//	functional.Compose(add1, half, square)(10) // returns 51
func Compose[T any](f ...func(T) T) func(T) T {
	return func(x T) T {
		for i := len(f) - 1; i >= 0; i-- {
			x = f[i](x)
		}
		return x
	}
}
