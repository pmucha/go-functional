package functional

import "testing"

func Test_Values(t *testing.T) {
	m := map[int]string{
		1: "Matthew",
		2: "Mark",
		3: "Luke",
		4: "John",
	}

	if vals, _ := Values(m); len(vals) != 4 {
		t.Log("the number of keys doesn't match the original map; expected 4, result:", len(vals))
		t.Fail()
	}
}
