package functional

import "testing"

func Test_Contains(t *testing.T) {
	intSlice := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}

	if result, err := Contains(intSlice)(5); !result || err != nil {
		t.Log("a number not found in a slice")
		t.Fail()
	}

	if result, err := Contains(intSlice)(10); result || err != nil {
		t.Log("a number should have not been found in the slice")
		t.Fail()
	}

	stringSlice := []string{"foo", "bar", "baz"}

	if result, err := Contains(stringSlice)("foo"); !result || err != nil {
		t.Log("a string not found in a slice")
		t.Fail()
	}

	if result, err := Contains(stringSlice)("fail"); result || err != nil {
		t.Log("a string should not have been found in the slice")
		t.Fail()
	}
}
