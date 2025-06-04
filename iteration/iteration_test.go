package iteration

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func randRange(max uint) uint {
	return rand.UintN(max)
}

func TestRepeat(t *testing.T) {
	t.Run("should repeat character 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectString(t, got, want)
	})

	t.Run("should return blank string", func(t *testing.T) {
		got := Repeat("a", 0)
		want := ""
		assertCorrectString(t, got, want)
	})
}

func assertCorrectString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Want %q, but got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", randRange(10))
	}
}

func ExampleRepeat() {
	str := Repeat("q", 10)
	fmt.Println(str)
	// Output: qqqqqqqqqq
}
