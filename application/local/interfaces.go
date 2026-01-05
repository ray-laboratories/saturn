package local

import (
	"context"
	"github.com/ray-laboratories/saturn/types"
)

type AccessorRepository interface {
	Get(ctx context.Context, username string) (*types.Accessor, error)
	Create(ctx context.Context, user *types.Accessor) error
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
