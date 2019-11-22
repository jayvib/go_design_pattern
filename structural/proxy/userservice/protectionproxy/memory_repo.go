package protectionproxy

import (
	"context"
	"errors"
	"github.com/jayvib/go_design_pattern/structural/proxy/userservice"
)

// this is just an example......

type ProtectionProxyRepository struct {
	username string
	password string
	repo userservice.Repository
}

func (r *ProtectionProxyRepository) authenticate() bool {
	// don't use this style in real code. ─=≡Σ((( つ＞＜)つ
	if r.username == "luffy" && r.password == "pirateking" {
		return true
	}
	return false
}


func (r *ProtectionProxyRepository) Create(ctx context.Context, user *userservice.User) error {
	if !r.authenticate() {
		return errors.New("unauthorized")
	}
	return r.repo.Create(ctx, user)
}

func (r *ProtectionProxyRepository) Update(ctx context.Context, user *userservice.User) error {
	if !r.authenticate() {
		return errors.New("unauthorized")
	}
	return r.repo.Update(ctx, user)
}

func (r *ProtectionProxyRepository) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

func (r *ProtectionProxyRepository) Get(ctx context.Context, id int) (*userservice.User, error) {
	if !r.authenticate() {
		return nil, errors.New("unauthorized")
	}
	return r.repo.Get(ctx, id)
}

func (r *ProtectionProxyRepository) GetAll(ctx context.Context) ([]*userservice.User, error) {
	if !r.authenticate() {
		return nil, errors.New("unauthorized")
	}
	return r.repo.GetAll(ctx)
}
