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
	t.Run("Add new key", func(t *testing.T) {
		d := Dictionary{}
		key := "key"
		value := "value"

		err := d.Add(key, value)

		assertErrors(t, err, nil)
		assertValues(t, d, key, value)
	})
	t.Run("Existing key", func(t *testing.T) {
		key := "key"
		value := "value"
		d := Dictionary{key: value}
		err := d.Add(key, "new value")

		assertErrors(t, err, ErrKeyExist)
		assertValues(t, d, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update existing key", func(t *testing.T) {
		key := "key"
		value := "value"
		d := Dictionary{key: value}
		newValue := "new value"

		err := d.Update(key, newValue)

		assertErrors(t, err, nil)
		assertValues(t, d, key, newValue)
	})
	t.Run("Trying to update new key", func(t *testing.T) {
		key := "key"
		newValue := "value"
		d := Dictionary{}

		err := d.Update(key, newValue)

		assertErrors(t, err, ErrKeyDoesNotExist)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertValues(t testing.TB, d Dictionary, key, value string) {
	t.Helper()
	got, err := d.Search(key)
	if err != nil {
		t.Fatal("shoud find added word:", err)
	}
	if got != value {
		t.Errorf("got %q want %q", got, value)
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
