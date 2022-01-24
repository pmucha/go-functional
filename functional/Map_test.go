package functional

import "testing"

func Test_Map(t *testing.T) {
	intSlice := []int{1, 2, 3}

	mapDouble := func(val int) int {
		return val * 2
	}

	if result := Map(mapDouble)(intSlice); result[0] != 2 {
		t.Log("mapping failed; input:", intSlice, "result:", result)
		t.Fail()
	}
}
