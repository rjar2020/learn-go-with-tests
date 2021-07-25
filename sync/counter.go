package sync

import "sync"

//Counter is a synchronized counter, ready to support concurrent operations
type Counter struct {
	mu    sync.Mutex
	value int
}

//NewCounter is the preferred way to initialise a counter
func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
