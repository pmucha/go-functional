package functional

import "testing"

func Test_All(t *testing.T) {
	positiveIntSlice := []int{1, 2, 3, 4, 5}
	positiveIntSliceFail := []int{0, 1, 2, 3, 4, 5}
	allPositiveInt := func(val int) bool {
		return val > 0
	}

	if !All(allPositiveInt)(positiveIntSlice) {
		t.Log("all numbers should have been positive")
		t.Fail()
	}

	if All(allPositiveInt)(positiveIntSliceFail) {
		t.Log("all numbers positive should have failed")
		t.Fail()
	}

	intSlice := []int{-2, -1, 1, 2}
	intSliceFail := []int{-2, -1, 0, 1, 2}
	allIntSlice := func(val int) bool {
		return val != 0
	}

	if !All(allIntSlice)(intSlice) {
		t.Log("no number should have been equal to zero")
		t.Fail()
	}

	if All(allIntSlice)(intSliceFail) {
		t.Log("no number equal to zero should have failed")
		t.Fail()
	}

	allStringSlice := []string{"foo", "bar", "baz"}
	allStringSliceFail := []string{"", "foo", "bar", "baz"}
	allNonEmpty := func(val string) bool {
		return len(val) > 0
	}

	if !All(allNonEmpty)(allStringSlice) {
		t.Log("no string should be empty")
		t.Fail()
	}

	if All(allNonEmpty)(allStringSliceFail) {
		t.Log("no string empty should have failed")
		t.Fail()
	}
}
