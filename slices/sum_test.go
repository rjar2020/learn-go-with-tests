package slices

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3})
	want := 6
	assert.Equal(t, want, got, "The result doesn't match the sum of all elements in the input slice")
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

	//Won't replace it with testify to let reflect.DeepEqual usage example
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
