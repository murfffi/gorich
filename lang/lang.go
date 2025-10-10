// Package lang container general purpose programming helpers
package lang

import (
	"bufio"
	"io"
	"regexp"
)

// IfEmpty returns alt if value is empty (zero-value), and value itself, if not empty
// Simplified variant of samber/lo.CoalesceOrEmpty
func IfEmpty[T comparable](value, alt T) T {
	var zero T
	if value != zero {
		return value
	}
	return alt
}

// Bind converts a single-parameter function to a no-parameter one by binding the given
// value to the parameter. Useful together with fi.NoErrorF or defer.
// Bind implements the missing Partial0 from samber/lo.
func Bind[T any, E any](f func(t T) E, t T) func() E {
	return func() E {
		return f(t)
	}
}

// ReaderContains reads the given stream until it finds substr or reading fails
// Unlike a solution based on bufio.Scanner, it doesn't depend on newlines in the stream.
func ReaderContains(reader io.Reader, substr string) bool {
	runeReader, ok := reader.(io.RuneReader)
	if !ok {
		runeReader = bufio.NewReader(reader)
	}
	return regexp.MustCompile(regexp.QuoteMeta(substr)).MatchReader(runeReader)
}
