package local

import (
	"context"
	"github.com/ray-laboratories/saturn/types"
)

type UserRepository interface {
	Get(ctx context.Context, id uint) (*types.User, error)
	GetByUsername(ctx context.Context, appName, username string) (*types.User, error)
}

type GroupRepository interface {
	Get(ctx context.Context, id uint) (*types.Group, error)
}
