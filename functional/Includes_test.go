package functional

import "testing"

func Test_Includes(t *testing.T) {
	intSlice := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}

	if result, _ := Includes(5)(intSlice); !result {
		t.Log("a number not found in a slice")
		t.Fail()
	}

	if result, _ := Includes(10)(intSlice); result {
		t.Log("a number should not have been found")
		t.Fail()
	}
}
