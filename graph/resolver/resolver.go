package resolver

import (
	authUsecase "github.com/homepage408/syncra/internal/domains/auth/usecase"
)

type Resolver struct {
	AuthService *authUsecase.Service
}

func NewResolver(authSvc *authUsecase.Service) *Resolver {
	return &Resolver{AuthService: authSvc}
}
