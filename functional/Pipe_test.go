package functional

import (
	"errors"
	"testing"
)

func Test_Pipe(t *testing.T) {
	square := func(x int) (int, error) {
		return x * x, nil
	}

	half := func(x int) (int, error) {
		return x / 2, nil
	}

	add1 := func(x int) (int, error) {
		return x + 1, nil
	}

	if result, _ := Pipe(square, half, add1)(10); result != 51 {
		t.Log("piping int failed; expected 51, result:", result)
		t.Fail()
	}

	repeat := func(s string) (string, error) {
		return s + s, nil
	}

	pick := func(s string) (string, error) {
		return s[2:8], nil
	}

	forcedErr := func(s string) (string, error) {
		return "", errors.New("forced error")
	}

	if result, e := Pipe(repeat, pick)("hello"); result != "llohel" || e != nil {
		t.Log("piping string failed; expected \"llohel\", result:", result)
		t.Fail()
	}

	if _, e := Pipe(repeat, forcedErr, pick)("hello"); e == nil {
		t.Log("forced error failed; error:", e)
		t.Fail()
	}
}
