package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0
	assert.Equal(t, want, got, "The rectangle doesn't have the expected perimeter")
}

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		got := Rectangle{12, 6}.Area()
		want := 72.0
		assert.Equal(t, want, got, "The shape doesn't have the expected area")
	})

	t.Run("circles", func(t *testing.T) {
		got := Circle{10}.Area()
		want := 314.1592653589793
		assert.Equal(t, want, got, "The shape doesn't have the expected area")
	})

}
