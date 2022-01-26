package functional

// Returns a slice containing the keys of the provided map `m`.
//
// Example:
// 	m := map[int]string{
// 		1: "Matthew",
// 		2: "Mark",
// 		3: "Luke",
// 		4: "John",
// 	}
//
//	 functional.Values(m) // returns "Matthew", "Mark", "Luke", "John"
//
// **Note:** Due to current Go generics' limitations the usability
// of this function in composing is limited.
func Values[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))

	for _, v := range m {
		result = append(result, v)
	}

	return result
}
