package assert

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/Snikimonkd/tf/internal/assertions"
)

// Equal - checks if two values are equal
func Equal(t *testing.T, want any, got any, opts ...cmp.Option) {
	t.Helper()
	assertions.Equal(t, want, got, t.Fatalf, opts...)
}

// NotEqual - checks if two values are not equal
func NotEqual(t *testing.T, want any, got any, opts ...cmp.Option) {
	t.Helper()
	assertions.NotEqual(t, want, got, t.Fatalf, opts...)
}

// Nil - checks if value is nil
func Nil(t *testing.T, got any) {
	t.Helper()
	assertions.Nil(t, got, t.Fatalf)
}

// NotNil - checks if value is not nil
func NotNil(t *testing.T, got any) {
	t.Helper()
	assertions.NotNil(t, got, t.Fatalf)
}

// True - checks if value is true
func True(t *testing.T, got bool) {
	t.Helper()
	assertions.True(t, got, t.Fatalf)
}

// False - checks if value is false
func False(t *testing.T, got bool) {
	t.Helper()
	assertions.False(t, got, t.Fatalf)
}
