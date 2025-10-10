package fi

import (
	"time"

	"github.com/murfffi/gorich/attempt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RequireConditionWithTimeout assert a complex condition, like require.Condition, but within the given timeout.
// messages optionally contains either the message if condition failed, or that message together with message if
// timeout exceeded.
// Note that this serves a different use case and operates differently than require.Eventually.
func RequireConditionWithTimeout(t require.TestingT, comp assert.Comparison, d time.Duration, messages ...string) {
	res, ok := attempt.WithTimeout(d, comp)
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
