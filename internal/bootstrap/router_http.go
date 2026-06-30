package bootstrap

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	authRest "github.com/homepage408/syncra/internal/domains/auth/interface/rest"
	postRest "github.com/homepage408/syncra/internal/domains/post/interface/rest"
	"github.com/homepage408/syncra/pkg/logger"
)

type Routes struct {
	httpRouter *gin.Engine
	port       string
	logger     logger.Logger
	// graphQLHandler *handler.Server // Untuk GraphQL
}

func SetupRoutes(authHandler *authRest.Handler, postHandler *postRest.Handler, log logger.Logger) *Routes {
	r := gin.Default()

	// Middleware stack
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"message": "Welcome to Syncra API"})
	})

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	authGroup := r.Group("/api/v1/auth")
	{
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/refresh", authHandler.RefreshToken)
	}

	postGroup := r.Group("/api/v1/posts")
	{
		postGroup.GET("", gin.WrapF(postHandler.ListPosts))
		postGroup.POST("", gin.WrapF(postHandler.CreatePost))
		postGroup.GET("/:id", gin.WrapF(postHandler.GetPost))
	}

	return &Routes{
		httpRouter: r,
		port:       ":4500",
		logger:     log,
	}
}

func (r *Routes) Start(ctx context.Context) error {
	r.logger.Info("Starting HTTP server on " + r.port)
	return r.httpRouter.Run(r.port)
}
