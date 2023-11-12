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

func TestAdd(t *testing.T) {
	d := Dictionary{}
	d.Add("test", "it's a test")
	want := "it's a test"
	got, err := d.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
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
