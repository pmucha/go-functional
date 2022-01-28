package functional

import (
	"testing"
)

func Test_ZipMap(t *testing.T) {
	keys := []int{-1, 0, 1, 2, 2, 3}
	vals := []interface{}{"minus one", 0.1, 1, nil, "three"}

	zip := ZipMap(keys)(vals)

	if len(zip) != 4 {
		t.Log("a zipped map has wrong length; expected 4, result:", len(zip))
		t.Fail()
	}

	if zip[2] != "three" {
		t.Log("zipping failed; expected \"three\", result:", zip[2])
		t.Fail()
	}

	keys2 := []string{"one", "two", "three"}
	vals2 := []int{1, 2, 3}

	valsAny := ToAny(vals2)

	zip2 := ZipMap(keys2)(valsAny)

	if zip2["one"] != 1 {
		t.Log("zipping any failed; expected 1, result:", zip2["one"])
		t.Fail()
	}
}
