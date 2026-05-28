package bootstrap

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {

	r := chi.NewRouter()

	SetupMiddlewares(r)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return r
}
