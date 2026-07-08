package domain

import (
	"context"

	"github.com/homepage408/syncra/internal/domains/auth/domain/entity"
)

type UserRepository interface {
	CheckEmailAndUsernameExisting(ctx context.Context, email, username string) (bool, bool, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
}
