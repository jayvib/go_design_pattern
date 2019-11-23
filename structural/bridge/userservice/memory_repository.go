package userservice

import (
	"context"
	"errors"
	"sync"
)

// This is just a simple implementation.

type MemoryRepository struct {
	d     map[int]*User
	count int
	rwmu  sync.RWMutex
}

func (r *MemoryRepository) Create(ctx context.Context, user *User) error {
	r.rwmu.Lock()
	defer r.rwmu.Unlock()
	r.count++
	user.ID = r.count
	r.d[user.ID] = user
	return nil
}

func (r *MemoryRepository) Update(ctx context.Context, user *User) error {
	r.rwmu.Lock()
	defer r.rwmu.Unlock()
	_, ok := r.d[user.ID]
	if !ok {
		return errors.New("not exists")
	}
	r.d[user.ID] = user
	return nil
}

func (r *MemoryRepository) Delete(ctx context.Context, id int) error {
	r.rwmu.Lock()
	defer r.rwmu.Unlock()
	delete(r.d, id)
	return nil
}

func (r *MemoryRepository) Get(ctx context.Context, id int) (*User, error) {
	u, ok := r.d[id]
	if !ok {
		return nil, errors.New("not exists")
	}
	return u, nil
}

func (r *MemoryRepository) GetAll(ctx context.Context) ([]*User, error) {
	var out []*User
	for _, u := range r.d {
		out = append(out, u)
	}
	return out, nil
}
