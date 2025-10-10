// Package limit provides means to limit the duration of functions
// It is meant to complement samber/lo utilities like Attempt and WaitFor .
//
// Functions in this package pair well with lang.ReaderContains
// or regexp.MatchReader which are not context-aware and can take a long time,
// especially on slow or unbounded readers like process pipes or network streams.
package limit

import (
	"context"
	"time"
)

// WithinCtx attempts to run the given function until the context expires.
// The function itself can, optionally, limit its own execution on the same context.
// Returns the result of the function and whether the function completed while the context was valid.
// Functions with multiple return values can use samber/lo Tuples to pack them.
//
// This utility can turn a function which is not context-aware, or only partially so, into a strictly
// context-aware one. Additionally, in combination with context.WithDeadline, it can put time limits on
// execution attempts.
func WithinCtx[T any](ctx context.Context, f func(context.Context) T) (result T, ok bool) {
	var zero T
	ch := make(chan T, 1) // capacity 1 to allow the goroutine to write to the channel and exit even if we stopped waiting
	go func() {
		res := f(ctx) // on separate lines for debugging
		ch <- res
	}()
	select {
	case <-ctx.Done():
		return zero, false
	case res := <-ch:
		return res, true
	}
}

// WithTimeout attempts to run the given context-unaware function within the given timeout duration.
// Use WithinCtx directly if there is parent context or the function is context-aware.
func WithTimeout[T any](d time.Duration, f func() T) (result T, ok bool) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), d)
	defer cancelCtx()
	return WithinCtx(ctx, func(context.Context) T {
		return f()
	})
}
