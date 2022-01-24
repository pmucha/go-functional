package functional

import "testing"

func Test_Uniq(t *testing.T) {
	intSlice := []int{1, 1, 2, 2, 3, 3}
	if uniqIntSlice := Uniq(intSlice); len(uniqIntSlice) != 3 {
		t.Log("a slice isn't unique; input:", intSlice, "result:", uniqIntSlice)
		t.Fail()
	}
}
