package cache

import (
	"context"
	"github.com/coocood/freecache"
)

type FreeCache struct {
	c *freecache.Cache
}

func (c *FreeCache) Get(ctx context.Context, key string) (value interface{}, err error) {
	return
}

func (c *FreeCache) Set(ctx context.Context, key string, value interface{}) (err error) {
	return
}
