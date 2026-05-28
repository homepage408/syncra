package persistence

import (
	"context"

	"github.com/homepage408/syncra/db/sqlc"
	"github.com/homepage408/syncra/internal/domains/post/domain"
	"github.com/homepage408/syncra/internal/domains/post/domain/entity"
	"github.com/homepage408/syncra/pkg/logger"
)

type PostgresPostRepository struct {
	queries *sqlc.Queries
	log     logger.Logger
}

func New(queries *sqlc.Queries, log logger.Logger) domain.PostRepository {
	return &PostgresPostRepository{queries: queries, log: log}
}

func (r *PostgresPostRepository) Create(ctx context.Context, post *entity.Post) error {
	return r.queries.CreatePost(ctx, sqlc.CreatePostParams{
		ID:        post.ID,
		AuthorID:  post.AuthorID,
		Title:     post.Title,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	})
}

func (r *PostgresPostRepository) FindByID(ctx context.Context, id string) (*entity.Post, error) {
	post, err := r.queries.GetPostByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entity.Post{
		ID:        post.ID,
		AuthorID:  post.AuthorID,
		Title:     post.Title,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}

func (r *PostgresPostRepository) List(ctx context.Context) ([]*entity.Post, error) {
	records, err := r.queries.ListPosts(ctx)
	if err != nil {
		return nil, err
	}
	posts := make([]*entity.Post, 0, len(records))
	for _, item := range records {
		posts = append(posts, &entity.Post{
			ID:        item.ID,
			AuthorID:  item.AuthorID,
			Title:     item.Title,
			Body:      item.Body,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}
	return posts, nil
}
