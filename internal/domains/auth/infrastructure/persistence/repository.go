package persistence

import (
	"context"

	"github.com/google/uuid"
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
	_, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &entity.User{}, nil
}

func (r *PostgresUserRepository) CheckEmailAndUsernameExisting(ctx context.Context, email, username string) (bool, bool, error) {
	result, err := r.queries.CheckEmailAndUsernameExisting(ctx, sqlc.CheckEmailAndUsernameExistingParams{
		Email:    email,
		Username: username,
	})
	if err != nil {
		return false, false, err
	}

	return result.EmailExists, result.UsernameExists, nil
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *entity.User) error {
	return r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		ID:           uuid.New(),
		Email:        user.Email,
		Username:     user.Username,
		FullName:     user.FullName,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	})
}
