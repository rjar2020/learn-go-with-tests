package iteration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {

	t.Run("Repeat repeats a character, 10 times", func(t *testing.T) {
		want := Repeat("a", 10)
		got := "aaaaaaaaaa"
		assert.Equal(t, want, got)
	})

	t.Run("Repeat doesn't repeat a character when pasing a negative multiplier", func(t *testing.T) {
		got := Repeat("z", -1)
		want := "z"
		assert.Equal(t, want, got)
	})

	t.Run("Repeat doesn't repeat a character when pasing a zero multiplier", func(t *testing.T) {
		got := Repeat("b", 0)
		want := "b"
		assert.Equal(t, want, got)
	})

	t.Run("EnhancedRepeat doesn't repeat a character when pasing neutral (1) multiplier", func(t *testing.T) {
		got := Repeat("W", 1)
		want := "W"
		assert.Equal(t, want, got)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func TestEnhancedRepeat(t *testing.T) {

	t.Run("EnhancedRepeat repeats a character, 10 times", func(t *testing.T) {
		got := EnhancedRepeat("a", 10)
		want := "aaaaaaaaaa"
		assert.Equal(t, want, got)
	})

	t.Run("EnhancedRepeat doesn't repeat a character when pasing a negative multiplier", func(t *testing.T) {
		got := EnhancedRepeat("z", -1)
		want := "z"
		assert.Equal(t, want, got)
	})

	t.Run("EnhancedRepeat doesn't repeat a character when pasing a zero multiplier", func(t *testing.T) {
		got := EnhancedRepeat("b", 0)
		want := "b"
		assert.Equal(t, want, got)
	})

	t.Run("EnhancedRepeat doesn't repeat a character when pasing neutral (1) multiplier", func(t *testing.T) {
		got := EnhancedRepeat("W", 1)
		want := "W"
		assert.Equal(t, want, got)
	})
}

func BenchmarkEnhancedRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EnhancedRepeat("a", 5)
	}
}

func ExampleEnhancedRepeat() {
	repeated := EnhancedRepeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
