package functional

import "testing"

func Test_Keys(t *testing.T) {
	m := map[string]int{
		"Matthew": 1,
		"Mark":    2,
		"Luke":    3,
		"John":    4,
	}

	if keys, _ := Keys(m); len(keys) != 4 {
		t.Log("the number of keys doesn't match the original map; expected 4, result:", len(keys))
		t.Fail()
	}
}
