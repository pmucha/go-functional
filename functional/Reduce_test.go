package functional

import (
	"strings"
	"testing"
)

func Test_Reduce(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}

	sum := func(result int, val int) (int, error) {
		return result + val, nil
	}

	if result, _ := Reduce(sum)(10)(intSlice); result != 25 {
		t.Log("reducing failed; expected: 25, got:", result)
		t.Fail()
	}

	stringSlice := []string{"Hello,", "world!"}
	concat := func(result string, val string) (string, error) {
		return result + " " + val, nil
	}

	if result, _ := Reduce(concat)("")(stringSlice); strings.TrimSpace(result) != "Hello, world!" {
		t.Log("reducing failed; slice reduced to:", result)
		t.Fail()
	}
}
