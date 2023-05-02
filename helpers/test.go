package helpers

import (
	"fmt"
	"testing"
)

func Assert(t testing.TB, condition bool) {
	t.Helper()
	if !condition {
		t.Errorf("assertion failed")
	}
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func ToString[T any](values T) string {
	return fmt.Sprintf("%v", values)
}
