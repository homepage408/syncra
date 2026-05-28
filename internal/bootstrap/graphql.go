package bootstrap

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/homepage408/syncra/graph/resolver"
	authUsecase "github.com/homepage408/syncra/internal/domains/auth/application/usecase"
	postUsecase "github.com/homepage408/syncra/internal/domains/post/application/usecase"
)

func NewGraphQLHandlers(authSvc *authUsecase.Service, postSvc *postUsecase.Service) (http.Handler, http.Handler) {
	gqlResolver := resolver.NewResolver(authSvc, postSvc)
	_ = gqlResolver

	graphqlServer := handler.New(nil)
	playgroundHandler := playground.Handler("Syncra GraphQL", "/graphql")

	return graphqlServer, playgroundHandler
}
