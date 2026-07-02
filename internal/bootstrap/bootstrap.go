package bootstrap

import (
	"context"
	"fmt"
	"log"

	"github.com/homepage408/syncra/config"
	"github.com/homepage408/syncra/db/sqlc"
	authRepo "github.com/homepage408/syncra/internal/domains/auth/infrastructure/persistence"
	authRest "github.com/homepage408/syncra/internal/domains/auth/interface/rest"
	authUsecase "github.com/homepage408/syncra/internal/domains/auth/usecase"
	"github.com/homepage408/syncra/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Application struct {
	db     *sqlx.DB
	logger logger.Logger
	routes *Routes // HTTP + GraphQL routes
}

// NewApplication - Main DI Container
func NewApplication(ctx context.Context) (*Application, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ✅ STEP 1: Load configuration (independent)
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("config load error: %w", err)
	}

	// ✅ STEP 2: Initialize logger (independent)
	log := logger.New(cfg.LogLevel)

	// ✅ STEP 3: Initialize database (independent)
	sqlDB, err := initDatabase(ctx, cfg, log)
	if err != nil {
		return nil, fmt.Errorf("database init error: %w", err)
	}

	// ✅ STEP 4: Create queries (depends on sqlDB)
	queries := sqlc.New(sqlDB)

	// ✅ STEP 5: Initialize repositories (depends on queries)
	authRepository := authRepo.New(queries, log)

	// ✅ STEP 6: Initialize use cases (depends on repositories)
	authService := authUsecase.New(authRepository, log)

	// ✅ STEP 7: Initialize HTTP handlers (depends on use cases)
	authHandler := authRest.New(authService, log)
	// postHandler := postRest.New(postService, log)

	// ✅ STEP 8: Setup routes
	routes := SetupRoutes(cfg, authHandler, log)

	// ✅ STEP 9: Setup GraphQL
	setupGraphQL(routes, authService)

	return &Application{
		db:     sqlDB,
		logger: log,
		routes: routes,
	}, nil
}

func (app *Application) Run(ctx context.Context) error {
	return app.routes.Start(ctx)
}
