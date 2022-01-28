package functional

import "testing"

func Test_Filter(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5, 6}
	intSliceFail := []int{1, 3, 5}

	isEven := func(x int) (bool, error) {
		return x%2 == 0, nil
	}

	if result, _ := Filter(isEven)(intSlice); len(result) != 3 {
		t.Log("filtering failed; expected 3, result:", len(result))
		t.Fail()
	}

	if result, _ := Filter(isEven)(intSliceFail); len(result) != 0 {
		t.Log("filtering was supposed to fail; expected 0, result:", len(result))
		t.Fail()
	}

	stringSlice := []string{"foo", "", "bar", "", "baz"}
	notEmpty := func(s string) (bool, error) {
		return s != "", nil
	}

	if result, _ := Filter(notEmpty)(stringSlice); len(result) != 3 {
		t.Log("filtering failed; expected 3, result:", len(result))
		t.Fail()
	}
}
