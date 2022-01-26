package functional

// Returns a slice containing the keys of the provided map `m`.
//
// Example:
//	 m := map[string]int{
//	 	"Matthew": 1,
//	 	"Mark": 2,
//	 	"Luke": 3,
//	 	"John": 4,
//	 }
//
//	 functional.Keys(m) // returns "Matthew", "Mark", "Luke", "John"
//
// **Note:** Due to current Go generics' limitations the usability
// of this function in composing is limited.
func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))

	for k := range m {
		result = append(result, k)
	}

	return result
}
