package bootstrap

import (
	"context"
	"net/http"
)

type Application struct {
	Server *http.Server
}

func NewApplication() (*Application, error) {

	router := SetupRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return &Application{
		Server: server,
	}, nil
}

func (a *Application) Run(ctx context.Context) error {
	return a.Server.ListenAndServe()
}
