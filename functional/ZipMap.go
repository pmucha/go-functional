package functional

// Combines provided slices into a `map` which has `keys[]` elements as keys,
// and `vals[]` elements as values (similar to `ZipObj` in _Ramda_).
// If the `keys[]` is empty, the result will be an empty map.
// If the `keys[]` isn't unique, the subsequent values will
// overwrite the previously set ones. If the `values[]` is shorter than the
// `keys[]`, the exceeding elements won't be used.
//
// Example:
//
// 	keys := []int{-1, 0, 1, 2, 2, 3}
// 	vals := []string{"minus one", "zero", "one", "two", "three"}
//
// 	functional.ZipMap(keys, vals)(vals)
// 		// returns -1 -> "minus one", 0 -> "zero", 1 -> "one", 2 -> "three"
//
// 	**Known issue**: Currently Go fails to infer the type
//	for `vals` in returned function, unless it is somehow
// 	provided in the primary function. That's why it's needed
// 	to provide `vals` in the primary function as well.
func ZipMap[K comparable, V any](keys []K, _ []V) func(vals []V) map[K]V {
	result := make(map[K]V)

	return func(vals []V) map[K]V {
		if len(keys) == 0 || len(vals) == 0 {
			return result
		}

		l := len(vals)
		for k, v := range keys {
			if k < l {
				result[v] = vals[k]
			}
		}
		return result
	}
}
