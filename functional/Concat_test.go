package functional

import "testing"

func Test_Concat(t *testing.T) {
	src1 := []int{1, 1, 2, 3, 3, 4, 5}
	src2 := []int{2, 6, 9}
	with1 := []int{7, 9, 10}
	with2 := []int{9, 11, 12}

	if concatenated := Concat(src1, src2)(with1, with2); len(concatenated) != 16 {
		t.Log("the Concat procedure failed; expected 16 elements, result:", len(concatenated))
		t.Fail()
	}
}
