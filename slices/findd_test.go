package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindFirstConsecutiveDigits(t *testing.T) {

	matchDigits := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("find the digits when in one position of input string slice", func(t *testing.T) {
		target := []string{"hola", "235ty", "rty"}

		got := FindFirstConsecutiveDigits(target)
		want := []int{2, 3, 5}
		matchDigits(t, got, want)
	})

	t.Run("find the digits when in 2 positions of input string slice", func(t *testing.T) {
		target := []string{"hola", "ty235", "78r", "4567388"}

		got := FindFirstConsecutiveDigits(target)
		want := []int{2, 3, 5, 7, 8}
		matchDigits(t, got, want)
	})

	t.Run("find the digits when in several positions of input string slice", func(t *testing.T) {
		target := []string{"hola", "3", "5", "456", "89", "1atgsg3552", "67152"}

		got := FindFirstConsecutiveDigits(target)
		want := []int{3, 5, 4, 5, 6, 8, 9, 1}
		matchDigits(t, got, want)
	})

	t.Run("return an empty slice when input string slice is empty", func(t *testing.T) {
		target := []string{}

		got := FindFirstConsecutiveDigits(target)
		want := []int{}
		matchDigits(t, got, want)
	})

	t.Run("return an empty slice when input string slice doesn't have numbers/digits", func(t *testing.T) {
		target := []string{"hola", "wdfty", "rty"}

		got := FindFirstConsecutiveDigits(target)
		want := []int{}
		matchDigits(t, got, want)
	})

	t.Run("return an empty slice when input string slice is nil", func(t *testing.T) {
		got := FindFirstConsecutiveDigits(nil)
		want := []int{}
		matchDigits(t, got, want)
	})
}

func BenchmarkFindFirstConsecutiveDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindFirstConsecutiveDigits([]string{"hola", "3", "5", "456", "89", "1atgsg3552", "67152"})
	}
}

func ExampleFindFirstConsecutiveDigits() {
	got := FindFirstConsecutiveDigits([]string{"hola", "ty235", "78r", "4567388"})
	fmt.Print(got)
	fmt.Print(" ")
	got = FindFirstConsecutiveDigits(nil)
	fmt.Println(got)
	/* Output: [2 3 5 7 8] []*/
}
