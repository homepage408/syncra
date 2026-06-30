package usecase

import (
	"context"

	"github.com/homepage408/syncra/internal/domains/post/domain"
	"github.com/homepage408/syncra/internal/domains/post/domain/entity"
	"github.com/homepage408/syncra/pkg/logger"
)

type Service struct {
	repo domain.PostRepository
	log  logger.Logger
}

func New(repo domain.PostRepository, log logger.Logger) *Service {
	return &Service{repo: repo, log: log}
}

func (s *Service) CreatePost(ctx context.Context, title, body, authorID string) (*entity.Post, error) {
	return nil, nil
}

func (s *Service) GetPost(ctx context.Context, id string) (*entity.Post, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) ListPosts(ctx context.Context) ([]*entity.Post, error) {
	return s.repo.List(ctx)
}
