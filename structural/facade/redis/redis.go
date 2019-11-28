package redis

import "context"
import "github.com/go-redis/redis/v7"

// Facade Interface

// Gonna use hash type
type Redis interface {
	Create(ctx context.Context, key string,obj interface{}) error
	Get(ctx context.Context, hash string, storeTo interface{}) (error)
	Delete(ctx context.Context, hash string) error
	Update(ctx context.Context, obj interface{}) error
}

type redisImpl struct {
	client *redis.Client
}

func (d *redisImpl) Create(ctx context.Context, key string,obj interface{}) error {
	h := ObjectMapper(obj)
	return d.client.HMSet(key, h).Err()
}

func (d *redisImpl) Get(ctx context.Context, hash string, storeTo interface{}) (error) {
	v, err := d.client.HGetAll(hash).Result()
	if err != nil {
		return err
	}
	return MapParser(v, storeTo)
}

func (d *redisImpl) Delete(ctx context.Context, key string) error {
	return d.client.Del(key).Err()
}

func (d *redisImpl) Update(ctx context.Context, key string,obj interface{}) error {
	h := ObjectMapper(obj)
	return d.client.HMSet(key, h).Err()
}


