package entity

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID               uuid.UUID `json:"id"`
	UserID           uuid.UUID `json:"use_id"`
	RefreshTokenHash string    `json:"refresh_token_hash"`
	ExpiredAt        time.Time `json:"expired_at"`
}
