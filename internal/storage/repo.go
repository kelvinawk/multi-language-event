package store

import (
	"context"
)

type BaseRepository[T any] interface {
	Create(ctx context.Context, t *T) error
	Update(ctx context.Context, t *T) error
	GetById(uuid string, c context.Context) (*T, error)
}
