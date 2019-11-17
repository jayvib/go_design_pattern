package counter

import "sync"

type Counter interface {
	AddOne() int
}

type counter struct {
	count int
}

func (c *counter) AddOne() int {
	c.count++ // prone to race condition
	return c.count
}

var (
	once sync.Once
	defaultCounter *counter
	defaultCounterInitializedCalls int // use for testing
)

func GetDefaultCounter() Counter {
	once.Do(func(){
		defaultCounter = &counter{}
		defaultCounterInitializedCalls++
	})
	return defaultCounter
}
