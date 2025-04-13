// Package lang container general purpose programming helpers
package lang

// IfEmpty returns alt if value is empty (zero-value), and value itself, if not empty
func IfEmpty[T comparable](value, alt T) T {
	// Simplified variant of lo.Coalesce.
	// lo.Coalesce returns two results which makes it harder to use in some cases
	var zero T
	if value != zero {
		return value
	}
	return alt
}
