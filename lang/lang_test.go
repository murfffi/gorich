package lang_test

import (
	"io"
	"strings"
	"sync"
	"testing"

	"github.com/murfffi/gorich/lang"
	"github.com/stretchr/testify/require"
)

func TestReaderContains(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		require.True(t, lang.ReaderContains(strings.NewReader("Hello World!"), "World!"))
	})
	t.Run("endless", func(t *testing.T) {
		r, w := io.Pipe()
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			var err error
			for err == nil {
				_, err = w.Write([]byte("ab"))
			}
			wg.Done()
		}()
		require.True(t, lang.ReaderContains(r, "babab"))
		require.NoError(t, r.Close())
		wg.Wait()
	})

}
