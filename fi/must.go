package fi

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// NoError defines a requirement that the result of the function
// is can be used and the associated error is nil
func NoError[T any](val T, err error) Requirement[T] {
	return Requirement[T]{
		val: val,
		check: func(t require.TestingT) {
			require.NoError(t, err)
		},
	}
}

// Requirement associate checks with a value
type Requirement[T any] struct {
	val   T
	check func(t require.TestingT)
}

// Require returns the Requirement value if the check doesn't fail the current test
func (r Requirement[T]) Require(t require.TestingT) T {
	r.check(t)
	return r.val
}

// NoErrorF fails the current test if f returns an error. Useful in defer.
func NoErrorF(f func() error, t assert.TestingT) {
	err := f()
	assert.NoError(t, err)
}
