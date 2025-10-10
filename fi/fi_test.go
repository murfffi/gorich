package fi_test

import (
	"testing"
	"time"

	"github.com/murfffi/gorich/fi"
)

func TestRequireConditionWithTimeout(t *testing.T) {
	t.Run("cond is good", func(t *testing.T) {
		fi.RequireConditionWithTimeout(t, func() bool {
			return true
		}, time.Second)
	})
}
