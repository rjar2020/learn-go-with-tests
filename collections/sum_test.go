package collections

import "testing"

func TestSum(t *testing.T) {

	assertCorrectResult := func(t testing.TB, numbers []int, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectResult(t, numbers, got, want)
	})

}
