package clockface

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSecondsInRadiant(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(testName((c.time)), func(t *testing.T) {
			got := secondsInRadians(c.time)
			require.Equal(t, c.angle, got)
		})
	}
}

func TestSecondHandVector(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName((c.time)), func(t *testing.T) {
			got := secondHandPoint(c.time)
			assertFloatEqual(c.point, got, t)
		})
	}
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func assertFloatEqual(expected, got Point, t *testing.T) {
	require.InDelta(t, expected.X, got.X, 1e-7)
}
