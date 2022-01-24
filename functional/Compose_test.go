package functional

import "testing"

func Test_Compose(t *testing.T) {
	square := func(x int) int {
		return x * x
	}

	half := func(x int) int {
		return x / 2
	}

	add1 := func(x int) int {
		return x + 1
	}

	if result := Compose(add1, half, square)(10); result != 51 {
		t.Log("composing int failed; expected 51, result:", result)
		t.Fail()
	}

	repeat := func(s string) string {
		return s + s
	}

	pick := func(s string) string {
		return s[2:8]
	}

	if result := Compose(pick, repeat)("hello"); result != "llohel" {
		t.Log("composing string failed; expected \"llohel\", result:", result)
		t.Fail()
	}
}
