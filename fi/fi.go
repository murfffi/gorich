package fi

import (
	"io"
	"time"

	"github.com/murfffi/gorich/lang"
	"github.com/murfffi/gorich/limit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RequireConditionWithTimeout asserts a complex condition within the given timeout.
// messages optionally contains either the message if condition failed, or that message together with the message if
// timeout exceeded.
// Note that this serves a different use case and operates differently than require.Eventually or require.Condition.
// This assertion pairs well with checks like lang.ReaderContains or regexp.MatchReader which may take a
// long time to complete or run indefinitely.
func RequireConditionWithTimeout(t require.TestingT, comp assert.Comparison, timeout time.Duration, messages ...string) {
	res, ok := limit.WithTimeout(timeout, comp)
	condMessage := "condition failed"
	if len(messages) > 0 {
		condMessage = messages[0]
	}
	timeoutMessage := "timeout exceeded"
	if len(messages) > 1 {
		timeoutMessage = messages[1]
	}
	require.True(t, ok, timeoutMessage)
	require.True(t, res, condMessage)
}

// RequireReaderContains asserts that substr appears is read by r within timeout
func RequireReaderContains(t require.TestingT, r io.Reader, substr string, timeout time.Duration, readerLabel string) {
	RequireConditionWithTimeout(
		t,
		func() bool {
			return lang.ReaderContains(r, substr)
		},
		timeout,
		readerLabel+" does not contain "+substr,
		"timeout exceeded while reading "+readerLabel,
	)
}
