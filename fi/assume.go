package fi

import "testing"

// Helpers for testing.T.Skip*

// SkipLongTest skips the current test when -short is set
func SkipLongTest(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
}
