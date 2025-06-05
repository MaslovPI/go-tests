package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Should give correct sum with slice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		got := Sum(slice)
		want := 15

		if want != got {
			t.Errorf("Want: %d, Got: %d, Given: %v", want, got, slice)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Should give slice of sums for provided slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{0, 9})
		want := []int{6, 9}
		assertSumsEqual(t, want, got)
	})
}

func TestSumAllTail(t *testing.T) {
	t.Run("Should give slice of tail sums for provided slices", func(t *testing.T) {
		got := SumAllTail([]int{1, 2, 3}, []int{0, 9})
		want := []int{5, 9}
		assertSumsEqual(t, want, got)
	})
	t.Run("Should safely process empty slice", func(t *testing.T) {
		got := SumAllTail([]int{1, 2, 3}, []int{})
		want := []int{5, 0}
		assertSumsEqual(t, want, got)
	})
}

func assertSumsEqual(t testing.TB, want, got []int) {
	t.Helper()
	if !slices.Equal(want, got) {
		t.Errorf("Want: %v, Got: %v", want, got)
	}
}
