package functional

import "testing"

func Test_Concat(t *testing.T) {
	foo := []int{1, 1, 2, 3, 3, 4, 5}
	bar1 := []int{2, 6, 9}
	bar2 := []int{7, 9, 10}
	bar3 := []int{9, 11, 12}

	if concatenated := Concat(foo)(bar1, bar2, bar3); len(concatenated) != 16 {
		t.Log("the Concat procedure failed; expected 16 elements, result:", len(concatenated))
		t.Fail()
	}
}
