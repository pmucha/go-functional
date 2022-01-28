package functional

import (
	"testing"
)

func Test_ToAny(t *testing.T) {
	ints := []int{-1, 0, 1, 2, 2, 3}
	strings := []string{"foo", "bar", "baz"}
	interfaces := []interface{}{"minus one", 0.1, 1, nil, "three"}

	anyInts := ToAny(ints)
	anyStrings := ToAny(strings)
	anyInterfaces := ToAny(interfaces)

	if len(anyInts) != 6 {
		t.Log("anyfying ints failed; expected 6, result:", len(anyInts))
		t.Fail()
	}

	if len(anyStrings) != 3 {
		t.Log("anyfying strings failed; expected 3, result:", len(anyStrings))
		t.Fail()
	}

	if len(anyInterfaces) != 5 {
		t.Log("anyfying interfaces failed; expected 5, result:", len(anyInterfaces))
		t.Fail()
	}
}
