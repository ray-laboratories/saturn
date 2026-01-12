package local

import (
	"context"
	"github.com/ray-laboratories/saturn/types"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*types.User, error)
	Create(ctx context.Context, user *types.User) error
}

type Hasher interface {
	Hash(password string) (string, error)
	Compare(password, hash string) bool
}

type SessionRepository interface {
	Save(ctx context.Context, session *types.Session) error
	Get(ctx context.Context, token string) (*types.Session, error)
	Delete(ctx context.Context, token string) error
}

type Tokenizer interface {
	New() string
}
