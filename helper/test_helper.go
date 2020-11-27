package helper

import (
	"runtime"
	"testing"
)

func ShouldBeEqual(t *testing.T, expected, actual interface{}) {
	skip := 1

	if expected != actual {
		_, fn, line, _ := runtime.Caller(skip)
		t.Fatalf("\nExp: %v\nGot: %v\nLocation: %s:%d", expected, actual, fn, line)
	}
}
