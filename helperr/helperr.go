// Package helperr contains error-handling helpers (get it?)
package helperr

import (
	"io"
	"strings"
)

// CloseQuietly closes the given object ignoring errors. Useful in defer.
func CloseQuietly(r io.Closer) {
	if r != nil {
		_ = r.Close()
	}
}

// ContainsAny checks if the error message contains any of the substrings
func ContainsAny(err error, subs ...string) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	for _, substr := range subs {
		if strings.Contains(msg, substr) {
			return true
		}
	}
	return false
}
