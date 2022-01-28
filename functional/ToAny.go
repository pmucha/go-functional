package functional

// Converts any slice of type `[]T` to `[]any` which might
// be helpful for functions such as `ZipMap()`.
//
// Example:
//
// 	strings := []string{"foo", "bar", "baz"}
// 	ToAny(strings) // returns []interface [foo bar baz]
func ToAny[T any](s []T) []any {
	result := make([]any, 0, len(s))

	if len(s) == 0 {
		return result
	}

	for _, v := range s {
		result = append(result, v)
	}

	return result
}
