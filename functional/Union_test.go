package functional

import (
	"testing"
)

func Test_Union(t *testing.T) {
	src1 := []int{1, 1, 2, 3, 3, 4, 5}
	src2 := []int{2, 6, 9}
	if uni, _ := Union(src1, src2); len(uni) != 7 {
		t.Log("the Union procedure failed; expected 7 elements, result:", len(uni))
		t.Fail()
	}
}
