package functional

import (
	"errors"
	"testing"
)

func Test_Compose(t *testing.T) {
	square := func(x int) (int, error) {
		return x * x, nil
	}

	half := func(x int) (int, error) {
		return x / 2, nil
	}

	add1 := func(x int) (int, error) {
		return x + 1, nil
	}

	composed := Compose(add1, half, square)

	if result, _ := composed(10); result != 51 {
		t.Log("composing int failed; expected 51, result:", result)
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

	if result, _ := Compose(pick, repeat)("hello"); result != "llohel" {
		t.Log("composing string failed; expected \"llohel\", result:", result)
		t.Fail()
	}

	if _, e := Compose(pick, forcedErr, repeat)("hello"); e == nil {
		t.Log("forced error failed; error:", e)
		t.Fail()
	}
}
