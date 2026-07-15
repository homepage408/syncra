package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/homepage408/syncra/config"
	"github.com/homepage408/syncra/db/sqlc"
	"github.com/homepage408/syncra/internal/domains/auth/domain"
	"github.com/homepage408/syncra/internal/domains/auth/domain/entity"
	"github.com/homepage408/syncra/pkg/jwt"
	"github.com/homepage408/syncra/pkg/logger"

	pswd "github.com/homepage408/syncra/pkg/password"
)

type Service struct {
	repo domain.UserRepository
	log  logger.Logger
	cfg  *config.Config
}

func New(repo domain.UserRepository, log logger.Logger, cfg *config.Config) *Service {
	return &Service{repo: repo, log: log, cfg: cfg}
}

func (s *Service) Login(ctx context.Context, email, password, userAgent string) (*entity.AuthMeta, error) {
	if s.cfg == nil {
		return nil, fmt.Errorf("config is not initialized")
	}

	if email == "" || password == "" {
		return nil, fmt.Errorf("invalid credentials")
	}

	sessionID := uuid.New()
	// UA := useragent.New(userAgent)
	// browser, _ := UA.Browser()
	// userAgentString := fmt.Sprintf("%s on %s", browser, UA.OS())

	existingUser, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		s.log.Error("failed to find user by email")
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	err = pswd.Verify(existingUser.PasswordHash, password)
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	token := jwt.JWT{
		SecretKey:           s.cfg.Token.JWTSecret,
		ExpiresAccessToken:  s.cfg.Token.JWTExpiry,
		ExpiresRefreshToken: s.cfg.Token.JWTRefreshExpiry,
	}
	accessToken, err := token.GenerateAccessToken(existingUser.ID.String(), sessionID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := token.GenerateRefreshToken(existingUser.ID.String(), sessionID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	expiresIn := int(s.cfg.Token.JWTExpiry.Seconds())

	// save to session
	err = s.repo.SaveSessions(ctx, sqlc.SaveSessionParams{
		ID:               sessionID,
		UserID:           existingUser.ID,
		UserAgent:        sql.NullString{String: userAgent, Valid: true},
		RefreshTokenHash: token.HashRefreshToken(refreshToken),
		ExpiredAt:        time.Now().Add(time.Duration(expiresIn)),
	})
	if err != nil {
		s.log.Error("")
		return nil, fmt.Errorf("Failed to create login session.: %w", err)
	}

	return &entity.AuthMeta{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}, nil
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

func (s *Service) GetAcctiveSessions(ctx context.Context, accessToken string) ([]*entity.Sessions, error) {
	if s == nil || s.cfg == nil {
		return nil, fmt.Errorf("service config is not initialized")
	}

	if accessToken == "" {
		return nil, fmt.Errorf("token is required")
	}

	response := make([]*entity.Sessions, 0)

	token := jwt.JWT{
		SecretKey:           s.cfg.Token.JWTSecret,
		ExpiresAccessToken:  s.cfg.Token.JWTExpiry,
		ExpiresRefreshToken: s.cfg.Token.JWTRefreshExpiry,
	}

	tokenVal, err := token.ParseToken(accessToken)
	if err != nil {
		return response, err
	}

	userIDstr, ok := tokenVal["sub"].(string)
	if !ok || userIDstr == "" {
		return nil, fmt.Errorf("invalid user ID in token")
	}

	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		return response, fmt.Errorf("invalid user ID in token: %w", err)
	}

	// repository
	data, err := s.repo.GetSessions(ctx, userID)
	if err != nil {
		return response, fmt.Errorf("failed to get sessions: %w", err)
	}

	if len(data) == 0 {
		return response, nil
	}

	for _, session := range data {
		var revokeTime *time.Time
		if session.RevokedAt.Valid {
			revokeTime = &session.RevokedAt.Time
		}

		response = append(response, &entity.Sessions{
			ID:               session.ID.String(),
			UserID:           session.UserID.String(),
			IpAddress:        session.IpAddress.IPNet.String(),
			RefreshTokenHash: session.RefreshTokenHash,
			RevokedAt:        revokeTime,
			UserAgent:        session.UserAgent.String,
			CreatedAt:        session.CreatedAt,
			ExpiredAt:        session.ExpiredAt,
		})
	}

	return response, err
}
