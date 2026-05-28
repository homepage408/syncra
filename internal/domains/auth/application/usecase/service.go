package usecase

import (
	"context"

	"github.com/homepage408/syncra/internal/domains/auth/domain"
	"github.com/homepage408/syncra/pkg/logger"
)

type Service struct {
	repo domain.UserRepository
	log  logger.Logger
}

func New(repo domain.UserRepository, log logger.Logger) *Service {
	return &Service{repo: repo, log: log}
}

func (s *Service) Login(ctx context.Context, email, password string) (interface{}, error) {
	return s.repo.FindByEmail(ctx, email)
}

func (s *Service) Register(ctx context.Context, email, password string) (interface{}, error) {
	return nil, nil
}

func (s *Service) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	return "", nil
}
