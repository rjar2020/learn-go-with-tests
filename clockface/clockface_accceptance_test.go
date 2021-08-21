package clockface_test

import (
	"testing"
	"time"

	"github.com/rjar2020/learn-go-with-tests/clockface"
	"github.com/stretchr/testify/require"
)

func TestSecondHandAtMidnight(t *testing.T) {

	assert := require.New(t)

	t.Run("second hand at midnight", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

		want := clockface.Point{X: 150, Y: 150 - 90}
		got := clockface.SecondHand(tm)

		assert.Equal(want, got)
	})

	t.Run("second hand at 30 seconds", func(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

		want := clockface.Point{X: 150, Y: 150 + 90}
		got := clockface.SecondHand(tm)

		assert.Equal(want, got)
	})

}
