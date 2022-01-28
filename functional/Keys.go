package functional

// Returns a slice containing the keys of the provided map `m`.
// Also returns an error if occured.
//
// Example:
//	 m := map[string]int{
//	 	"Matthew": 1,
//	 	"Mark": 2,
//	 	"Luke": 3,
//	 	"John": 4,
//	 }
//
//	 functional.Keys(m) // returns ["Matthew" "Mark" "Luke" "John"], nil
func Keys[K comparable, V any](m map[K]V) ([]K, error) {
	result := make([]K, 0, len(m))

	for k := range m {
		result = append(result, k)
	}

	return result, nil
}
