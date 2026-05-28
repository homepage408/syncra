package main

import (
	"context"
	"log"

	"github.com/homepage408/syncra/internal/bootstrap"
)

func main() {
	app, err := bootstrap.NewApplication(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
