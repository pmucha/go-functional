package functional

// Converts any slice of type `[]T` to `[]any`.
// Returns the copy of the slice and the error
// if occurred.
//
// Example:
//
// 	strings := []string{"foo", "bar", "baz"}
// 	ToAny(strings) // returns []interface [foo bar baz], nil
func ToAny[T any](s []T) ([]any, error) {
	result := make([]any, 0, len(s))

	if len(s) == 0 {
		return result, nil
	}

	for _, v := range s {
		result = append(result, v)
	}

	return result, nil
}
