package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing counter 3 times leaves it at 3", func(t *testing.T) {
		c := NewCounter()
		c.Inc()
		c.Inc()
		c.Inc()
		assertCounter(t, c, 3)
	})
	t.Run("run safety concurrently", func(t *testing.T) {
		wantedCount := 1000
		c := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				defer wg.Done()
				c.Inc()
			}()
		}

		wg.Wait()

		assertCounter(t, c, wantedCount)
	})
}

func assertCounter(t testing.TB, c *Counter, wantValue int) {
	t.Helper()
	if c.Value() != wantValue {
		t.Errorf("got %d, want %d", c.Value(), wantValue)
	}
}
