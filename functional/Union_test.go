package functional

import "testing"

func Test_Union(t *testing.T) {
	foo := []int{1, 1, 2, 3, 3, 4, 5}
	bar1 := []int{2, 6, 9}
	bar2 := []int{7, 9, 10}
	bar3 := []int{9, 11, 12}

	if uni := Union(foo)(bar1, bar2, bar3); len(uni) != 11 {
		t.Log("the Union procedure failed; expected 11 elements, result:", len(uni))
		t.Fail()
	}
}
