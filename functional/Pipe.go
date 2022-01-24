package functional

// Performs left-to-right function composition.
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
//	functional.Pipe(square, half, add1)(10) // returns 51
func Pipe[T any](f ...func(T) T) func(T) T {
	return func(x T) T {
		for _, v := range f {
			x = v(x)
		}
		return x
	}
}
