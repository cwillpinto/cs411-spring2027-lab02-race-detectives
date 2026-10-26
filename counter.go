package counter

import "sync"

// Counter is shared by concurrent workers. Preserve its public API.
type Counter struct {
	mu    sync.Mutex
	value int
}

// Add increments by delta. TODO: guard the full read-modify-write operation.
func (c *Counter) Add(delta int) {
	current := c.value
	c.value = current + delta
}

// Value returns a snapshot. TODO: synchronize this access too.
func (c *Counter) Value() int {
	return c.value
}
