package slices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstConsecutiveDigits(t *testing.T) {

	t.Run("find the digits when in one position of input string slice", func(t *testing.T) {
		got := FindFirstConsecutiveDigits([]string{"hola", "235ty", "rty"})
		want := []int{2, 3, 5}
		assert.Equal(t, want, got)
	})

	t.Run("find the digits when in 2 positions of input string slice", func(t *testing.T) {
		got := FindFirstConsecutiveDigits([]string{"hola", "ty235", "78r", "4567388"})
		want := []int{2, 3, 5, 7, 8}
		assert.Equal(t, want, got)
	})

	t.Run("find the digits when in several positions of input string slice", func(t *testing.T) {
		got := FindFirstConsecutiveDigits([]string{"hola", "3", "5", "456", "89", "1atgsg3552", "67152"})
		want := []int{3, 5, 4, 5, 6, 8, 9, 1}
		assert.Equal(t, want, got)
	})

	t.Run("return an empty slice when input string slice is empty", func(t *testing.T) {
		got := FindFirstConsecutiveDigits([]string{})
		want := []int{}
		assert.Equal(t, want, got)
	})

	t.Run("return an empty slice when input string slice doesn't have numbers/digits", func(t *testing.T) {
		got := FindFirstConsecutiveDigits([]string{"hola", "wdfty", "rty"})
		want := []int{}
		assert.Equal(t, want, got)
	})

	t.Run("return an empty slice when input string slice is nil", func(t *testing.T) {
		got := FindFirstConsecutiveDigits(nil)
		want := []int{}
		assert.Equal(t, want, got)
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
