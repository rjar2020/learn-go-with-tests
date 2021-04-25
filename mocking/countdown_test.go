package main

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCountdown(t *testing.T) {

	assert := require.New(t)

	t.Run("prints 3, 2, 1, Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &CountdownOperationSpy{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		assert.Equal(want, got)
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		assert.Equal(want, spySleepPrinter.Calls)
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	require.Equal(t, sleepTime, spyTime.durationSlept)
}

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const sleep = "sleep"
const write = "write"

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func BenchmarkCountdown(b *testing.B) {
	sleeper := &ConfigurableSleeper{1 * time.Millisecond, time.Sleep}
	buffer := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		Countdown(&buffer, sleeper)
	}
}

func ExampleCountdown() {
	sleeper := &ConfigurableSleeper{1 * time.Millisecond, time.Sleep}
	Countdown(os.Stdout, sleeper)
	/* Output: 3
2
1
Go!*/
}
