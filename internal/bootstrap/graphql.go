package bootstrap

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	graph "github.com/homepage408/syncra/graph/generated"
	resolver "github.com/homepage408/syncra/graph/resolver"
	authUsecase "github.com/homepage408/syncra/internal/domains/auth/usecase"
)

func NewGraphQLHandlers(authSvc *authUsecase.Service) (http.Handler, http.Handler) {
	gqlResolver := resolver.NewResolver(authSvc)
	_ = gqlResolver

	graphqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &resolver.Resolver{},
	}))
	playgroundHandler := playground.Handler("Syncra GraphQL", "/query")

	return graphqlServer, playgroundHandler
}
