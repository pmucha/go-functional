package functional

import "testing"

func Test_Pipe(t *testing.T) {
	square := func(x int) int {
		return x * x
	}

	half := func(x int) int {
		return x / 2
	}

	add1 := func(x int) int {
		return x + 1
	}

	if result := Pipe(square, half, add1)(10); result != 51 {
		t.Log("piping int failed; expected 51, result:", result)
		t.Fail()
	}

	repeat := func(s string) string {
		return s + s
	}

	pick := func(s string) string {
		return s[2:8]
	}

	if result := Pipe(repeat, pick)("hello"); result != "llohel" {
		t.Log("piping string failed; expected \"llohel\", result:", result)
		t.Fail()
	}
}
