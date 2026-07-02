package bootstrap

import (
	"github.com/gin-gonic/gin"
	authUsecase "github.com/homepage408/syncra/internal/domains/auth/usecase"
)

func setupGraphQL(routes *Routes, authSvc *authUsecase.Service) {
	graphqlHandler, playgroundHandler := NewGraphQLHandlers(authSvc)
	routes.httpRouter.Any("/query", gin.WrapH(graphqlHandler))
	routes.httpRouter.Any("/playground", gin.WrapH(playgroundHandler))
}
