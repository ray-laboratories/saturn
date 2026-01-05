package sqlite

import (
	"context"
	"database/sql"
	db "github.com/ray-laboratories/saturn/db/sqlc"
	"github.com/ray-laboratories/saturn/types"
)

type AccessorRepository struct {
	q *db.Queries
}

func (a AccessorRepository) Get(ctx context.Context, username string) (*types.Accessor, error) {
	acc, err := a.q.GetAccessor(ctx, username)
	if err != nil {
		return nil, err
	}
	return types.NewAccessor(acc.Username, acc.HashedPassword), nil
}

func (a AccessorRepository) Create(ctx context.Context, user *types.Accessor) error {
	_, err := a.q.InsertAccessor(ctx, db.InsertAccessorParams{
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
	})
	return err
}

func NewAccessorRepository(sqlite *sql.DB) *AccessorRepository {
	return &AccessorRepository{
		q: db.New(sqlite),
	}
}
