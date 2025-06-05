package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Should give correct sum with array of fixed size", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5}
		got := Sum(array)
		want := 15

		if want != got {
			t.Errorf("Want: %d, Got: %d, Given: %v", want, got, array)
		}
	})
}
