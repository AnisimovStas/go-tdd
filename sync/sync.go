package sync

import "sync"

type Counter struct {
	my    sync.Mutex
	count int
}

func (c *Counter) Inc() {
	c.my.Lock()
	defer c.my.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}
