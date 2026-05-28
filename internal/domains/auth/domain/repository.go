package domain

import (
	"context"

	"github.com/homepage408/syncra/internal/domains/auth/domain/entity"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
}
