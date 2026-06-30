package resolver

import (
	authUsecase "github.com/homepage408/syncra/internal/domains/auth/usecase"
	postUsecase "github.com/homepage408/syncra/internal/domains/post/usecase"
)

type Resolver struct {
	AuthService *authUsecase.Service
	PostService *postUsecase.Service
}

func NewResolver(authSvc *authUsecase.Service, postSvc *postUsecase.Service) *Resolver {
	return &Resolver{AuthService: authSvc, PostService: postSvc}
}
