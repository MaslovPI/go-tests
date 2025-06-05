package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStringsEqual(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("love")
		if err == nil {
			t.Fatal("Expected an exception")
		}
		assertError(t, err, errNotFound)
	})
}

func assertStringsEqual(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want: %q, got: %q", want, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
