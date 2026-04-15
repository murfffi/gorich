package fi_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/murfffi/gorich/fi"
	"github.com/murfffi/gorich/lang"
	"github.com/stretchr/testify/require"
)

type testingT func(format string, args ...any)

func (f testingT) Errorf(format string, args ...any) {
	f(format, args...)
}

func TestNoErrorF(t *testing.T) {
	called := false
	var calledStub testingT = func(string, ...any) {
		called = true
	}
	fi.NoErrorF(lang.Bind(errors.New, "test"), calledStub)
	require.Equal(t, true, called)
}

func TestNoError(t *testing.T) {
	re := fi.NoError(regexp.Compile(".")).Require(t)
	require.NotNil(t, re)
}
