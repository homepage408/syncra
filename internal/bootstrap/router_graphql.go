package bootstrap

import (
	"github.com/gin-gonic/gin"
	authUsecase "github.com/homepage408/syncra/internal/domains/auth/usecase"
	postUsecase "github.com/homepage408/syncra/internal/domains/post/usecase"
)

func setupGraphQL(routes *Routes, authSvc *authUsecase.Service, postSvc *postUsecase.Service) {
	graphqlHandler, playgroundHandler := NewGraphQLHandlers(authSvc, postSvc)
	routes.httpRouter.Any("/query", gin.WrapH(graphqlHandler))
	routes.httpRouter.Any("/playground", gin.WrapH(playgroundHandler))
}
