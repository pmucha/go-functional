package functional

import "testing"

func Test_Contains(t *testing.T) {
	intSlice := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}

	if !Contains(intSlice)(5) {
		t.Log("a number not found in a slice")
		t.Fail()
	}

	if Contains(intSlice)(10) {
		t.Log("a number in a slice failed")
		t.Fail()
	}

	stringSlice := []string{"foo", "bar", "baz"}

	if !Contains(stringSlice)("foo") {
		t.Log("a string not found in a slice")
		t.Fail()
	}

	if Contains(stringSlice)("fail") {
		t.Log("a string in a slice failed")
		t.Fail()
	}
}
