package fi

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Requirement associates checks with a value
type Requirement[T any] struct {
	// The value returned by Require if the check passes
	Val   T
	Check func(t require.TestingT)
}

// NoError defines a requirement that the result of the function
// can be used, and the associated error is nil
func NoError[T any](val T, err error) Requirement[T] {
	return Requirement[T]{
		Val: val,
		Check: func(t require.TestingT) {
			require.NoError(t, err)
		},
	}
}

// FileExists defines a requirement that the file exists (experimental)
func FileExists(path string) Requirement[string] {
	return Requirement[string]{
		Val: path,
		Check: func(t require.TestingT) {
			require.FileExists(t, path)
		},
	}
}

// Require returns the Requirement value if the Check doesn't fail the current test
func (r Requirement[T]) Require(t require.TestingT) T {
	r.Check(t)
	return r.Val
}

// NoErrorF fails the current test if f returns an error. Useful in defer.
func NoErrorF(f func() error, t assert.TestingT) {
	err := f()
	assert.NoError(t, err)
}
