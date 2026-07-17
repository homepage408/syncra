package entity

import (
	"time"
)

type AuthMeta struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	SessionId    string `json:"session_id"`
}

type Sessions struct {
	ID               string     `json:"id"`
	UserID           string     `json:"user_id"`
	RefreshTokenHash string     `json:"refresh_token_hash"`
	UserAgent        string     `json:"user_agent,omitempty"`
	IpAddress        string     `json:"ip_address,omitempty"`
	ExpiredAt        time.Time  `json:"expired_at"`
	RevokedAt        *time.Time `json:"revoked_at"`
	CreatedAt        time.Time  `json:"created_at"`
}
