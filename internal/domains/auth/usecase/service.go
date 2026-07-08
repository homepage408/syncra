package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/homepage408/syncra/internal/domains/auth/domain"
	"github.com/homepage408/syncra/internal/domains/auth/domain/entity"
	"github.com/homepage408/syncra/pkg/logger"

	pswd "github.com/homepage408/syncra/pkg/password"
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

func (s *Service) Register(ctx context.Context, user *entity.User) (interface{}, error) {
	emailExists, usernameExists, err := s.repo.CheckEmailAndUsernameExisting(ctx, user.Email, user.Username)
	if err != nil {
		return nil, err
	}

	if emailExists {
		return nil, fmt.Errorf("email already exists")
	}

	if usernameExists {
		return nil, fmt.Errorf("username already exists")
	}

	passwordHash, err := pswd.Hash(user.PasswordHash)
	if err != nil {
		return nil, err
	}

	newUser := &entity.User{
		ID:           uuid.NewString(),
		Email:        user.Email,
		Username:     user.Username,
		FullName:     user.FullName,
		PasswordHash: passwordHash,
	}

	err = s.repo.Save(ctx, newUser)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *Service) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	return "", nil
}
