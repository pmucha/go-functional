package functional

import "testing"

func Test_Filter(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5, 6}
	intSliceFail := []int{1, 3, 5}

	isEven := func(x int) bool {
		return x%2 == 0
	}

	if result := Filter(isEven)(intSlice); len(result) != 3 {
		t.Log("filtering failed; expected 3, result:", len(result))
		t.Fail()
	}

	if result := Filter(isEven)(intSliceFail); len(result) != 0 {
		t.Log("filtering was supposed to fail; expected 0, result:", len(result))
		t.Fail()
	}

	stringSlice := []string{"foo", "", "bar", "", "baz"}
	notEmpty := func(s string) bool {
		return s != ""
	}

	if result := Filter(notEmpty)(stringSlice); len(result) != 3 {
		t.Log("filtering failed; expected 3, result:", len(result))
		t.Fail()
	}
}
