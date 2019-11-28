package cache

import "context"

type Cache interface {
	CacheGetter // composition
	CacheSetter
}

type CacheGetter interface {
	Get(ctx context.Context, key string) (value interface{}, err error)
}

type CacheSetter interface {
	Set(ctx context.Context, key string, value interface{}) error
}

func DoSomethingWithCache(c CacheSetter) error { // following the interface segragation principle
	key := "123456789"
	value := "hello world"
	return c.Set(context.Background(), key, value)
}
