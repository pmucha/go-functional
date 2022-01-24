package functional

import "testing"

func Test_Includes(t *testing.T) {
	intSlice := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}

	if !Includes(5)(intSlice) {
		t.Log("a number not found in a slice")
		t.Fail()
	}

	if Includes(10)(intSlice) {
		t.Log("a number in a slice failed")
		t.Fail()
	}

	stringSlice := []string{"foo", "bar", "baz"}

	if !Includes("foo")(stringSlice) {
		t.Log("a string not found in a slice")
		t.Fail()
	}

	if Includes("fail")(stringSlice) {
		t.Log("a string in a slice failed")
		t.Fail()
	}
}
