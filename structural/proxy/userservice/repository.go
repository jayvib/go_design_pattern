package userservice

import "context"

// Basic CRUD
type Repository interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
}
