package functional

import (
	"strings"
	"testing"
)

func Test_Reduce(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}

	reduceSum := func(result int, val int) int {
		return result + val
	}

	if result := Reduce(reduceSum)(10)(intSlice); result != 25 {
		t.Log("reducing failed; expected: 25, got:", result)
		t.Fail()
	}

	stringSlice := []string{"Hello,", "world!"}
	reduceConcat := func(result string, val string) string {
		return result + " " + val
	}

	if result := strings.TrimSpace(Reduce(reduceConcat)("")(stringSlice)); result != "Hello, world!" {
		t.Log("reducing failed; slice reduced to:", result)
		t.Fail()
	}
}
