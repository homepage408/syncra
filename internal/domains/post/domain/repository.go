package domain

import (
	"context"

	"github.com/homepage408/syncra/internal/domains/post/domain/entity"
)

type PostRepository interface {
	Create(ctx context.Context, post *entity.Post) error
	FindByID(ctx context.Context, id string) (*entity.Post, error)
	List(ctx context.Context) ([]*entity.Post, error)
}
