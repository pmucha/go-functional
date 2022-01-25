package functional

import "testing"

func Test_Any(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	intSliceFail := []int{1, 3, 5, 7, 9}
	isEven := func(val int) bool {
		return val%2 == 0
	}

	if !Any(isEven)(intSlice) {
		t.Log("there should be at least one even number")
		t.Fail()
	}

	if Any(isEven)(intSliceFail) {
		t.Log("one even number should have failed")
		t.Fail()
	}

	intSlice = []int{-1, 0, 1, 2, 3}
	intSliceFail = []int{0, 1, 2, 3}
	isNegative := func(val int) bool {
		return val < 0
	}

	if !Any(isNegative)(intSlice) {
		t.Log("at least one number should have been negative")
		t.Fail()
	}

	if Any(isNegative)(intSliceFail) {
		t.Log("one number negative should have failed")
		t.Fail()
	}
}
