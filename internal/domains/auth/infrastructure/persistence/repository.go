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

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (sqlc.GetUserByEmailRow, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return sqlc.GetUserByEmailRow{}, err
	}

	return user, nil
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
	})
}

func (r *PostgresUserRepository) SaveSessions(ctx context.Context, params sqlc.SaveSessionParams) error {
	return r.queries.SaveSession(ctx, params)
}

func (r *PostgresUserRepository) GetSessions(ctx context.Context, userID uuid.UUID) ([]sqlc.GetAllSessionsRow, error) {
	return r.queries.GetAllSessions(ctx, userID)
}

func (r *PostgresUserRepository) GetSessionById(ctx context.Context, userID, sessionID uuid.UUID) (sqlc.GetSessionByIdRow, error) {
	return r.queries.GetSessionById(ctx, sqlc.GetSessionByIdParams{UserID: userID, ID: sessionID})
}

func (r *PostgresUserRepository) RemoveSession(ctx context.Context, sessionID uuid.UUID) error {
	return r.queries.UpdateSession(ctx, sessionID)
}
