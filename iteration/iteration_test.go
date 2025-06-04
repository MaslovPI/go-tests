package iteration

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func randRange(max uint) uint {
	return rand.UintN(max)
}

func assertCorrectString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Want %q, but got %q", want, got)
	}
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

func TestRepeatWithSeparation(t *testing.T) {
	t.Run("should repeat character 5 times with separation", func(t *testing.T) {
		got := RepeatWithSeparation("a", ",", 5)
		want := "a,a,a,a,a"
		assertCorrectString(t, got, want)
	})

	t.Run("should return blank string", func(t *testing.T) {
		got := RepeatWithSeparation("a", ",", 0)
		want := ""
		assertCorrectString(t, got, want)
	})
}

func BenchmarkRepeatWithSeparation(b *testing.B) {
	for b.Loop() {
		RepeatWithSeparation("a", ",", randRange(10))
	}
}

func ExampleRepeatWithSeparation() {
	str := RepeatWithSeparation("q", ",", 10)
	fmt.Println(str)
	// Output: q,q,q,q,q,q,q,q,q,q
}
