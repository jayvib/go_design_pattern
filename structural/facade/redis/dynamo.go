package redis

import "context"

// Facade Interface

type Redis interface {
	Create(ctx context.Context, obj interface{}) error
	Get(ctx context.Context, hash string) (interface{}, error)
	Delete(ctx context.Context, hash string) error
	Update(ctx context.Context, obj interface{}) error
}

type redisImpl struct {
}

func (d *redisImpl) Create(ctx context.Context, obj interface{}) error {
	return nil
}

func (d *redisImpl) Get(ctx context.Context, hash string) (interface{}, error) {
	return nil, nil
}

func (d *redisImpl) Delete(ctx context.Context, hash string) error {
	return nil
}

func (d *redisImpl) Update(ctx context.Context, obj interface{}) error {
	return nil
}
