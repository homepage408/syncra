package persistence

import (
	"context"
	"errors"

	"github.com/homepage408/syncra/db/sqlc"
	"github.com/homepage408/syncra/internal/domains/auth/domain"
	"github.com/homepage408/syncra/internal/domains/auth/domain/entity"
	"github.com/homepage408/syncra/pkg/logger"
)

type PostgresUserRepository struct {
	queries *sqlc.Queries
	log     logger.Logger
}

func New(queries *sqlc.Queries, log logger.Logger) domain.UserRepository {
	return &PostgresUserRepository{queries: queries, log: log}
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user.Email == "" {
		return nil, errors.New("user not found")
	}
	return &entity.User{
		// ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		// Role:         user.Role,
		// IsActive:     user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *entity.User) error {
	return r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		// ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		// Role:         user.Role,
		// IsActive:     user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}
