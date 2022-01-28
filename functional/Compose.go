package functional

// Performs right-to-left function composition.
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
//	functional.Compose(add1, half, square)(10) // returns 51, nil
func Compose[T any](f ...func(T) (T, error)) func(T) (T, error) {
	return func(x T) (T, error) {
		var zero T
		var err error
		for i := len(f) - 1; i >= 0; i-- {
			x, err = f[i](x)
			if err != nil {
				return zero, err
			}
		}
		return x, nil
	}
}
