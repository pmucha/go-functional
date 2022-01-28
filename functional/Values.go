package functional

// Returns a slice containing the values of the provided map `m`.
// Also returns an error if occurred.
//
// Example:
// 	m := map[int]string{
// 		1: "Matthew",
// 		2: "Mark",
// 		3: "Luke",
// 		4: "John",
// 	}
//
//	 functional.Values(m) // returns ["Matthew" "Mark" "Luke" "John"], nil
func Values[K comparable, V any](m map[K]V) ([]V, error) {
	result := make([]V, 0, len(m))

	for _, v := range m {
		result = append(result, v)
	}

	return result, nil
}
