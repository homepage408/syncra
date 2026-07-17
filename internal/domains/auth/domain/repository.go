package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/homepage408/syncra/db/sqlc"
	"github.com/homepage408/syncra/internal/domains/auth/domain/entity"
)

type UserRepository interface {
	CheckEmailAndUsernameExisting(ctx context.Context, email, username string) (bool, bool, error)
	FindByEmail(ctx context.Context, email string) (sqlc.GetUserByEmailRow, error)
	Save(ctx context.Context, user *entity.User) error
	SaveSessions(ctx context.Context, params sqlc.SaveSessionParams) error
	GetSessions(ctx context.Context, userID uuid.UUID) ([]sqlc.GetAllSessionsRow, error)
	GetSessionById(ctx context.Context, userID, sessionID uuid.UUID) (sqlc.GetSessionByIdRow, error)
	RemoveSession(ctx context.Context, sessionID uuid.UUID) error
}
