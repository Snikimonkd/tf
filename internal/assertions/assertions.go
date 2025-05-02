package assertions

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// FailFn - func to fail test
type FailFn func(format string, args ...any)

// Equal - checks if two values are equal
func Equal(t *testing.T, want any, got any, failFn FailFn, opts ...cmp.Option) {
	t.Helper()
	if !cmp.Equal(want, got, opts...) {
		failFn("want != got, diff: %v", cmp.Diff(want, got, opts...))
	}
}

// NotEqual - checks if two values are not equal
func NotEqual(t *testing.T, want any, got any, failFn FailFn, opts ...cmp.Option) {
	t.Helper()
	if cmp.Equal(want, got, opts...) {
		failFn("want == got")
	}
}

// Nil - checks if value is nil
func Nil(t *testing.T, got any, failFn FailFn) {
	t.Helper()
	if got != nil {
		failFn("got != nil, got: %v", got)
	}
}

// Nil - checks if value is not nil
func NotNil(t *testing.T, got any, failFn FailFn) {
	t.Helper()
	if got == nil {
		failFn("got == nil")
	}
}

// True - checks if value is true
func True(t *testing.T, got bool, failFn FailFn) {
	t.Helper()
	if got != true {
		failFn("got != true")
	}
}

// False - checks if value is false
func False(t *testing.T, got bool, failFn FailFn) {
	t.Helper()
	if got != false {
		failFn("got != false")
	}
}
