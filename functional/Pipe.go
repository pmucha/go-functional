package functional

// Performs left-to-right function composition.
//
// Example:
//
// 	square := func(x int) (int, error) {
// 		return x * x, nil
// 	}
// 	half := func(x int) (int, error) {
// 		return x / 2, nil
// 	}
// 	add1 := func(x int) (int, error) {
// 		return x + 1, nil
// 	}
//	functional.Pipe(square, half, add1)(10) // returns 51, nil
func Pipe[T any](f ...func(T) (T, error)) func(T) (T, error) {
	return func(x T) (T, error) {
		var zero T
		var err error
		for _, v := range f {
			x, err = v(x)
			if err != nil {
				return zero, err
			}
		}
		return x, nil
	}
}
