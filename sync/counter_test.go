package sync

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCounter(t *testing.T) {

	assert := require.New(t)

	t.Run("increment the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		three := 3

		for i := 0; i < three; i++ {
			counter.Inc()
		}

		assert.Equal(three, counter.Value())
	})

	t.Run("it runs safely cuncurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		concurrentIncrement(counter, wantedCount)

		assert.Equal(wantedCount, counter.Value())
	})

}

func ExampleCounter() {
	wantedCount := 1000
	counter := NewCounter()

	var wg sync.WaitGroup
	wg.Add(wantedCount)

	for i := 0; i < wantedCount; i++ {
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("%d", counter.Value())
	/*Output: 1000*/
}

func BenchmarkCounter(b *testing.B) {
	wantedCount := b.N
	counter := NewCounter()

	concurrentIncrement(counter, wantedCount)
}

func concurrentIncrement(counter *Counter, wantedCount int) {
	var wg sync.WaitGroup
	wg.Add(wantedCount)

	for i := 0; i < wantedCount; i++ {
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
}
