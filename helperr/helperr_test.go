package helperr_test

import (
	"testing"

	"github.com/murfffi/gorich/helperr"
)

type mockCloser struct {
	closeCalled bool
}

func (m *mockCloser) Close() error {
	m.closeCalled = true
	return nil
}

func TestCloseQuietly(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		helperr.CloseQuietly(nil)
	})

	t.Run("non-nil", func(t *testing.T) {
		mc := &mockCloser{}
		helperr.CloseQuietly(mc)
		if !mc.closeCalled {
			t.Error("expected Close() to be called, but it was not")
		}
	})
}
