package singleton

import (
	"sync"
)

type AddOner interface {
	AddOne()
}

type Count struct {
	count int
}

func (c *Count) AddOne() {
	c.count++
}

var (
	instance *Count
	once     sync.Once
)

func GetInstance() AddOner {
	// This is an idiomatic way for the singleton
	// http://marcio.io/2015/07/singleton-pattern-in-go/
	once.Do(func() {
		instance = &Count{}
	})
	return instance
}
