package application

import (
	"context"
	"github.com/ray-laboratories/saturn/types"
)

type AuthService interface {
	Login(ctx context.Context, username, password string) (string, error)
	Logout(ctx context.Context, token string) error
	Register(ctx context.Context, username, password string) (string, error)
	Validate(ctx context.Context, token string) (*types.Accessor, error)
}
