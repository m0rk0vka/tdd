package maps

import "testing"

func TestSearch(t *testing.T) {
	d := Dictionary{"test": "it's test!"}
	t.Run("known word", func(t *testing.T) {

		got, _ := d.Search("test")

		want := "it's test!"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := d.Search("word")

		assertErrors(t, got, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Fatal("expected to get an error")
	}
}
