package bootstrap

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	graph "github.com/homepage408/syncra/graph/generated"
	resolver "github.com/homepage408/syncra/graph/resolver"
	authUsecase "github.com/homepage408/syncra/internal/domains/auth/usecase"
	postUsecase "github.com/homepage408/syncra/internal/domains/post/usecase"
)

func NewGraphQLHandlers(authSvc *authUsecase.Service, postSvc *postUsecase.Service) (http.Handler, http.Handler) {
	gqlResolver := resolver.NewResolver(authSvc, postSvc)
	_ = gqlResolver

	graphqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &resolver.Resolver{},
	}))
	playgroundHandler := playground.Handler("Syncra GraphQL", "/query")

	return graphqlServer, playgroundHandler
}
