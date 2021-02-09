package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("Repeat a character, 10 times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Repeat a character, pasing a negative multiplier", func(t *testing.T) {
		repeated := Repeat("z", -1)
		expected := "z"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Repeat a character, 0 times", func(t *testing.T) {
		repeated := Repeat("b", 0)
		expected := "b"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Repeat a character, 1 time", func(t *testing.T) {
		repeated := Repeat("W", 1)
		expected := "W"
		assertCorrectMessage(t, repeated, expected)
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

	assertCorrectMessage := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("EnhancedRepeat repeats a character, 10 times", func(t *testing.T) {
		repeated := EnhancedRepeat("a", 10)
		expected := "aaaaaaaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("EnhancedRepeat repeats a character, pasing a negative multiplier", func(t *testing.T) {
		repeated := EnhancedRepeat("z", -1)
		expected := "z"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("EnhancedRepeat repeats a character, 0 times", func(t *testing.T) {
		repeated := EnhancedRepeat("b", 0)
		expected := "b"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("EnhancedRepeat repeats a character, 1 time", func(t *testing.T) {
		repeated := EnhancedRepeat("W", 1)
		expected := "W"
		assertCorrectMessage(t, repeated, expected)
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
