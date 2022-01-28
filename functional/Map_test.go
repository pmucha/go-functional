package functional

import "testing"

func Test_Map(t *testing.T) {
	intSlice := []int{1, 2, 3}

	double := func(val int) (int, error) {
		return val * 2, nil
	}

	if result, _ := Map(double)(intSlice); result[0] != 2 {
		t.Log("mapping failed; input:", intSlice, "result:", result)
		t.Fail()
	}
}
