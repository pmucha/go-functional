package functional

import "testing"

func Test_All(t *testing.T) {
	positiveIntSlice := []int{1, 2, 3, 4, 5}
	positiveIntSliceFail := []int{0, 1, 2, 3, 4, 5}
	allPositiveInt := func(val int) (bool, error) {
		return val > 0, nil
	}

	if result, _ := All(allPositiveInt)(positiveIntSlice); !result {
		t.Log("all numbers should have been positive")
		t.Fail()
	}

	if result, _ := All(allPositiveInt)(positiveIntSliceFail); result {
		t.Log("at least one number should have been considered negative")
		t.Fail()
	}

	intSlice := []int{-2, -1, 1, 2}
	intSliceFail := []int{-2, -1, 0, 1, 2}
	allIntSlice := func(val int) (bool, error) {
		return val != 0, nil
	}

	if result, _ := All(allIntSlice)(intSlice); !result {
		t.Log("no number should have been equal to zero")
		t.Fail()
	}

	if result, _ := All(allIntSlice)(intSliceFail); result {
		t.Log("no number equal to zero should have failed")
		t.Fail()
	}
}
