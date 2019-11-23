package userservice

import (
	"context"
	"errors"
	"time"
)

// the Bridge pattern compose of:
// Base Class
// Abstract Class
// Refined Abstraction
// Implementation

// https://refactoring.guru/design-patterns/bridge

//go:generate mockery --name Service

type User struct {
	ID          int
	FirstName   string
	LastName    string
	CreatedDate *time.Time
}

type Service interface {
	Register(ctx context.Context, user *User) error
	Delete(ctx context.Context, user *User) error
	Find(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, user *User) error
}

type serviceImpl struct {
	repo Repository
}

func (s *serviceImpl) Register(ctx context.Context, user *User) error {
	user.CreatedDate = TimePointer(time.Now())
	return s.repo.Create(ctx, user)
}

func (s *serviceImpl) Delete(ctx context.Context, user *User) error {
	if user.ID == 0 {
		return errors.New("empty id")
	}
	return s.repo.Delete(ctx, user.ID)
}

func (s *serviceImpl) Find(ctx context.Context, id int) (*User, error) {
	return s.repo.Get(ctx, id)
}

func (s *serviceImpl) Update(ctx context.Context, user *User) error {
	return s.repo.Update(ctx, user)
}
