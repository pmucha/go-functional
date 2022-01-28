package functional

import "testing"

func Test_Any(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	intSliceFail := []int{1, 3, 5, 7, 9}
	isEven := func(val int) (bool, error) {
		return val%2 == 0, nil
	}

	if result, err := Any(isEven)(intSlice); !result || err != nil {
		t.Log("there should be at least one even number")
		t.Fail()
	}

	if result, err := Any(isEven)(intSliceFail); !result || err != nil {
		t.Log("there should not have been any even numbers")
		t.Fail()
	}

	intSlice = []int{-1, 0, 1, 2, 3}
	intSliceFail = []int{0, 1, 2, 3}
	isNegative := func(val int) (bool, error) {
		return val < 0, nil
	}

	if result, err := Any(isNegative)(intSlice); !result || err != nil {
		t.Log("at least one number should have been negative")
		t.Fail()
	}

	if result, err := Any(isNegative)(intSliceFail); !result || err != nil {
		t.Log("there should not have been any negative numbers")
		t.Fail()
	}
}
