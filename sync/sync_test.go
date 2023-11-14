package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing counter 3 times leaves it at 3", func(t *testing.T) {
		c := Counter{}
		c.Inc()
		c.Inc()
		c.Inc()
		assertCounter(t, c, 3)
	})
}

func assertCounter(t testing.TB, c Counter, wantValue int) {
	t.Helper()
	if c.Value() != 3 {
		t.Errorf("got %d, want %d", c.Value(), 3)
	}
}
