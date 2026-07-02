package bootstrap

import (
	"context"
	"fmt"

	"github.com/homepage408/syncra/config"
	"github.com/homepage408/syncra/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func initDatabase(ctx context.Context, cfg *config.Config, log logger.Logger) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.Database.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)

	log.Info("Database connected successfully")
	return db, nil
}
