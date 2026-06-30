package gql

import (
	"context"
	"time"

	postUsecase "github.com/homepage408/syncra/internal/domains/post/usecase"
)

type Post struct {
	ID        string
	AuthorID  string
	Title     string
	Body      string
	CreatedAt time.Time
}

type Resolver struct {
	service *postUsecase.Service
}

func New(service *postUsecase.Service) *Resolver {
	return &Resolver{service: service}
}

func (r *Resolver) Posts(ctx context.Context) ([]*Post, error) {
	return []*Post{}, nil
}

func (r *Resolver) Post(ctx context.Context, id string) (*Post, error) {
	return &Post{ID: id}, nil
}

func (r *Resolver) CreatePost(ctx context.Context, title, body, authorID string) (*Post, error) {
	return &Post{ID: "", Title: title, Body: body, AuthorID: authorID, CreatedAt: time.Now()}, nil
}
