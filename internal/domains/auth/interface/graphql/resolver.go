package gql

import (
	"context"

	authUsecase "github.com/homepage408/syncra/internal/domains/auth/usecase"
)

type AuthPayload struct {
	AccessToken  string
	RefreshToken string
}

type Resolver struct {
	service *authUsecase.Service
}

func New(service *authUsecase.Service) *Resolver {
	return &Resolver{service: service}
}

func (r *Resolver) Login(ctx context.Context, email, password string) (*AuthPayload, error) {
	// Implement domain interaction here
	return &AuthPayload{AccessToken: "", RefreshToken: ""}, nil
}

func (r *Resolver) Register(ctx context.Context, email, password string) (*AuthPayload, error) {
	return &AuthPayload{AccessToken: "", RefreshToken: ""}, nil
}
