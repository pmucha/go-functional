package functional

// Combines provided slices into a `map` which has `keys[]` elements as keys,
// and `vals[]` elements as values (similar to `ZipObj` in _Ramda_).
// If the `keys[]` is empty, the result will be an empty map.
// If the `keys[]` isn't unique, the subsequent values will
// overwrite the previously set ones. If the `values[]` is shorter than the
// `keys[]`, the exceeding elements won't be used.
// Returns the new map and an error if occurred.
//
// Example:
//
// 	keys := []int{-1, 0, 1, 2, 2, 3}
// 	vals := []string{"minus one", "zero", "one", "two", "three"}
//
// 	functional.ZipMap(keys, vals)
// 		// returns [-1: "minus one", 0: "zero", 1: "one", 2: "three"], nil
func ZipMap[K comparable, V any](keys []K, vals []V) (map[K]V, error) {
	result := make(map[K]V)

	if len(keys) == 0 || len(vals) == 0 {
		return result, nil
	}

	l := len(vals)

	for k, v := range keys {
		if k < l {
			result[v] = vals[k]
		}
	}

	return result, nil

}
