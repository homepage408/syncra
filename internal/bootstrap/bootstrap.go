package bootstrap

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/homepage408/syncra/config"
	"github.com/homepage408/syncra/db/sqlc"
	authUsecase "github.com/homepage408/syncra/internal/domains/auth/application/usecase"
	authRepo "github.com/homepage408/syncra/internal/domains/auth/infrastructure/persistence"
	authRest "github.com/homepage408/syncra/internal/domains/auth/interface/rest"
	postUsecase "github.com/homepage408/syncra/internal/domains/post/application/usecase"
	postRepo "github.com/homepage408/syncra/internal/domains/post/infrastructure/persistence"
	postRest "github.com/homepage408/syncra/internal/domains/post/interface/rest"
	"github.com/homepage408/syncra/pkg/logger"
)

type Application struct {
	db     *sql.DB
	logger logger.Logger
	routes *Routes // HTTP + GraphQL routes
}

// NewApplication - Main DI Container
func NewApplication(ctx context.Context) (*Application, error) {
	// ✅ STEP 1: Load configuration (independent)
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("config load error: %w", err)
	}

	// ✅ STEP 2: Initialize logger (independent)
	log := logger.New(cfg.LogLevel)

	// ✅ STEP 3: Initialize database (independent)
	db, err := initDatabase(ctx, cfg, log)
	if err != nil {
		return nil, fmt.Errorf("database init error: %w", err)
	}

	// ✅ STEP 4: Create queries (depends on db)
	queries := sqlc.New(db)

	// ✅ STEP 5: Initialize repositories (depends on queries)
	authRepository := authRepo.New(queries, log)
	postRepository := postRepo.New(queries, log)

	// ✅ STEP 6: Initialize use cases (depends on repositories)
	authService := authUsecase.New(authRepository, log)
	postService := postUsecase.New(postRepository, log)

	// ✅ STEP 7: Initialize HTTP handlers (depends on use cases)
	authHandler := authRest.New(authService, log)
	postHandler := postRest.New(postService, log)

	// ✅ STEP 8: Setup routes
	routes := SetupRoutes(authHandler, postHandler, log)

	// ✅ STEP 9: Setup GraphQL
	setupGraphQL(routes, authService, postService)

	return &Application{
		db:     db,
		logger: log,
		routes: routes,
	}, nil
}

func (app *Application) Run(ctx context.Context) error {
	return app.routes.Start(ctx)
}
