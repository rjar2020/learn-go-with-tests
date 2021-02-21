package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0
	assert.Equal(t, want, got, "The rectangle doesn't have the expected perimeter")
}

func TestArea(t *testing.T) {

	assert := require.New(t)

	t.Run("calculates a Rectangle's area", func(t *testing.T) {
		got := Rectangle{12, 6}.Area()
		want := 72.0
		assert.Equal(want, got)
	})

	t.Run("calculates a Circles's area", func(t *testing.T) {
		got := Circle{10}.Area()
		want := 314.1592653589793
		assert.Equal(want, got)
	})

}
