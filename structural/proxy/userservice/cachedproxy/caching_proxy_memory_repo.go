package cachedproxy

import (
	"context"
	"github.com/jayvib/go_design_pattern/structural/proxy/userservice"
)

// The difference between the decorator pattern and the proxy pattern is
// the intent. With decorator pattern you extend the functionality of the
// existing service. While the Proxy pattern is the object do take control
// the life cycle, and access to the existing service.

type CachedProxyRepository struct {
	repo userservice.Repository
	cache map[int]*userservice.User // I could use LRU
}

func (c *CachedProxyRepository) Create(ctx context.Context, user *userservice.User) error {
	err := c.repo.Create(ctx, user)
	if err != nil {
		return err
	}
	c.cache[user.ID] = user
	return nil
}

func (c *CachedProxyRepository) Update(ctx context.Context, user *userservice.User) error {
	err := c.repo.Update(ctx, user)
	if err != nil {
		return err
	}
	c.cache[user.ID] = user
	return nil
}

func (c *CachedProxyRepository) Delete(ctx context.Context, id int) error {
	err := c.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	delete(c.cache, id)

	return nil
}

func (c *CachedProxyRepository) Get(ctx context.Context, id int) (*userservice.User, error) {
	if user, ok := c.cache[id]; ok {
		return user, nil
	}
	user, err := c.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	c.cache[id] = user
	return user, nil
}

func (c *CachedProxyRepository) GetAll(ctx context.Context) ([]*userservice.User, error) {
	return c.repo.GetAll(ctx)
}

