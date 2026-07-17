package bootstrap

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/homepage408/syncra/config"
	authRest "github.com/homepage408/syncra/internal/domains/auth/interface/rest"
	"github.com/homepage408/syncra/internal/shared/middleware"
	"github.com/homepage408/syncra/pkg/logger"
)

type Routes struct {
	httpRouter *gin.Engine
	port       string
	logger     logger.Logger
}

func SetupRoutes(cfg *config.Config, authHandler *authRest.Handler, log logger.Logger) *Routes {
	r := gin.Default()

	port := strconv.Itoa(cfg.Server.Port)

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
		authGroup.GET("/sessions", middleware.Auth(), authHandler.GetActiveSessions)
		authGroup.DELETE("/sessions/:id", middleware.Auth(), authHandler.RemoveActiveSession)
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/refresh", authHandler.RefreshToken)
	}

	return &Routes{
		httpRouter: r,
		port:       ":" + port,
		logger:     log,
	}
}

func (r *Routes) Start(ctx context.Context) error {
	r.logger.Info("Starting HTTP server on " + r.port)
	return r.httpRouter.Run(r.port)
}
