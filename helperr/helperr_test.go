package helperr_test

import (
	"testing"

	"github.com/murfffi/gorich/helperr"
)

func TestCloseQuietly(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		helperr.CloseQuietly(nil)
	})
}
