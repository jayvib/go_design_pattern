package virtualproxy

import (
	"context"
	"github.com/jayvib/go_design_pattern/structural/proxy/userservice"
	"sync"
)

// https://roberto.selbach.ca/zero-values-in-go-and-lazy-initialization/

// Lazy Initialization(Virtual Proxy):; This is when you have a heavyweight service object that wastes system resources
// by being always up, even though you only need it from time to time. Instead of creating the object when the app
// launches, you can delay the object's initialization to a time when it's really needed.

type VirtualProxyMemoryRepo struct {
	repo *userservice.MemoryRepository
	once sync.Once
}

func (r *VirtualProxyMemoryRepo) lazyInit() {
	r.once.Do(func() {
		r.repo = userservice.NewMemoryRepository()
	})
}

func (r *VirtualProxyMemoryRepo) Create(ctx context.Context, user *userservice.User) error {
	r.lazyInit()
	return r.repo.Create(ctx, user)
}

func (r *VirtualProxyMemoryRepo) Update(ctx context.Context, user *userservice.User) error {
	r.lazyInit()
	return r.repo.Update(ctx, user)
}

func (r *VirtualProxyMemoryRepo) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

func (r *VirtualProxyMemoryRepo) Get(ctx context.Context, id int) (*userservice.User, error) {
	r.lazyInit()
	return r.repo.Get(ctx, id)
}

func (r *VirtualProxyMemoryRepo) GetAll(ctx context.Context) ([]*userservice.User, error) {
	r.lazyInit()
	return r.repo.GetAll(ctx)
}
