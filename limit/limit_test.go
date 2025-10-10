package limit_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/murfffi/gorich/limit"
	"github.com/stretchr/testify/require"
)

func TestWithinCtx(t *testing.T) {
	t.Run("hit limit", func(t *testing.T) {
		ch := make(chan bool)
		ctx, cancel := context.WithCancel(t.Context())
		defer cancel()

		wgStart := sync.WaitGroup{}
		wgStart.Add(1)

		go func() {
			wgStart.Wait()
			cancel() // make sure we cancel only after f started
		}()

		wg := sync.WaitGroup{}
		wg.Add(1)
		res, ok := limit.WithinCtx(ctx, func(ctx context.Context) bool {
			wgStart.Done()
			res := <-ch
			wg.Done()
			return res
		})

		// ensure f exists so we don't leak a goroutine
		t.Cleanup(func() {
			ch <- true
			wg.Wait()
		})

		require.False(t, ok)
		require.False(t, res)
	})
	t.Run("simple", func(t *testing.T) {
		res, ok := limit.WithinCtx(t.Context(), func(ctx context.Context) int {
			return 1
		})
		require.True(t, ok)
		require.Equal(t, 1, res)
	})
}

func TestWithTimeout(t *testing.T) {
	t.Run("hit limit", func(t *testing.T) {
		ch := make(chan bool)

		wg := sync.WaitGroup{}
		wg.Add(1)
		res, ok := limit.WithTimeout(20*time.Millisecond, func() bool {
			res := <-ch
			wg.Done()
			return res
		})

		// ensure f exists so we don't leak a goroutine
		t.Cleanup(func() {
			ch <- true
			wg.Wait()
		})

		require.False(t, ok)
		require.False(t, res)
	})
	t.Run("simple", func(t *testing.T) {
		res, ok := limit.WithTimeout(10*time.Second, func() int {
			return 1
		})
		require.True(t, ok)
		require.Equal(t, 1, res)
	})
}
