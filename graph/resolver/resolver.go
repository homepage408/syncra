package resolver

import (
	authUsecase "github.com/homepage408/syncra/internal/domains/auth/application/usecase"
	postUsecase "github.com/homepage408/syncra/internal/domains/post/application/usecase"
)

type Resolver struct {
	AuthService *authUsecase.Service
	PostService *postUsecase.Service
}

func NewResolver(authSvc *authUsecase.Service, postSvc *postUsecase.Service) *Resolver {
	return &Resolver{AuthService: authSvc, PostService: postSvc}
}
