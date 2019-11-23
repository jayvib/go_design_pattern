package cache

import (
	"context"
	"github.com/allegro/bigcache"
)

type BigCache struct {
	bc *bigcache.BigCache
}

func (c *BigCache) Get(ctx context.Context, key string) (value interface{}, err error) {
	return
}

func (c *BigCache) Set(ctx context.Context, key string, value interface{}) error {
	return nil
}
