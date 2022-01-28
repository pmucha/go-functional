package functional

import "testing"

func Test_Concat(t *testing.T) {
	src1 := []int{1, 1, 2, 3, 3, 4, 5}
	src2 := []int{2, 6, 9}

	if concatenated, _ := Concat(src1, src2); len(concatenated) != 10 {
		t.Log("the Concat procedure failed; expected 10 elements, result:", len(concatenated))
		t.Fail()
	}
}
