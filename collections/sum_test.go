package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}

}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3})
	}
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 10, 1000}
	fmt.Println(Sum(numbers))
	// Output: 1016
}

func TestSumAllTails(t *testing.T) {

	checksum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checksum(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checksum(t, got, want)
	})

}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{1, 2, 3}, []int{0, 9})
	}
}

func ExampleSumAllTails() {
	got := SumAllTails([]int{}, []int{3, 4, 5}, []int{1, 2, 3, 10, 1000})
	fmt.Println(got)
	// Output: [0 9 1015]
}
